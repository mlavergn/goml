// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	. "golog"
	"reflect"
)

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
    LogWarnf("unhandled: %s", reflect.TypeOf(arg1))
	}

	switch arg2.(type) {
	case Matrix, [][]float64:
		flags |= ARG2_MATRIX
	case Vector, []float64:
		flags |= ARG2_VECTOR
	case float64, float32, int:
		flags |= ARG2_SCALAR
  default:
    LogWarnf("unhandled: %s", reflect.TypeOf(arg2))
	}

	return flags
}

//
// Generic Add method
//
func Add(dataA Data, dataB Data) (sum Data) {
	flags := _argBitmask(dataA, dataB)

	switch flags {
	case ARG1_MATRIX | ARG2_MATRIX:
		LogDebug("MM")
		sum = AddMM(dataA.(Matrix), dataB.(Matrix))
	case ARG1_VECTOR | ARG2_VECTOR:
		LogDebug("VV")
		sum = AddVV(dataA.(Vector), dataB.(Vector))
	case ARG1_MATRIX | ARG2_VECTOR:
		LogDebug("MV")
		sum = AddMV(dataA.(Matrix), dataB.(Vector))
	case ARG1_VECTOR | ARG2_MATRIX:
		LogDebug("VM")
		sum = AddMV(dataB.(Matrix), dataA.(Vector))
	case ARG1_MATRIX | ARG2_SCALAR:
		LogDebug("MS")
		sum = AddMS(dataA.(Matrix), dataB.(float64))
	case ARG1_SCALAR | ARG2_MATRIX:
		LogDebug("SM")
		sum = AddMS(dataB.(Matrix), dataA.(float64))
	case ARG1_VECTOR | ARG2_SCALAR:
		LogDebug("VS")
		sum = AddVS(dataA.(Vector), dataB.(float64))
	case ARG1_SCALAR | ARG2_VECTOR:
		LogDebug("SV")
		sum = AddVS(dataB.(Vector), dataA.(float64))
	default:
		LogWarnf("unhandled flag set")
	}

	return sum
}

//
// Creates a matrix of the sums of two matricies.
//
func AddMM(matrix Matrix, matrix2 Matrix) (sum Matrix) {
	rows, cols := Size(matrix)
	rows2, cols2 := Size(matrix2)

	if rows != rows2 || cols != cols2 {
		LogErrorf("nonconformant arguments %dx%d + %dx%d\n", rows, cols, rows2, cols2)
		return sum
	}

	sum = NewMatrix(rows, cols)

	for i, row := range matrix {
		for j, val := range row {
			sum[i][j] += val + matrix2[i][j]
		}
	}

	return sum
}

//
// Creates a vector of the sums of two vectors.
//
func AddVV(vectorA Vector, vectorB Vector) (sum Vector) {
	cols := len(vectorA)
	colsB := len(vectorB)

	if cols != colsB {
		LogError("undefined")
		return sum
	}

	sum = NewVector(cols)

	for i, val := range vectorA {
		sum[i] += val + vectorB[i]
	}

	return sum
}

//
// Creates a vector of the sum of a matrix and a vector
// NOTE: arg order is irrelevant
//
func AddMV(matrix Matrix, vector Vector) (sum Matrix) {
	rows, cols := Size(matrix)
	rows2, cols2 := Size(vector)

	if cols != 1 {
		LogErrorf("nonconformant arguments %dx%d + %dx%d\n", rows, cols, rows2, cols2)
		return sum
	}

	sum = NewMatrix(rows, cols2)

	for i, row := range matrix {
		for j, val := range vector {
			sum[i][j] += row[0] + val
		}
	}

	return sum
}

//
// Creates a matrix of the sum of a matrix and a scalar
// NOTE: arg order is irrelevant
//
func AddMS(matrix Matrix, scalar float64) (sum Matrix) {
	rows, cols := Size(matrix)
	sum = NewMatrix(rows, cols)

	for i, row := range matrix {
		for j, val := range row {
			sum[i][j] += val + scalar
		}
	}

	return sum
}

//
// Creates a vector of the sum of a vector and a scalar
// NOTE: arg order is irrelevant
//
func AddVS(vector Vector, scalar float64) (sum Vector) {
	sum = NewVector(len(vector))

	for i, val := range vector {
		sum[i] = val + scalar
	}

	return sum
}

