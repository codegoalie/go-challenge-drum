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

	p = &Pattern{version: "0.808-alpha", tempo: 120, tracks: []track{
		track{id: 0, name: "kick", steps: []step{
			step{active: true},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: false},
			step{active: false}}},
		track{id: 1, name: "snare", steps: []step{
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: false},
			step{active: false}}},
		track{id: 2, name: "clap", steps: []step{
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false}}},
		track{id: 3, name: "hh-open", steps: []step{
			step{active: false},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: true},
			step{active: false}}},
		track{id: 4, name: "hh-close", steps: []step{
			step{active: true},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: false},
			step{active: true}}},
		track{id: 5, name: "cowbell", steps: []step{
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: true},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false},
			step{active: false}}}}}
	return p, nil
}
