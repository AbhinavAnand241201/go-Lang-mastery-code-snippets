package main

import (
	"testing"
)

type BinaryMessage struct {
	Type byte   // 1-byte message type
	Data []byte // Variable-length data
}

func FuzzBinaryParser(f *testing.F) {
	// Seed corpus with a valid binary message: [Type: 1 byte][Data: "Hello"]
	f.Add([]byte{1, 'H', 'e', 'l', 'l', 'o'})
	f.Add([]byte{2, 0xFF, 0x00, 0xAA})

	f.Fuzz(func(t *testing.T, data []byte) {
		if len(data) < 1 {
			return // Ignore empty or invalid input
		}

		msg := BinaryMessage{
			Type: data[0],
			Data: data[1:],
		}

		// Invariants: Type should be 1 or 2, and Data should not be too large
		if msg.Type != 1 && msg.Type != 2 {
			t.Errorf("Invalid message type: %d", msg.Type)
		}
		if len(msg.Data) > 1000 {
			t.Errorf("Data too large: %d bytes", len(msg.Data))
		}
	})
}
