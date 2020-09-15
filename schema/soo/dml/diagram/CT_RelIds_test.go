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

	"github.com/wubin1989/unioffice/schema/soo/dml/diagram"
)

func TestCT_RelIdsConstructor(t *testing.T) {
	v := diagram.NewCT_RelIds()
	if v == nil {
		t.Errorf("diagram.NewCT_RelIds must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_RelIds should validate: %s", err)
	}
}

func TestCT_RelIdsMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_RelIds()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_RelIds()
	xml.Unmarshal(buf, v2)
}
