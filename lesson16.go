func FuzzPNGDecode(f *testing.F) {
	// Seed with a small PNG (hex-encoded)
	seedPNG := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	f.Add(seedPNG)

	f.Fuzz(func(t *testing.T, data []byte) {
		_, err := png.Decode(bytes.NewReader(data))
		if err != nil {
			return // Skip invalid PNGs
		}
		// If we get here, the fuzzer found a valid PNG!
	})
}

// fuzzing a png decoder

// fuzzing a json api handler 

type Request struct {
    Action string `json:"action"`
    ID     int    `json:"id"`
}


// this  `json:"id"`  means json me ye id ko point karega ...action ko point karega similarly 
func FuzzJSONHandler(f *testing.F) {
    // Seed corpus
    f.Add([]byte(`{"action":"create","id":1}`))
    f.Add([]byte(`{"action":"delete","id":0}`))

    f.Fuzz(func(t *testing.T, data []byte) {
        var req Request
        if err := json.Unmarshal(data, &req); err != nil {
            return // Skip invalid JSON
        }

        // Test business logic
        if req.ID < 0 {
            t.Errorf("Invalid ID: %d", req.ID)
        }
        if req.Action != "create" && req.Action != "delete" {
            t.Errorf("Unknown action: %s", req.Action)
        }
    })




	go test -fuzz=. -fuzztime=20s


	this command runs all the fuzz tests






	// this is for the fuzzing of a network packet decoder ...
	// this is a network packet decoder
	so it has three things one is version 1 byte, payload length 1 byte  and then payload which is N bytes 

	type Packet struct {
		Version uint8
		Payload []byte
	}
	
	func FuzzDecodePacket(f *testing.F) {
		// Seed corpus: [Version][PayloadLen][Payload...]
		f.Add([]byte{1, 3, 'A', 'B', 'C'})  // Valid packet
		f.Add([]byte{0, 0})                  // Minimal packet
	// here the seed corpus shows 1 verisn , 3 length and the payload is A,B,C ....
		f.Fuzz(func(t *testing.T, data []byte) {
			if len(data) < 2 {
				return // Too short
			}
			p := Packet{
				Version: data[0],
				Payload: data[2:], // Assume data[1] = length
			}
			if p.Version != 1 {
				t.Errorf("Unsupported version: %d", p.Version)
			}
		})
	}