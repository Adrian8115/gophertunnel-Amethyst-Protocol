package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

// Amethyst is an extra custom packet that is used by the Amethyst Protocol for extending the bedrock protocol
type Amethyst struct {
	// Placeholder data
	Data byte
}

// ID ...
func (*Amethyst) ID() uint32 {
	return IDAmethyst
}

func (pk *Amethyst) Marshal(io protocol.IO) {
	io.Uint8(&pk.Data)
}
