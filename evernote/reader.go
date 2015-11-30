package evernote

import (
	"encoding/xml"
	"io"
)

type WalkFunc func(note *Note) error

func Parse(r io.Reader, fn WalkFunc) {
	decoder := xml.NewDecoder(r)
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			if token.Name.Local == "note" {
				note := &Note{}
				if e2 := decoder.DecodeElement(note, &token); e2 != nil {
					break
				}
				if e3 := fn(note); e3 != nil {
					break
				}
			}
		}
	}
}
