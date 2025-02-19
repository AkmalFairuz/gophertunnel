package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AddPlayer is sent by the server to the client to make a player entity show up client-side. It is one of the
// few entities that cannot be sent using the AddActor packet.
type AddPlayer struct {
	// UUID is the UUID of the player. It is the same UUID that the client sent in the Login packet at the
	// start of the session. A player with this UUID must exist in the player list (built up using the
	// PlayerList packet), for it to show up in-game.
	UUID uuid.UUID
	// Username is the name of the player. This username is the username that will be set as the initial
	// name tag of the player.
	Username string
	// EntityRuntimeID is the runtime ID of the player. The runtime ID is unique for each world session, and
	// entities are generally identified in packets using this runtime ID.
	EntityRuntimeID uint64
	// PlatformChatID is an identifier only set for particular platforms when chatting (presumably only for
	// Nintendo Switch). It is otherwise an empty string, and is used to decide which players are able to
	// chat with each other.
	PlatformChatID string
	// Position is the position to spawn the player on. If the player is on a distance that the viewer cannot
	// see it, the player will still show up if the viewer moves closer.
	Position mgl32.Vec3
	// Velocity is the initial velocity the player spawns with. This velocity will initiate client side
	// movement of the player.
	Velocity mgl32.Vec3
	// Pitch is the vertical rotation of the player. Facing straight forward yields a pitch of 0. Pitch is
	// measured in degrees.
	Pitch float32
	// Yaw is the horizontal rotation of the player. Yaw is also measured in degrees.
	Yaw float32
	// HeadYaw is the same as Yaw, except that it applies specifically to the head of the player. A different
	// value for HeadYaw than Yaw means that the player will have its head turned.
	HeadYaw float32
	// HeldItem is the item that the player is holding. The item is shown to the viewer as soon as the player
	// itself shows up. Needless to say that this field is rather pointless, as additional packets still must
	// be sent for armour to show up.
	HeldItem protocol.ItemInstance
	// GameType is the game type of the player. If set to GameTypeSpectator, the player will not be shown to viewers.
	GameType int32
	// EntityMetadata is a map of entity metadata, which includes flags and data properties that alter in
	// particular the way the player looks. Flags include ones such as 'on fire' and 'sprinting'.
	// The metadata values are indexed by their property key.
	EntityMetadata map[uint32]any
	// EntityUniqueID is a unique identifier of the player. It appears it is not required to fill this field
	// out with a correct value. Simply writing 0 seems to work.
	EntityUniqueID int64
	// PlayerPermissions is the permission level of the player as it shows up in the player list built up using
	// the PlayerList packet.
	PlayerPermissions byte
	// CommandPermissions is a set of permissions that specify what commands a player is allowed to execute.
	CommandPermissions byte
	// Layers contains all ability layers and their potential values. This should at least have one entry, being the
	// base layer.
	Layers []protocol.AbilityLayer
	// EntityLinks is a list of entity links that are currently active on the player. These links alter the
	// way the player shows up when first spawned in terms of it shown as riding an entity. Setting these
	// links is important for new viewers to see the player is riding another entity.
	EntityLinks []protocol.EntityLink
	// DeviceID is the device ID set in one of the files found in the storage of the device of the player. It
	// may be changed freely, so it should not be relied on for anything.
	DeviceID string
	// BuildPlatform is the build platform/device OS of the player that is about to be added, as it sent in
	// the Login packet when joining.
	BuildPlatform int32
}

// ID ...
func (*AddPlayer) ID() uint32 {
	return IDAddPlayer
}

// Marshal ...
func (pk *AddPlayer) Marshal(w *protocol.Writer) {
	w.UUID(&pk.UUID)
	w.String(&pk.Username)
	if w.ProtocolID() < protocol.ID534 {
		actorUid := int64(pk.EntityRuntimeID)
		w.Varint64(&actorUid)
	}
	w.Varuint64(&pk.EntityRuntimeID)
	w.String(&pk.PlatformChatID)
	w.Vec3(&pk.Position)
	w.Vec3(&pk.Velocity)
	w.Float32(&pk.Pitch)
	w.Float32(&pk.Yaw)
	w.Float32(&pk.HeadYaw)
	w.ItemInstance(&pk.HeldItem)
	if w.ProtocolID() >= protocol.ID503 {
		w.Varint32(&pk.GameType)
	}
	w.EntityMetadata(&pk.EntityMetadata)
	if w.ProtocolID() >= protocol.ID534 {
		w.Int64(&pk.EntityUniqueID)
		w.Uint8(&pk.PlayerPermissions)
		w.Uint8(&pk.CommandPermissions)
		protocol.SliceUint8Length(w, &pk.Layers)
	} else {
		// TODO: this shouldn't hardcoded
		flags := uint32(0)
		w.Varuint32(&flags)
		commandPermission := uint32(pk.CommandPermissions)
		w.Varuint32(&commandPermission)
		actionPermission := uint32(0)
		w.Varuint32(&actionPermission)
		playerPermission := uint32(pk.PlayerPermissions)
		w.Varuint32(&playerPermission)
		customStoredPermission := uint32(0)
		w.Varuint32(&customStoredPermission)
		w.Int64(&pk.EntityUniqueID)
	}
	protocol.Slice(w, &pk.EntityLinks)
	w.String(&pk.DeviceID)
	w.Int32(&pk.BuildPlatform)
}

// Unmarshal ...
func (pk *AddPlayer) Unmarshal(r *protocol.Reader) {
	r.UUID(&pk.UUID)
	r.String(&pk.Username)
	if r.ProtocolID() < protocol.ID534 {
		var actorUid int64
		r.Varint64(&actorUid)
	}
	r.Varuint64(&pk.EntityRuntimeID)
	r.String(&pk.PlatformChatID)
	r.Vec3(&pk.Position)
	r.Vec3(&pk.Velocity)
	r.Float32(&pk.Pitch)
	r.Float32(&pk.Yaw)
	r.Float32(&pk.HeadYaw)
	r.ItemInstance(&pk.HeldItem)
	if r.ProtocolID() >= protocol.ID503 {
		r.Varint32(&pk.GameType)
	}
	r.EntityMetadata(&pk.EntityMetadata)
	if r.ProtocolID() >= protocol.ID534 {
		r.Int64(&pk.EntityUniqueID)
		r.Uint8(&pk.PlayerPermissions)
		r.Uint8(&pk.CommandPermissions)
		protocol.SliceUint8Length(r, &pk.Layers)
	} else {
		var flags uint32
		r.Varuint32(&flags)

		var commandPermission uint32
		r.Varuint32(&commandPermission)
		pk.CommandPermissions = byte(commandPermission)

		var actionPermissions uint32
		r.Varuint32(&actionPermissions)

		var playerPermission uint32
		r.Varuint32(&playerPermission)
		pk.PlayerPermissions = byte(playerPermission)

		var customStoredPermission uint32
		r.Varuint32(&customStoredPermission)

		r.Int64(&pk.EntityUniqueID)

		pk.Layers = make([]protocol.AbilityLayer, 0)
	}
	protocol.Slice(r, &pk.EntityLinks)
	r.String(&pk.DeviceID)
	r.Int32(&pk.BuildPlatform)
}
