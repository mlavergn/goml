// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	"testing"
)

func TestAddVV(t *testing.T) {
	x := Add(Vector{1, 2, 3}, Vector{10, 100, 1000})
	exp := Vector{11, 102, 1003}
	if !Equal(x, exp) {
		t.Errorf("%v != %v", x, exp)
	}
}

func TestSubVV(t *testing.T) {
	x := Sub(Vector{1, 2, 3}, Vector{10, 100, 1000})
	exp := Vector{-9, -98, -997}
	if !Equal(x, exp) {
		t.Errorf("%v != %v", x, exp)
	}
}

func TestMulVS(t *testing.T) {
	x := Mul(10.0, Vector{1, 2, 3})
	exp := Vector{10, 20, 30}
	if !Equal(x, exp) {
		t.Errorf("%v != %v", x, exp)
	}
}

func TestDivVS(t *testing.T) {
	x := Div(Vector{1, 2, 3}, 2.0)
	exp := Vector{0.5, 1.0, 1.5}
	if !Equal(x, exp) {
		t.Errorf("%v != %v", x, exp)
	}
}

func TestAddMM(t *testing.T) {
	x := Add(Matrix{{1, 2}, {3, 4}, {5, 6}}, Matrix{{6, 5}, {4, 3}, {2, 1}})
	exp := Matrix{{7, 7}, {7, 7}, {7, 7}}
	if !Equal(x, exp) {
		t.Errorf("%v != %v", x, exp)
	}
}

func TestSubMM(t *testing.T) {
	x := Sub(Matrix{{1, 2}, {3, 4}, {5, 6}}, Matrix{{6, 5}, {4, 3}, {2, 1}})
	exp := Matrix{{-5, -3}, {-1, 1}, {3, 5}}
	if !Equal(x, exp) {
		t.Errorf("%v != %v", x, exp)
	}
}

func TestMulMM(t *testing.T) {
	x := Mul(Matrix{{1, 2}, {3, 4}, {5, 6}}, Matrix{{10}, {100}})
	exp := Matrix{{210}, {430}, {650}}
	if !Equal(x, exp) {
		t.Errorf("%v != %v", x, exp)
	}
}

func TestMulVM(t *testing.T) {
	x := Mul(Vector{1, 2}, Matrix{{1, 2, 3}, {4, 5, 6}})
	exp := Vector{9, 12, 15}
	if !Equal(x, exp) {
		t.Errorf("%v != %v", x, exp)
	}
}

func TestMulSM(t *testing.T) {
	x := Mul(2.0, Matrix{{1, 2}, {3, 4}, {5, 6}})
	exp := Matrix{{2, 4}, {6, 8}, {10, 12}}
	if !Equal(x, exp) {
		t.Errorf("%v != %v", x, exp)
	}
}

func TestDivMS(t *testing.T) {
	x := Div(Matrix{{1, 2}, {3, 4}, {5, 6}}, 2.0)
	exp := Matrix{{0.5, 1.0}, {1.5, 2.0}, {2.5, 3.0}}
	if !Equal(x, exp) {
		t.Errorf("%v != %v", x, exp)
	}
}
