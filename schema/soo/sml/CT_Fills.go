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
	"strconv"

	"github.com/unidoc/unioffice"
)

type CT_Fills struct {
	// Fill Count
	CountAttr *uint32
	// Fill
	Fill []*CT_Fill
}

func NewCT_Fills() *CT_Fills {
	ret := &CT_Fills{}
	return ret
}

func (m *CT_Fills) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.Fill != nil {
		sefill := xml.StartElement{Name: xml.Name{Local: "ma:fill"}}
		for _, c := range m.Fill {
			e.EncodeElement(c, sefill)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Fills) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_Fills:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "fill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "fill"}:
				tmp := NewCT_Fill()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Fill = append(m.Fill, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_Fills %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Fills
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Fills and its children
func (m *CT_Fills) Validate() error {
	return m.ValidateWithPath("CT_Fills")
}

// ValidateWithPath validates the CT_Fills and its children, prefixing error messages with path
func (m *CT_Fills) ValidateWithPath(path string) error {
	for i, v := range m.Fill {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Fill[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
