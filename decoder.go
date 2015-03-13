package drum

import (
	"bufio"
	"encoding/binary"
	"os"
	"strings"
)

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
func DecodeFile(path string) (*Pattern, error) {
	file, err := os.Open(path)
	if err != nil {
		return &(Pattern{}), err
	}
	reader := bufio.NewReader(file)

	var decoded decodedFile
	var pattern Pattern

	binary.Read(reader, binary.LittleEndian, &decoded)
	pattern.Version = strings.TrimRight(string(decoded.Version[:]), string([]byte{0x00}))
	pattern.Tempo = decoded.Tempo
	for reader.Buffered() > 0 {
		var id int32
		binary.Read(reader, binary.LittleEndian, &id)

		next, err := reader.Peek(1)
		var name []byte
		for err == nil && next[0] != 0 && next[0] != 1 {
			charByte, _ := reader.ReadByte()
			name = append(name, charByte)
			next, err = reader.Peek(1)
		}
		var decodedSteps [16]int8
		binary.Read(reader, binary.LittleEndian, &decodedSteps)
		convertedTrack := Track{Id: int(id), Name: string(name[1:])}
		for i, intStep := range decodedSteps {
			convertedTrack.Steps[i] = Step{Active: intStep != 0}
		}

		pattern.Tracks = append(pattern.Tracks, convertedTrack)
	}

	return &pattern, nil
}

type decodedFile struct {
	Header  [13]byte
	_       [1]byte
	Version [11]byte
	_       [21]byte
	Tempo   float32
}
