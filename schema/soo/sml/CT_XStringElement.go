// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"
)

type CT_XStringElement struct {
	// Value
	VAttr string
}

func NewCT_XStringElement() *CT_XStringElement {
	ret := &CT_XStringElement{}
	return ret
}

func (m *CT_XStringElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "v"},
		Value: fmt.Sprintf("%v", m.VAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_XStringElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "v" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.VAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_XStringElement: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_XStringElement and its children
func (m *CT_XStringElement) Validate() error {
	return m.ValidateWithPath("CT_XStringElement")
}

// ValidateWithPath validates the CT_XStringElement and its children, prefixing error messages with path
func (m *CT_XStringElement) ValidateWithPath(path string) error {
	return nil
}
