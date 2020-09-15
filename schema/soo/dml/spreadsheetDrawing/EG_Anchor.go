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

	"github.com/wubin1989/unioffice"
)

type EG_Anchor struct {
	TwoCellAnchor  *CT_TwoCellAnchor
	OneCellAnchor  *CT_OneCellAnchor
	AbsoluteAnchor *CT_AbsoluteAnchor
}

func NewEG_Anchor() *EG_Anchor {
	ret := &EG_Anchor{}
	return ret
}

func (m *EG_Anchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TwoCellAnchor != nil {
		setwoCellAnchor := xml.StartElement{Name: xml.Name{Local: "xdr:twoCellAnchor"}}
		e.EncodeElement(m.TwoCellAnchor, setwoCellAnchor)
	}
	if m.OneCellAnchor != nil {
		seoneCellAnchor := xml.StartElement{Name: xml.Name{Local: "xdr:oneCellAnchor"}}
		e.EncodeElement(m.OneCellAnchor, seoneCellAnchor)
	}
	if m.AbsoluteAnchor != nil {
		seabsoluteAnchor := xml.StartElement{Name: xml.Name{Local: "xdr:absoluteAnchor"}}
		e.EncodeElement(m.AbsoluteAnchor, seabsoluteAnchor)
	}
	return nil
}

func (m *EG_Anchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_Anchor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "twoCellAnchor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "twoCellAnchor"}:
				m.TwoCellAnchor = NewCT_TwoCellAnchor()
				if err := d.DecodeElement(m.TwoCellAnchor, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "oneCellAnchor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "oneCellAnchor"}:
				m.OneCellAnchor = NewCT_OneCellAnchor()
				if err := d.DecodeElement(m.OneCellAnchor, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "absoluteAnchor"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "absoluteAnchor"}:
				m.AbsoluteAnchor = NewCT_AbsoluteAnchor()
				if err := d.DecodeElement(m.AbsoluteAnchor, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on EG_Anchor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_Anchor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_Anchor and its children
func (m *EG_Anchor) Validate() error {
	return m.ValidateWithPath("EG_Anchor")
}

// ValidateWithPath validates the EG_Anchor and its children, prefixing error messages with path
func (m *EG_Anchor) ValidateWithPath(path string) error {
	if m.TwoCellAnchor != nil {
		if err := m.TwoCellAnchor.ValidateWithPath(path + "/TwoCellAnchor"); err != nil {
			return err
		}
	}
	if m.OneCellAnchor != nil {
		if err := m.OneCellAnchor.ValidateWithPath(path + "/OneCellAnchor"); err != nil {
			return err
		}
	}
	if m.AbsoluteAnchor != nil {
		if err := m.AbsoluteAnchor.ValidateWithPath(path + "/AbsoluteAnchor"); err != nil {
			return err
		}
	}
	return nil
}
