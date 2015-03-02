// Package drum is supposed to implement the decoding of .splice drum machine files.
// See golang-challenge.com/go-challenge1/ for more information
package drum

import (
	"bytes"
	"fmt"
)

// Pattern is the high level representation of the
// drum pattern contained in a .splice file.
// TODO: implement
type Pattern struct {
	version string
	tempo   int
	tracks  []track
}

type track struct {
	id    int
	name  string
	steps []bool
}

type step struct {
	active bool
}

// Default string representation of Pattern
func (p *Pattern) String() string {
	var buffer bytes.Buffer

	var template = "Saved with HW Version: %s\nTempo: %d"
	buffer.WriteString(fmt.Sprintf(template, p.version, p.tempo))
	for _, track := range p.tracks {
		buffer.WriteString(fmt.Sprintf("\n%v", track))
	}
	return buffer.String()
}

func (t *track) String() string {
	var template = "(%d) %s\t|%s%s%s%s|%s%s%s%s|%s%s%s%s|%s%s%s%s|"
	return fmt.Sprintf(template, t.id, t.name, t.steps)
}

func (s *step) String() string {
	if s.active {
		return "x"
	} else {
		return "-"
	}
}
