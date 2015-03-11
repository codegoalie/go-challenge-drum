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
	steps []step
}

type step struct {
	active bool
}

// Default string representation of Pattern
func (p Pattern) String() string {
	var buffer bytes.Buffer

	var template = "Saved with HW Version: %s\nTempo: %d"
	buffer.WriteString(fmt.Sprintf(template, p.version, p.tempo))
	for _, track := range p.tracks {
		buffer.WriteString(fmt.Sprintf("\n%v", track))
	}
	buffer.WriteString("\n")
	return buffer.String()
}

func (t track) String() string {
	var buffer bytes.Buffer
	var template = "(%d) %s\t"
	buffer.WriteString(fmt.Sprintf(template, t.id, t.name))
	for i, step := range t.steps {
		if i%4 == 0 {
			buffer.WriteString("|")
		}
		buffer.WriteString(fmt.Sprintf("%v", step))
	}
	buffer.WriteString("|")
	return buffer.String()
}

func (s step) String() string {
	if s.active {
		return "x"
	} else {
		return "-"
	}
}
