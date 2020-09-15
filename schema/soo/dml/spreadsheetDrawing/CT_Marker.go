// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package spreadsheetDrawing

import (
	"encoding/xml"
	"fmt"

	"github.com/wubin1989/unioffice"
	"github.com/wubin1989/unioffice/schema/soo/dml"
)

type CT_Marker struct {
	Col    int32
	ColOff dml.ST_Coordinate
	Row    int32
	RowOff dml.ST_Coordinate
}

func NewCT_Marker() *CT_Marker {
	ret := &CT_Marker{}
	ret.Col = 0
	ret.Row = 0
	return ret
}

func (m *CT_Marker) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secol := xml.StartElement{Name: xml.Name{Local: "xdr:col"}}
	e.EncodeElement(m.Col, secol)
	secolOff := xml.StartElement{Name: xml.Name{Local: "xdr:colOff"}}
	e.EncodeElement(m.ColOff, secolOff)
	serow := xml.StartElement{Name: xml.Name{Local: "xdr:row"}}
	e.EncodeElement(m.Row, serow)
	serowOff := xml.StartElement{Name: xml.Name{Local: "xdr:rowOff"}}
	e.EncodeElement(m.RowOff, serowOff)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Marker) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Col = 0
	m.Row = 0
lCT_Marker:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "col"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "col"}:
				if err := d.DecodeElement(&m.Col, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "colOff"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "colOff"}:
				if err := d.DecodeElement(&m.ColOff, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "row"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "row"}:
				if err := d.DecodeElement(&m.Row, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "rowOff"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "rowOff"}:
				if err := d.DecodeElement(&m.RowOff, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_Marker %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Marker
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Marker and its children
func (m *CT_Marker) Validate() error {
	return m.ValidateWithPath("CT_Marker")
}

// ValidateWithPath validates the CT_Marker and its children, prefixing error messages with path
func (m *CT_Marker) ValidateWithPath(path string) error {
	if m.Col < 0 {
		return fmt.Errorf("%s/m.Col must be >= 0 (have %v)", path, m.Col)
	}
	if err := m.ColOff.ValidateWithPath(path + "/ColOff"); err != nil {
		return err
	}
	if m.Row < 0 {
		return fmt.Errorf("%s/m.Row must be >= 0 (have %v)", path, m.Row)
	}
	if err := m.RowOff.ValidateWithPath(path + "/RowOff"); err != nil {
		return err
	}
	return nil
}
