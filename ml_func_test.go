// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	"testing"
)

func _equalsVector(t *testing.T, x Vector, y Vector) {
	if len(x) != len(y) {
		t.Errorf("row count x:%d != y:%d", len(x), len(y))
		return
	}

	for i, val := range x {
		if val != y[i] {
			t.Errorf("value at %d is %f != %f", i, val, y[i])
			return
		}
	}

}

func _equalsMatrix(t *testing.T, x Matrix, y Matrix) {
	if len(x) != len(y) {
		t.Errorf("rows %d vs expected %d", len(x), len(y))
		return
	}

	if len(x[0]) != len(y[0]) {
		t.Errorf("cols %d vs expected %d", len(x[0]), len(y[0]))
		return
	}

	for rowi, row := range x {
		for coli, v := range row {
			if v != y[rowi][coli] {
				t.Errorf("value at %d:%d is %f != %f", rowi, coli, v, y[rowi][coli])
				return
			}
		}
	}
}

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
	e := Matrix{{1, 5}, {2, 6}, {3, 7}, {4, 8}}
	_equalsMatrix(t, y, e)
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
		if x[0][i] != 0 {
			t.Errorf("index %f is not 0", i)
		}
	}
}

func TestOnes(t *testing.T) {
	x := Ones(1, 10)
	for i := 0; i < 10; i++ {
		if x[0][i] != 1 {
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
	e := Vector{1, 2, 3, 4, 5, 6}
	_equalsVector(t, y, e)
}

func TestReshape(t *testing.T) {
	x := Vector{1, 2, 3, 4, 5, 6}
	y := Reshape(x, 2, 3)
	e := Matrix{{1, 2, 3}, {4, 5, 6}}
	_equalsMatrix(t, y, e)
}

func TestJoinV(t *testing.T) {
	x := JoinV(Vector{1, 2, 3}, Vector{4, 5, 6}, Vector{7, 8, 9})
	y := Vector{1, 2, 3, 4, 5, 6, 7, 8, 9}
	_equalsVector(t, x, y)
}

func TestJoinM(t *testing.T) {
	x := JoinM(Matrix{{1, 2}, {3, 4}, {5, 6}}, Matrix{{10, 20}, {30, 40}, {50, 60}})
	y := Matrix{{1, 2, 10, 20}, {3, 4, 30, 40}, {5, 6, 50, 60}}
	_equalsMatrix(t, x, y)
}
