// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
	"github.com/unidoc/unioffice/schema/soo/dml"
)

type EG_AxShared struct {
	AxId           *CT_UnsignedInt
	Scaling        *CT_Scaling
	Delete         *CT_Boolean
	AxPos          *CT_AxPos
	MajorGridlines *CT_ChartLines
	MinorGridlines *CT_ChartLines
	Title          *CT_Title
	NumFmt         *CT_NumFmt
	MajorTickMark  *CT_TickMark
	MinorTickMark  *CT_TickMark
	TickLblPos     *CT_TickLblPos
	SpPr           *dml.CT_ShapeProperties
	TxPr           *dml.CT_TextBody
	CrossAx        *CT_UnsignedInt
	Choice         *EG_AxSharedChoice
}

func NewEG_AxShared() *EG_AxShared {
	ret := &EG_AxShared{}
	ret.AxId = NewCT_UnsignedInt()
	ret.Scaling = NewCT_Scaling()
	ret.AxPos = NewCT_AxPos()
	ret.CrossAx = NewCT_UnsignedInt()
	return ret
}

func (m *EG_AxShared) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	seaxId := xml.StartElement{Name: xml.Name{Local: "c:axId"}}
	e.EncodeElement(m.AxId, seaxId)
	sescaling := xml.StartElement{Name: xml.Name{Local: "c:scaling"}}
	e.EncodeElement(m.Scaling, sescaling)
	if m.Delete != nil {
		sedelete := xml.StartElement{Name: xml.Name{Local: "c:delete"}}
		e.EncodeElement(m.Delete, sedelete)
	}
	seaxPos := xml.StartElement{Name: xml.Name{Local: "c:axPos"}}
	e.EncodeElement(m.AxPos, seaxPos)
	if m.MajorGridlines != nil {
		semajorGridlines := xml.StartElement{Name: xml.Name{Local: "c:majorGridlines"}}
		e.EncodeElement(m.MajorGridlines, semajorGridlines)
	}
	if m.MinorGridlines != nil {
		seminorGridlines := xml.StartElement{Name: xml.Name{Local: "c:minorGridlines"}}
		e.EncodeElement(m.MinorGridlines, seminorGridlines)
	}
	if m.Title != nil {
		setitle := xml.StartElement{Name: xml.Name{Local: "c:title"}}
		e.EncodeElement(m.Title, setitle)
	}
	if m.NumFmt != nil {
		senumFmt := xml.StartElement{Name: xml.Name{Local: "c:numFmt"}}
		e.EncodeElement(m.NumFmt, senumFmt)
	}
	if m.MajorTickMark != nil {
		semajorTickMark := xml.StartElement{Name: xml.Name{Local: "c:majorTickMark"}}
		e.EncodeElement(m.MajorTickMark, semajorTickMark)
	}
	if m.MinorTickMark != nil {
		seminorTickMark := xml.StartElement{Name: xml.Name{Local: "c:minorTickMark"}}
		e.EncodeElement(m.MinorTickMark, seminorTickMark)
	}
	if m.TickLblPos != nil {
		setickLblPos := xml.StartElement{Name: xml.Name{Local: "c:tickLblPos"}}
		e.EncodeElement(m.TickLblPos, setickLblPos)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "c:spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.TxPr != nil {
		setxPr := xml.StartElement{Name: xml.Name{Local: "c:txPr"}}
		e.EncodeElement(m.TxPr, setxPr)
	}
	secrossAx := xml.StartElement{Name: xml.Name{Local: "c:crossAx"}}
	e.EncodeElement(m.CrossAx, secrossAx)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	return nil
}

func (m *EG_AxShared) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.AxId = NewCT_UnsignedInt()
	m.Scaling = NewCT_Scaling()
	m.AxPos = NewCT_AxPos()
	m.CrossAx = NewCT_UnsignedInt()