//
// Generic Sub-tract method
//
func Sub(dataA Data, dataB Data) (diff Data) {
	flags := _argBitmask(dataA, dataB)

	switch flags {
	case ARG1_MATRIX | ARG2_MATRIX:
		LogDebug("MM")
		diff = AddMM(dataA.(Matrix), dataB.(Matrix))
	case ARG1_VECTOR | ARG2_VECTOR:
		LogDebug("VV")
		diff = AddVV(dataA.(Vector), dataB.(Vector))
	case ARG1_MATRIX | ARG2_VECTOR:
		LogDebug("MV")
		diff = AddMV(dataA.(Matrix), dataB.(Vector))
	case ARG1_VECTOR | ARG2_MATRIX:
		LogDebug("VM")
		diff = AddMV(dataB.(Matrix), dataA.(Vector))
	case ARG1_MATRIX | ARG2_SCALAR:
		LogDebug("MS")
		diff = AddMS(dataA.(Matrix), dataB.(float64))
	case ARG1_SCALAR | ARG2_MATRIX:
		LogDebug("SM")
		diff = AddMS(dataB.(Matrix), dataA.(float64))
	case ARG1_VECTOR | ARG2_SCALAR:
		LogDebug("VS")
		diff = AddVS(dataA.(Vector), dataB.(float64))
	case ARG1_SCALAR | ARG2_VECTOR:
		LogDebug("SV")
		diff = AddVS(dataB.(Vector), dataA.(float64))
	default:
		LogWarnf("unhandled flag set")
	}

	return diff
}


//
// Creates a matrix of the differences of two matricies.
//
func SubMM(matrixA Matrix, matrixB Matrix) (diff Matrix) {
	rows, cols := Size(matrixA)
	diff = NewMatrix(rows, cols)

	for i, row := range matrixA {
		for j, val := range row {
			diff[i][j] += val - matrixB[i][j]
		}
	}

	return diff
}

//
// Creates a vector of the differences of two vectors.
//
func SubVV(vectorA Vector, vectorB Vector) (diff Vector) {
	_, cols := Size(vectorA)
	diff = NewVector(cols)

	for i, val := range vectorA {
		diff[i] = val - vectorB[i]
	}

	return diff
}

//
// Creates a vector of the products of a vector and a value.
//
func MulSV(factor float64, vector Vector) (prod Vector) {
	return MulVS(vector, factor)
}

//
// Creates a vector of the products of a vector and a value.
//
func MulVS(vector Vector, factor float64) (prod Vector) {
	prod = NewVector(len(vector))

	for i, val := range vector {
		prod[i] = val * factor
	}

	return prod
}

//
// Creates a matrix of the products of two matricies.
//
func MulMM(matrixA Matrix, matrixB Matrix) (prod Matrix) {
	prod = NewMatrix(len(matrixA), 1)

	// matrix * matrix multiplication => col[0] * row[0] + col[1] * row[1]
	for i, row := range matrixA {
		for j, val := range row {
			prod[i][0] += val * matrixB[j][0]
		}
	}

	return prod
}

//
// Creates a matrix of the products of a value and matrix.
//
func MulVM(vector Vector, matrix Matrix) (prod Vector) {
	_, cols := Size(matrix)
	prod = NewVector(cols)

	// vetor * matrix multiplication => v[0] * m[0][0] + v[0] * m[1][0] + ...
	for i, row := range matrix {
		for j, val := range row {
			prod[j] += vector[i] * val
		}
	}

	return prod
}

//
// Creates a matrix of the products of a value and matrix.
//
func MulSM(factor float64, matrix Matrix) (prod Matrix) {
	rows, cols := Size(matrix)
	prod = NewMatrix(rows, cols)

	for i, row := range matrix {
		for j, val := range row {
			prod[i][j] = factor * val
		}
	}

	return prod
}

//
// Creates a matrix of the quotients of a matrix and a value.
//
func DivMS(matrix Matrix, divisor float64) (quot Matrix) {
	rows, cols := Size(matrix)
	quot = NewMatrix(rows, cols)

	for i, row := range matrix {
		for j, val := range row {
			quot[i][j] += val / divisor
		}
	}

	return quot
}

//
// Creates a vector of the quotients of a vector and a value.
//
func DivVS(vector Vector, divisor float64) (quot Vector) {
	quot = NewVector(len(vector))

	for i, val := range vector {
		quot[i] = val / divisor
	}

	return quot
}
