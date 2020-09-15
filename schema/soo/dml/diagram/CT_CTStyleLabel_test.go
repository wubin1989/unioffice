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

func TestCT_CTStyleLabelConstructor(t *testing.T) {
	v := diagram.NewCT_CTStyleLabel()
	if v == nil {
		t.Errorf("diagram.NewCT_CTStyleLabel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_CTStyleLabel should validate: %s", err)
	}
}

func TestCT_CTStyleLabelMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_CTStyleLabel()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_CTStyleLabel()
	xml.Unmarshal(buf, v2)
}