lEG_AxShared:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "axId"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "axId"}:
				if err := d.DecodeElement(m.AxId, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "scaling"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "scaling"}:
				if err := d.DecodeElement(m.Scaling, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "delete"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "delete"}:
				m.Delete = NewCT_Boolean()
				if err := d.DecodeElement(m.Delete, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "axPos"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "axPos"}:
				if err := d.DecodeElement(m.AxPos, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "majorGridlines"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "majorGridlines"}:
				m.MajorGridlines = NewCT_ChartLines()
				if err := d.DecodeElement(m.MajorGridlines, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "minorGridlines"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "minorGridlines"}:
				m.MinorGridlines = NewCT_ChartLines()
				if err := d.DecodeElement(m.MinorGridlines, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "title"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "title"}:
				m.Title = NewCT_Title()
				if err := d.DecodeElement(m.Title, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "numFmt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "numFmt"}:
				m.NumFmt = NewCT_NumFmt()
				if err := d.DecodeElement(m.NumFmt, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "majorTickMark"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "majorTickMark"}:
				m.MajorTickMark = NewCT_TickMark()
				if err := d.DecodeElement(m.MajorTickMark, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "minorTickMark"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "minorTickMark"}:
				m.MinorTickMark = NewCT_TickMark()
				if err := d.DecodeElement(m.MinorTickMark, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "tickLblPos"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "tickLblPos"}:
				m.TickLblPos = NewCT_TickLblPos()
				if err := d.DecodeElement(m.TickLblPos, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "spPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "spPr"}:
				m.SpPr = dml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "txPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "txPr"}:
				m.TxPr = dml.NewCT_TextBody()
				if err := d.DecodeElement(m.TxPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "crossAx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "crossAx"}:
				if err := d.DecodeElement(m.CrossAx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "crosses"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "crosses"}:
				m.Choice = NewEG_AxSharedChoice()
				if err := d.DecodeElement(&m.Choice.Crosses, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "crossesAt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "crossesAt"}:
				m.Choice = NewEG_AxSharedChoice()
				if err := d.DecodeElement(&m.Choice.CrossesAt, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on EG_AxShared %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_AxShared
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_AxShared and its children
func (m *EG_AxShared) Validate() error {
	return m.ValidateWithPath("EG_AxShared")
}

// ValidateWithPath validates the EG_AxShared and its children, prefixing error messages with path
func (m *EG_AxShared) ValidateWithPath(path string) error {
	if err := m.AxId.ValidateWithPath(path + "/AxId"); err != nil {
		return err
	}
	if err := m.Scaling.ValidateWithPath(path + "/Scaling"); err != nil {
		return err
	}
	if m.Delete != nil {
		if err := m.Delete.ValidateWithPath(path + "/Delete"); err != nil {
			return err
		}
	}
	if err := m.AxPos.ValidateWithPath(path + "/AxPos"); err != nil {
		return err
	}
	if m.MajorGridlines != nil {
		if err := m.MajorGridlines.ValidateWithPath(path + "/MajorGridlines"); err != nil {
			return err
		}
	}
	if m.MinorGridlines != nil {
		if err := m.MinorGridlines.ValidateWithPath(path + "/MinorGridlines"); err != nil {
			return err
		}
	}
	if m.Title != nil {
		if err := m.Title.ValidateWithPath(path + "/Title"); err != nil {
			return err
		}
	}
	if m.NumFmt != nil {
		if err := m.NumFmt.ValidateWithPath(path + "/NumFmt"); err != nil {
			return err
		}
	}
	if m.MajorTickMark != nil {
		if err := m.MajorTickMark.ValidateWithPath(path + "/MajorTickMark"); err != nil {
			return err
		}
	}
	if m.MinorTickMark != nil {
		if err := m.MinorTickMark.ValidateWithPath(path + "/MinorTickMark"); err != nil {
			return err
		}
	}
	if m.TickLblPos != nil {
		if err := m.TickLblPos.ValidateWithPath(path + "/TickLblPos"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.TxPr != nil {
		if err := m.TxPr.ValidateWithPath(path + "/TxPr"); err != nil {
			return err
		}
	}
	if err := m.CrossAx.ValidateWithPath(path + "/CrossAx"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	return nil
}
