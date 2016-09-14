// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	. "golog"
	"reflect"
)

type Matrix [][]float64
type Vector []float64

// type Scalar float64
type Data interface{}

type ArgBitmask uint16

const (
	ARG1_MATRIX ArgBitmask = 1 << iota
	ARG1_VECTOR
	ARG1_SCALAR
	ARG2_MATRIX
	ARG2_VECTOR
	ARG2_SCALAR
)

//
// Generate a bitmask based on the argument data types
//
func _argBitmask(arg1 Data, arg2 Data) (flags ArgBitmask) {
	switch arg1.(type) {
	case Matrix, [][]float64:
		flags |= ARG1_MATRIX
	case Vector, []float64:
		flags |= ARG1_VECTOR
	case float64, float32, int:
		flags |= ARG1_SCALAR
	default:
		LogErrorf("Unhandled argument type: %s", reflect.TypeOf(arg1))
	}

	switch arg2.(type) {
	case Matrix, [][]float64:
		flags |= ARG2_MATRIX
	case Vector, []float64:
		flags |= ARG2_VECTOR
	case float64, float32, int:
		flags |= ARG2_SCALAR
	default:
		LogErrorf("Unhandled argument type: %s", reflect.TypeOf(arg2))
	}

	return flags
}
