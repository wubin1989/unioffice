// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/wubin1989/unioffice/schema/soo/dml"
)

func TestThemeConstructor(t *testing.T) {
	v := dml.NewTheme()
	if v == nil {
		t.Errorf("dml.NewTheme must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.Theme should validate: %s", err)
	}
}

func TestThemeMarshalUnmarshal(t *testing.T) {
	v := dml.NewTheme()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewTheme()
	xml.Unmarshal(buf, v2)
}
