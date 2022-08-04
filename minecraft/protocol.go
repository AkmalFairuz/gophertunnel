package minecraft

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// Protocol represents the Minecraft protocol used to communicate over network. It comprises a unique set of packets
// that may be changed in any version.
// Protocol specifically handles the conversion of packets between the most recent protocol (as in the
// minecraft/protocol package) and the protocol as specified in Protocol.
type Protocol interface {
	// ID returns the unique ID of the Protocol. It generally goes up for every new Minecraft version released.
	ID() int32
	// Ver returns the Minecraft version associated with this Protocol, such as "1.18.10".
	Ver() string
	// Packets returns a packet.Pool with all packets registered for this Protocol. It is used to lookup packets by a
	// packet ID.
	Packets() packet.Pool
	// ConvertToLatest converts a packet.Packet obtained from the other end of a Conn to a slice of packet.Packets from
	// the latest protocol. Any packet.Packet implementation in the packet.Pool obtained through a call to Packets that
	// is not identical to the most recent version of that packet.Packet must be converted to the most recent version of
	// that packet adequately in this function. ConvertToLatest returns pk if the packet.Packet was unchanged in this
	// version compared to the latest. Note that packets must also be converted if only their ID changes.
	ConvertToLatest(pk packet.Packet, conn *Conn) []packet.Packet
	// ConvertFromLatest converts a packet.Packet of the most recent Protocol to a slice of packet.Packets of this
	// specific Protocol. ConvertFromLatest must be synonymous to ConvertToLatest, in that it should convert any
	// packet.Packet to the correct one from the packet.Pool returned through a call to Packets if its payload or ID was
	// changed in this Protocol compared to the latest one.
	ConvertFromLatest(pk packet.Packet, conn *Conn) []packet.Packet
}

// Proto is the default Protocol implementation. It returns the current protocol, version and packet pool and does not
// convert any packets, as they are already of the right type.
type Proto struct {
	Id int32
}

func (p Proto) ID() int32                                                 { return p.Id }
func (p Proto) Ver() string                                               { return protocol.CurrentVersion }
func (p Proto) Packets() packet.Pool                                      { return packet.NewPool() }
func (p Proto) ConvertToLatest(pk packet.Packet, _ *Conn) []packet.Packet { return []packet.Packet{pk} }
func (p Proto) ConvertFromLatest(pk packet.Packet, _ *Conn) []packet.Packet {
	return []packet.Packet{pk}
}
