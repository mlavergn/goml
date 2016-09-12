// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	"testing"
)


func TestSigmoid(t *testing.T) {
	x := Sigmoid(1.0).(float64)
	exp := 1.581977
	if Round(x, 6) != exp {
		t.Errorf("Sigmoid of %f is not %f", x, exp)
	}
}
