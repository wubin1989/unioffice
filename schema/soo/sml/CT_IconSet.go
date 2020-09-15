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

type CT_IconSet struct {
	// Icon Set
	IconSetAttr ST_IconSetType
	// Show Value
	ShowValueAttr *bool
	// Percent
	PercentAttr *bool
	// Reverse Icons
	ReverseAttr *bool
	// Conditional Formatting Object
	Cfvo []*CT_Cfvo
}

func NewCT_IconSet() *CT_IconSet {
	ret := &CT_IconSet{}
	return ret
}

func (m *CT_IconSet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IconSetAttr != ST_IconSetTypeUnset {
		attr, err := m.IconSetAttr.MarshalXMLAttr(xml.Name{Local: "iconSet"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ShowValueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "showValue"},
			Value: fmt.Sprintf("%d", b2i(*m.ShowValueAttr))})
	}
	if m.PercentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "percent"},
			Value: fmt.Sprintf("%d", b2i(*m.PercentAttr))})
	}
	if m.ReverseAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "reverse"},
			Value: fmt.Sprintf("%d", b2i(*m.ReverseAttr))})
	}
	e.EncodeToken(start)
	secfvo := xml.StartElement{Name: xml.Name{Local: "ma:cfvo"}}
	for _, c := range m.Cfvo {
		e.EncodeElement(c, secfvo)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_IconSet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "iconSet" {
			m.IconSetAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "showValue" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowValueAttr = &parsed
			continue
		}
		if attr.Name.Local == "percent" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PercentAttr = &parsed
			continue
		}
		if attr.Name.Local == "reverse" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ReverseAttr = &parsed
			continue
		}
	}
lCT_IconSet:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cfvo"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "cfvo"}:
				tmp := NewCT_Cfvo()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cfvo = append(m.Cfvo, tmp)
			default:
				unioffice.Log("skipping unsupported element on CT_IconSet %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_IconSet
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_IconSet and its children
func (m *CT_IconSet) Validate() error {
	return m.ValidateWithPath("CT_IconSet")
}

// ValidateWithPath validates the CT_IconSet and its children, prefixing error messages with path
func (m *CT_IconSet) ValidateWithPath(path string) error {
	if err := m.IconSetAttr.ValidateWithPath(path + "/IconSetAttr"); err != nil {
		return err
	}
	for i, v := range m.Cfvo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cfvo[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
