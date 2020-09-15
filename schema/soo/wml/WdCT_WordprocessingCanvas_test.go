// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package wml_test

import (
	"encoding/xml"
	"testing"

	"github.com/wubin1989/unioffice/schema/soo/wml"
)

func TestWdCT_WordprocessingCanvasConstructor(t *testing.T) {
	v := wml.NewWdCT_WordprocessingCanvas()
	if v == nil {
		t.Errorf("wml.NewWdCT_WordprocessingCanvas must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WordprocessingCanvas should validate: %s", err)
	}
}

func TestWdCT_WordprocessingCanvasMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WordprocessingCanvas()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WordprocessingCanvas()
	xml.Unmarshal(buf, v2)
}
