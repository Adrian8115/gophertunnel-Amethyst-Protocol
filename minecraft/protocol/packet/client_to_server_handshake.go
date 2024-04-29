package packet

import (
	"github.com/Adrian8115/gophertunnel-Amethyst-Protocol/minecraft/protocol"
)

// ClientToServerHandshake is sent by the client in response to a ServerToClientHandshake packet sent by the
// server. It is the first encrypted packet in the login handshake and serves as a confirmation that
// encryption is correctly initialised client side.
type ClientToServerHandshake struct {
	// ClientToServerHandshake has no fields.
}

// ID ...
func (*ClientToServerHandshake) ID() uint32 {
	return IDClientToServerHandshake
}

func (*ClientToServerHandshake) Marshal(protocol.IO) {}
