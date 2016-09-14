// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	. "golog"
	"math/rand"
	"reflect"
	"time"
)

//
// Creates an empty vector.
//
func NewEmptyVector() (vector Vector) {
	vector = Vector{}

	return vector
}

//
// Creates a vector.
//
func NewVector(cols int) (vector Vector) {
	vector = make(Vector, cols)

	return vector
}

//
// Creates an empty matrix with no rows initialized.
//
func NewEmptyMatrix(rows int) (matrix Matrix) {
	matrix = make(Matrix, rows)

	return matrix
}

//
// Creates a matrix.
//
func NewMatrix(rows int, cols int) (matrix Matrix) {
	matrix = make(Matrix, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = NewVector(cols)
	}

	return matrix
}

//
// Size of the data structure as rows, cols.
//
func Size(data Data) (rows int, cols int) {
	switch data.(type) {
	case Matrix, [][]float64:
		rows = len(data.(Matrix))
		cols = len(data.(Matrix)[0])
	case Vector, []float64:
		rows = 1
		cols = len(data.(Vector))
	case float64:
		rows = 1
		cols = 1
	default:
		LogWarnf("unhandled: %s", reflect.TypeOf(data))
	}

	return rows, cols
}

func Equal(x Data, y Data) (result bool) {
	bitmask := _argBitmask(x, y)
	switch bitmask {
	case ARG1_MATRIX | ARG2_MATRIX:
		result = _equalMM(x.(Matrix), y.(Matrix))
	case ARG1_VECTOR | ARG2_VECTOR:
		result = _equalVV(x.(Vector), y.(Vector))
	}

	return result
}

func _equalVV(x Vector, y Vector) (result bool) {
	result = true

	if len(x) != len(y) {
		result = false
	} else {
		for i, val := range x {
			if val != y[i] {
				result = false
			}
		}
	}

	return result
}

func _equalMM(x Matrix, y Matrix) (result bool) {
	result = true

	if len(x) != len(y) {
		result = false
	} else if len(x[0]) != len(y[0]) {
		result = false
	} else {
		for rowi, row := range x {
			for coli, v := range row {
				if v != y[rowi][coli] {
					result = false
				}
			}
		}
	}

	return result
}

//
// Transpose a vector or matrix.
//
func Transpose(inputData Data) (data Data) {
	// invert the col and row vals
	inputVector := false
	outputVector := false
	switch inputData.(type) {
	case Matrix, [][]float64:
		// we invert the returns here
		cols, rows := Size(inputData)
		if rows > 1 {
			data = NewMatrix(rows, cols)
		} else {
			// only 1 row, so vectorize the return
			data = NewVector(cols)
			outputVector = true
		}
	case Vector, []float64:
		inputVector = true
		// we invert the returns here
		cols, rows := Size(inputData)
		data = NewMatrix(rows, cols)
	default:
		LogWarnf("unhandled: %s", reflect.TypeOf(inputData))
	}

	if !inputVector {
		for j, row := range inputData.(Matrix) {
			if !outputVector {
				for i, colval := range row {
					data.(Matrix)[i][j] = colval
				}
			} else {
				data.(Vector)[j] = row[0]
			}
		}
	} else {
		for i, val := range inputData.(Vector) {
			data.(Matrix)[i][0] = val
		}
	}

	return data
}

//
// Shorthand for Transpose
//
type _transpose func(Data) Data

var T _transpose = Transpose

//
// Creates a size x size identity matrix.
//
func Eye(size int) (matrix Matrix) {
	matrix = NewMatrix(size, size)

	for i := 0; i < size; i++ {
		matrix[i][i] = 1
	}

	return matrix
}

//
// Creates a data structure filled with 0's.
//
func Zeros(rows int, cols int) (data Data) {
	if rows > 1 {
		data = NewMatrix(rows, cols)
	} else {
		data = NewVector(cols)
	}

	return data
}

//
// Creates a data structure filled with 1's.
//
func Ones(rows int, cols int) (data Data) {
	if rows > 1 {
		data = NewMatrix(rows, cols)
		for _, row := range data.(Matrix) {
			for i := range row {
				row[i] = 1
			}
		}
	} else {
		data = NewVector(cols)
		for i := range data.(Vector) {
			data.(Vector)[i] = 1
		}
	}

	return data
}

//
// Creates a data structure filled with a sequence.
//
func Seq(rows int, cols int) (matrix Matrix) {
	matrix = NewMatrix(rows, cols)

	// looks awkward, but it's the most performant way
	seq := 1.0
	for _, row := range matrix {
		for i := range row {
			row[i] = seq
			seq += 1.0
		}
	}

	return matrix
}

//
// Creates a vector from a matrix column. The col parameter is 1-based.
//
func Cols(inmatrix Matrix, colFrom int, colTo int) (matrix Matrix) {
	matrix = NewEmptyMatrix(len(inmatrix))

	cols := colTo - colFrom + 1
	for i, row := range inmatrix {
		newRow := NewVector(cols)
		k := 0
		for j := colFrom - 1; j < colTo; j++ {
			newRow[k] = row[j]
			k += 1
		}
		matrix[i] = newRow
	}

	return matrix
}

//
// Creates a matrix filled with 1's.
//
func Rand(rows int, cols int) (matrix Matrix) {
	matrix = NewMatrix(rows, cols)

	rand.Seed(time.Now().UTC().UnixNano())

	// looks awkward, but it's the most performant way
	for _, row := range matrix {
		for i := range row {
			row[i] = rand.Float64()
		}
	}

	return matrix
}

//
// Unrolls a matrix into a vector.
//
func Unroll(matrix Matrix) (vector Vector) {
	rows, cols := Size(matrix)
	vector = NewVector(rows * cols)

	for i, row := range matrix {
		for j, val := range row {
			vector[(i*cols)+j] = val
		}
	}

	return vector
}

//
// Reshape a vector into a matrix.
//
func Reshape(vector Vector, rows int, cols int) (matrix Matrix) {
	matrix = NewMatrix(rows, cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = vector[(i*cols)+j]
		}
	}

	return matrix
}

func Join(datas ...interface{}) (data Data) {
	switch datas[0].(type) {
	case Matrix, [][]float64:
		matricies := make([]Matrix, len(datas))
		for i, data := range datas {
			matricies[i] = data.(Matrix)
		}
		return _joinM(matricies...)
	case Vector, []float64:
		vectors := make([]Vector, len(datas))
		for i, data := range datas {
			vectors[i] = data.(Vector)
		}
		return _joinV(vectors...)
	default:
		LogWarnf("unhandled: %s", reflect.TypeOf(datas))
	}

	return data
}

//
// Join multiple arrays.
//
func _joinV(vectors ...Vector) (vector Vector) {
	vector = Vector{}

	// keep it simple
	for _, arg := range vectors {
		vector = append(vector, arg...)
	}

	return vector
}

//
// Join multiple arrays.
//
func _joinM(matrices ...Matrix) (matrix Matrix) {
	matrix = nil
	rows := 0
	cols := 0

	rows, _ = Size(matrices[0])
	for _, arg := range matrices {
		cols += len(arg[0])
	}

	matrix = NewMatrix(rows, cols)

	offset := 0
	for _, arg := range matrices {
		for i, row := range arg {
			for j, val := range row {
				matrix[i][j+offset] = val
			}
		}
		// calculate the new column offset for the next matrix
		offset += len(arg[0])
	}

	return matrix
}
