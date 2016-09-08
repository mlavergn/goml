// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	"testing"
)

func TestRandomWholeFloatSet(t *testing.T) {
	x := RandomWholeFloatSet(1000, 2.0, 99.0)

	n := len(x)
	if n != 1000 {
		t.Errorf("%f vs expected %f", n, 1000)
	}

	for _, v := range x {
		if v < 2.0 || v > 99.0 {
			t.Errorf("%f vs expected range %f:%f", v, 2.0, 99.0)
			break
		}
	}
}

func TestRandomFloatSet(t *testing.T) {
	x := RandomFloatSet(1000, 2.0, 99.0)

	n := len(x)
	if n != 1000 {
		t.Errorf("%f vs expected %f", n, 1000)
	}

	for _, v := range x {
		if v < 2.0 || v > 99.0 {
			t.Errorf("%f vs expected range %f:%f", v, 2.0, 99.0)
			break
		}
	}
}

func TestRandomTrainingTestingSets(t *testing.T) {
	in := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	x, y := RandomTrainTestSets(in, 0.2)
	if len(x) != 4 || len(y) != 14 {
		t.Errorf("lens %f : %f vs expected %f : %f", len(x), len(y), 4, 14)
	}
}

func TestRandomTrainValidationTestSets(t *testing.T) {
	in := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	x, y, z := RandomTrainValidationTestSets(in, 0.2, 0.2)
	if len(x) != 12 || len(y) != 4 || len(z) != 4 {
		t.Errorf("lens %d : %d : %d vs expected %d : %d : %d", len(x), len(y), len(z), 12, 4, 4)
	}
}
