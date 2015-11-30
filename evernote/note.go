package evernote

import (
	"encoding/base64"
	"encoding/xml"
	// "fmt"
)

// https://golang.org/pkg/encoding/xml/
// https://gist.github.com/evernotegists/6116886
// https://blog.evernote.com/tech/2013/08/08/evernote-export-format-enex/
// https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/07.1.md
type Note struct {
	XMLName   xml.Name   `xml:note`
	Title     string     `xml:"title"`
	Content   string     `xml:"content"`
	Created   string     `xml:"created"`
	Updated   string     `xml:"updated"`
	Tag       string     `xml:"tag"`
	Attrs     Attributes `xml:"note-attributes"`
	Resources []Resource `xml:"resource"`
}

// Simple String Attributes
// <note-attributes>
//     <latitude>33.88394692352314</latitude>
//     <longitude>-117.9191355110099</longitude>
//     <altitude>96</altitude>
//     <author>Brett Kelly</author>
// </note-attributes>
//
// <resource-attributes>
//     <file-name>snapshot-DAE9FC15-88E3-46CF-B744-DA9B1B56EB57.jpg</file-name>
// </resource-attributes>
type Attributes map[string]string

func (a *Attributes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	if *a == nil {
		*a = map[string]string{}
	}
	m := *a
	for t, err := d.Token(); err == nil; t, err = d.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			var s string
			err = d.DecodeElement(&s, &token)
			m[token.Name.Local] = s
		}
	}
	return nil
}

// xml:note>resource
type Resource struct {
	XMLName xml.Name   `xml:"resource"`
	Data    string     `xml:"data"` // <data encoding="base64">...</data>
	Mime    string     `xml:"mime"`
	Width   int        `xml:"width"`
	Height  int        `xml:"height"`
	Attrs   Attributes `xml:"resource-attributes"`
}

func (r *Resource) Filename() string {
	if name, ok := r.Attrs["file-name"]; ok {
		return name
	}
	return ""
}

func (r *Resource) Bytes() []byte {
	if b, err := base64.StdEncoding.DecodeString(r.Data); err == nil {
		return b
	}
	return nil
}
