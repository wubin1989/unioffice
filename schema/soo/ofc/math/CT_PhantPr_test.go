// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/wubin1989/unioffice/schema/soo/ofc/math"
)

func TestCT_PhantPrConstructor(t *testing.T) {
	v := math.NewCT_PhantPr()
	if v == nil {
		t.Errorf("math.NewCT_PhantPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_PhantPr should validate: %s", err)
	}
}

func TestCT_PhantPrMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_PhantPr()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_PhantPr()
	xml.Unmarshal(buf, v2)
}
