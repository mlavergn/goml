// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	"testing"
)

func TestAddVV(t *testing.T) {
	x := AddVV(Vector{1, 2, 3}, Vector{10, 100, 1000})
	exp := Vector{11, 102, 1003}
	_equalsVector(t, x, exp)
}

func TestSubVV(t *testing.T) {
	x := SubVV(Vector{1, 2, 3}, Vector{10, 100, 1000})
	exp := Vector{-9, -98, -997}
	_equalsVector(t, x, exp)
}

func TestMulVX(t *testing.T) {
	x := MulVX(Vector{1, 2, 3}, 10.0)
	exp := Vector{10, 20, 30}
	_equalsVector(t, x, exp)
}

func TestDivVX(t *testing.T) {
	x := DivVX(Vector{1, 2, 3}, 2.0)
	exp := Vector{0.5, 1.0, 1.5}
	_equalsVector(t, x, exp)
}

func TestAddMM(t *testing.T) {
	x := AddMM(Matrix{{1, 2}, {3, 4}, {5, 6}}, Matrix{{6, 5}, {4, 3}, {2, 1}})
	exp := Matrix{{7, 7}, {7, 7}, {7, 7}}
	_equalsMatrix(t, x, exp)
}

func TestSubMM(t *testing.T) {
	x := SubMM(Matrix{{1, 2}, {3, 4}, {5, 6}}, Matrix{{6, 5}, {4, 3}, {2, 1}})
	exp := Matrix{{-5, -3}, {-1, 1}, {3, 5}}
	_equalsMatrix(t, x, exp)
}

func TestMulMM(t *testing.T) {
	x := MulMM(Matrix{{1, 2}, {3, 4}, {5, 6}}, Matrix{{10}, {100}})
	exp := Matrix{{210}, {430}, {650}}
	_equalsMatrix(t, x, exp)
}

func TestMulVM(t *testing.T) {
	x := MulVM(Vector{1, 2}, Matrix{{1, 2, 3}, {4, 5, 6}})
	exp := Vector{9, 12, 15}
	_equalsVector(t, x, exp)
}

func TestMulXM(t *testing.T) {
	x := MulXM(2.0, Matrix{{1, 2}, {3, 4}, {5, 6}})
	exp := Matrix{{2, 4}, {6, 8}, {10, 12}}
	_equalsMatrix(t, x, exp)
}

func TestDivMX(t *testing.T) {
	x := DivMX(Matrix{{1, 2}, {3, 4}, {5, 6}}, 2.0)
	exp := Matrix{{0.5, 1.0}, {1.5, 2.0}, {2.5, 3.0}}
	_equalsMatrix(t, x, exp)
}
