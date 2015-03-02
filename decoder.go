package drum

import (
	"bufio"
	"encoding/binary"
	"os"
)

// DecodeFile decodes the drum machine file found at the provided path
// and returns a pointer to a parsed pattern which is the entry point to the
// rest of the data.
// TODO: implement
func DecodeFile(path string) (*Pattern, error) {
	file, err := os.Open(path)
	if err != nil {
		return &Pattern{}, err
	}
	reader := bufio.NewReader(file)

	p := &Pattern{}
	binary.Read(reader, binary.LittleEndian, &p)
	return p, nil
}
