// Copyright 2012, Bryan Matsuo. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simplesettings

/*  Filename:    simplesettings_test.go
 *  Author:      Bryan Matsuo <bryan.matsuo [at] gmail.com>
 *  Created:     2012-11-25 16:44:51.995633 -0800 PST
 *  Description: Main test file for go-simplesettings
 */

import (
    "testing"
)

func TestSimplesettingsBottom(t *testing.T) {
	settings := New("test")
	if s, err := settings.String(); err != nil {
		t.Errorf("String() error: %v", err)
	} else if s != "test" {
		t.Errorf("String() mismatch: %v", s)
	}
	if p, err := settings.Bytes(); err != nil {
		t.Errorf("Bytes() error: %v", err)
	} else if string(p) != "test" {
		t.Errorf("Bytes() mismatch: %s", p)
	}
}

func TestSimplesettings(t *testing.T) {
	settings := New(Map{
		"testkey": "testvalue",
	})
	if s, err := settings.Get("testkey").String(); err != nil {
		t.Errorf("Get() error: %v", err)
	} else if s != "testvalue" {
		t.Errorf("Get() mismatch: %v", s)
	}
}
