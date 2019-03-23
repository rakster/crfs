// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "testing"

func TestMakeDev(t *testing.T) {
	tests := []struct {
		major, minor int
		want         uint32
	}{
		{0, 0, 0},
		{0, 1, 1},
		{1, 0, 256},
		{1, 1, 257},
		{1, 2, 258},
		{100, 1, 25601},
		{255, 1, 65281},
		{256, 1, 65537},
		{257, 1, 65793},
		{1, 254, 510},
		{1, 255, 511},
		{1, 256, 1048832},
		{1, 257, 1048833},
	}
	for _, tt := range tests {
		got := makedev(uint32(tt.major), uint32(tt.minor))
		if got != tt.want {
			t.Errorf("makedev(%v, %v) = %v; want %v", tt.major, tt.minor, got, tt.want)
		}
	}

}
