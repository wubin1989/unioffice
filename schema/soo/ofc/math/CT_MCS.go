// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_MCS struct {
	Mc []*CT_MC
}

func NewCT_MCS() *CT_MCS {
	ret := &CT_MCS{}
	return ret
}

func (m *CT_MCS) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	semc := xml.StartElement{Name: xml.Name{Local: "m:mc"}}
	for _, c := range m.Mc {
		e.EncodeElement(c, semc)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_MCS) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_MCS:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "mc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/math", Local: "mc"}:
				tmp := NewCT_MC()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Mc = append(m.Mc, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_MCS %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_MCS
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_MCS and its children
func (m *CT_MCS) Validate() error {
	return m.ValidateWithPath("CT_MCS")
}

// ValidateWithPath validates the CT_MCS and its children, prefixing error messages with path
func (m *CT_MCS) ValidateWithPath(path string) error {
	for i, v := range m.Mc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Mc[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
