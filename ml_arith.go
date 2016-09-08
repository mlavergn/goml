// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	"log"
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
	default:
		flags |= ARG1_SCALAR
	}

	switch arg2.(type) {
	case Matrix, [][]float64:
		flags |= ARG2_MATRIX
	case Vector, []float64:
		flags |= ARG2_VECTOR
	default:
		flags |= ARG2_SCALAR
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
		log.Println("MM")
	case ARG1_VECTOR | ARG2_VECTOR:
		log.Println("VV")
		sum = AddVV(dataA.(Vector), dataB.([]float64))
	case ARG1_VECTOR | ARG2_SCALAR:
		log.Println("VX")
	default:
		log.Println("other")
	}

	return sum
}

//
// Creates a vector of the sums of two vectors.
//
func AddVV(vectorA Vector, vectorB Vector) (sum Vector) {
	sum = NewVector(len(vectorA))

	for i, val := range vectorA {
		sum[i] = val + vectorB[i]
	}

	return sum
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
func MulXV(factor float64, vectorA Vector) (prod Vector) {
	return MulVX(vectorA, factor)
}

//
// Creates a vector of the products of a vector and a value.
//
func MulVX(vectorA Vector, factor float64) (prod Vector) {
	prod = NewVector(len(vectorA))

	for i, val := range vectorA {
		prod[i] = val * factor
	}

	return prod
}

//
// Creates a vector of the quotients of a vector and a value.
//
func DivVX(vectorA Vector, divisor float64) (quot Vector) {
	quot = NewVector(len(vectorA))

	for i, val := range vectorA {
		quot[i] = val / divisor
	}

	return quot
}

//
// Creates a matrix of the sums of two matricies.
//
func AddMM(matrixA Matrix, matrixB Matrix) (sum Matrix) {
	rows, cols := Size(matrixA)
	sum = NewMatrix(rows, cols)

	for i, row := range matrixA {
		for j, val := range row {
			sum[i][j] += val + matrixB[i][j]
		}
	}

	return sum
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
func MulXM(factor float64, matrix Matrix) (prod Matrix) {
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
func DivMX(matrixA Matrix, divisor float64) (quot Matrix) {
	rows, cols := Size(matrixA)
	quot = NewMatrix(rows, cols)

	for i, row := range matrixA {
		for j, val := range row {
			quot[i][j] += val / divisor
		}
	}

	return quot
}
