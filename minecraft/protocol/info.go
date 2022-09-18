package protocol

var AcceptedProtocols = []int32{
	ID545,
	ID544,
	ID534,
	ID527,
	ID503,
	ID486,
}

const (
	// CurrentProtocol is the current protocol version for the version below.
	CurrentProtocol = 545
	// CurrentVersion is the current version of Minecraft as supported by the `packet` package.
	CurrentVersion = "1.19.21"
	// ID545 is the protocol version for 1.19.21.
	ID545 = 545
	// ID544 is the protocol version for 1.19.20.
	ID544 = 544
	// ID534 is the protocol version for 1.19.10.
	ID534 = 534
	// ID527 is the protocol version for 1.19.0.
	ID527 = 527
	// ID503 is the protocol version for 1.18.30.
	ID503 = 503
	// ID486 is the protocol version for 1.18.10.
	ID486 = 486
)
