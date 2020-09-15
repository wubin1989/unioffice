// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram

import (
	"encoding/xml"
	"fmt"

	"github.com/wubin1989/unioffice"
	"github.com/wubin1989/unioffice/schema/soo/dml"
)

type CT_Colors struct {
	MethAttr       ST_ClrAppMethod
	HueDirAttr     ST_HueDir
	EG_ColorChoice []*dml.EG_ColorChoice
}

func NewCT_Colors() *CT_Colors {
	ret := &CT_Colors{}
	return ret
}

func (m *CT_Colors) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.MethAttr != ST_ClrAppMethodUnset {
		attr, err := m.MethAttr.MarshalXMLAttr(xml.Name{Local: "meth"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HueDirAttr != ST_HueDirUnset {
		attr, err := m.HueDirAttr.MarshalXMLAttr(xml.Name{Local: "hueDir"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.EG_ColorChoice != nil {
		for _, c := range m.EG_ColorChoice {
			c.MarshalXML(e, xml.StartElement{})
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Colors) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "meth" {
			m.MethAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "hueDir" {
			m.HueDirAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_Colors:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			default:
				unioffice.Log("skipping unsupported element on CT_Colors %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Colors
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Colors and its children
func (m *CT_Colors) Validate() error {
	return m.ValidateWithPath("CT_Colors")
}

// ValidateWithPath validates the CT_Colors and its children, prefixing error messages with path
func (m *CT_Colors) ValidateWithPath(path string) error {
	if err := m.MethAttr.ValidateWithPath(path + "/MethAttr"); err != nil {
		return err
	}
	if err := m.HueDirAttr.ValidateWithPath(path + "/HueDirAttr"); err != nil {
		return err
	}
	for i, v := range m.EG_ColorChoice {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/EG_ColorChoice[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
