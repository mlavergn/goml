// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	"testing"
)

func TestNewVector(t *testing.T) {
	x := NewVector(10)
	if len(x) != 10 {
		t.Errorf("rows %f vs expected %f", len(x), 10)
		return
	}
}

func TestNewMatrix(t *testing.T) {
	x := NewMatrix(10, 20)
	if len(x) != 10 {
		t.Errorf("rows %f vs expected %f", len(x), 10)
		return
	}

	if len(x[0]) != 20 {
		t.Errorf("cols %f vs expected %f", len(x[0]), 10)
	}
}

func TestSize(t *testing.T) {
	x := Matrix{{1, 2, 3}, {4, 5, 6}}
	rows, cols := Size(x)

	if rows != 2 {
		t.Errorf("rows %f vs expected %f", len(x), 2)
		return
	}
	if cols != 3 {
		t.Errorf("cols %f vs expected %f", len(x), 3)
		return
	}
}

func TestTranspose(t *testing.T) {
	x := Matrix{{1, 2, 3, 4}, {5, 6, 7, 8}}
	y := Transpose(x).(Matrix)
	exp := Matrix{{1, 5}, {2, 6}, {3, 7}, {4, 8}}
	if !Equal(y, exp) {
		t.Errorf("%v != %v", y, exp)
	}
}

func TestEye(t *testing.T) {
	x := Eye(3)
	for i := 0; i < 3; i++ {
		if x[i][i] != 1 {
			t.Errorf("non-one %f vs expected %f", x[i][i], 1)
			break
		}
	}
}

func TestZeros(t *testing.T) {
	x := Zeros(1, 10)
	for i := 0; i < 10; i++ {
		if x.(Vector)[i] != 0 {
			t.Errorf("index %f is not 0", i)
		}
	}
}

func TestOnes(t *testing.T) {
	x := Ones(1, 10)
	for i := 0; i < 10; i++ {
		if x.(Vector)[i] != 1 {
			t.Errorf("index %f is not 1", i)
		}
	}
}

func TestSeq(t *testing.T) {
	x := Seq(1, 10)
	seq := 1.0
	for i := 0; i < 10; i++ {
		if x[0][i] != seq {
			t.Errorf("index %f is not 1", i)
		}
		seq += 1.0
	}
}

func TestRand(t *testing.T) {
	x := Rand(1, 10)
	for i := 0; i < 10; i++ {
		if x[0][i] < 0.0 || x[0][i] > 1.0 {
			t.Errorf("index %f value %f is not between 0.0 and 1.0", i, x[0][i])
		}
	}
}

func TestUnroll(t *testing.T) {
	x := Matrix{{1, 2, 3}, {4, 5, 6}}
	y := Unroll(x)
	exp := Vector{1, 2, 3, 4, 5, 6}
	if !Equal(y, exp) {
		t.Errorf("%v != %v", y, exp)
	}
}

func TestReshape(t *testing.T) {
	x := Vector{1, 2, 3, 4, 5, 6}
	y := Reshape(x, 2, 3)
	exp := Matrix{{1, 2, 3}, {4, 5, 6}}
	if !Equal(y, exp) {
		t.Errorf("%v != %v", x, exp)
	}
}

func TestJoinV(t *testing.T) {
	x := Join(Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9})
	y := Vector{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if !Equal(x.(Vector), y) {
		t.Errorf("%v != %v", x, y)
	}
}

func TestJoinM(t *testing.T) {
	x := Join(Matrix{{1, 2}, {3, 4}, {5, 6}}, Matrix{{10, 20}, {30, 40}, {50, 60}})
	y := Matrix{{1, 2, 10, 20}, {3, 4, 30, 40}, {5, 6, 50, 60}}
	if !Equal(x.(Matrix), y) {
		t.Errorf("%v != %v", x, y)
	}
}
