// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package diagram_test

import (
	"encoding/xml"
	"testing"

	"github.com/unidoc/unioffice/schema/soo/dml/diagram"
)

func TestCT_StyleLabelConstructor(t *testing.T) {
	v := diagram.NewCT_StyleLabel()
	if v == nil {
		t.Errorf("diagram.NewCT_StyleLabel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_StyleLabel should validate: %s", err)
	}
}

func TestCT_StyleLabelMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_StyleLabel()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_StyleLabel()
	xml.Unmarshal(buf, v2)
}