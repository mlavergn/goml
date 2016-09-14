// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	. "golog"
	"math"
)

//
// Generic element-wise addition method
//
func DotAdd(arg1 Data, arg2 Data) (sum Data) {
	sum = Add(arg1, arg2)
	return sum
}

//
// Generic addition method
//
func Add(arg1 Data, arg2 Data) (sum Data) {
	flags := _argBitmask(arg1, arg2)

	switch flags {
	case ARG1_MATRIX | ARG2_MATRIX:
		LogDebug("AddMM")
		sum = _addMM(arg1.(Matrix), arg2.(Matrix))
	case ARG1_VECTOR | ARG2_VECTOR:
		LogDebug("AddVV")
		sum = _addVV(arg1.(Vector), arg2.(Vector))
	case ARG1_MATRIX | ARG2_VECTOR:
		LogDebug("AddMV")
		sum = _addMV(arg1.(Matrix), arg2.(Vector))
	case ARG1_VECTOR | ARG2_MATRIX:
		LogDebug("AddVM")
		sum = _addMV(arg2.(Matrix), arg1.(Vector))
	case ARG1_MATRIX | ARG2_SCALAR:
		LogDebug("AddMS")
		sum = _addMS(arg1.(Matrix), arg2.(float64))
	case ARG1_SCALAR | ARG2_MATRIX:
		LogDebug("AddSM")
		sum = _addMS(arg2.(Matrix), arg1.(float64))
	case ARG1_VECTOR | ARG2_SCALAR:
		LogDebug("AddVS")
		sum = _addVS(arg1.(Vector), arg2.(float64))
	case ARG1_SCALAR | ARG2_VECTOR:
		LogDebug("AddSV")
		sum = _addVS(arg2.(Vector), arg1.(float64))
	default:
		LogError("Unhandled argument type / combination")
	}

	return sum
}

//
// Creates a matrix of the sums of two matricies.
//
func _addMM(matrix Matrix, matrix2 Matrix) (sum Matrix) {
	rows, cols := Size(matrix)
	rows2, cols2 := Size(matrix2)

	if rows != rows2 || cols != cols2 {
		LogErrorf("error: operator +: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
		return sum
	}

	sum = NewMatrix(rows, cols)

	for i, row := range matrix {
		for j, val := range row {
			sum[i][j] = val + matrix2[i][j]
		}
	}

	return sum
}

//
// Creates a vector of the sums of two vectors.
//
func _addVV(vector Vector, vector2 Vector) (sum Vector) {
	rows, cols := Size(vector)
	rows2, cols2 := Size(vector2)

	if cols != cols2 {
		LogErrorf("error: operator +: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
		return sum
	}

	sum = NewVector(cols)

	for i, val := range vector {
		sum[i] = val + vector2[i]
	}

	return sum
}

//
// Creates a vector of the sum of a matrix and a vector
// NOTE: arg order is irrelevant
//
func _addMV(matrix Matrix, vector Vector) (sum Matrix) {
	rows, cols := Size(matrix)
	rows2, cols2 := Size(vector)

	if cols != cols2 {
		LogErrorf("error: operator +: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
		return sum
	}

	sum = NewMatrix(rows, cols2)

	for i, row := range matrix {
		for j, val := range row {
			sum[i][j] = val + vector[j]
		}
	}

	return sum
}

//
// Creates a matrix of the sum of a matrix and a scalar
// NOTE: arg order is irrelevant
//
func _addMS(matrix Matrix, scalar float64) (sum Matrix) {
	rows, cols := Size(matrix)
	sum = NewMatrix(rows, cols)

	for i, row := range matrix {
		for j, val := range row {
			sum[i][j] = val + scalar
		}
	}

	return sum
}

//
// Creates a vector of the sum of a vector and a scalar
// NOTE: arg order is irrelevant
//
func _addVS(vector Vector, scalar float64) (sum Vector) {
	sum = NewVector(len(vector))

	for i, val := range vector {
		sum[i] = val + scalar
	}

	return sum
}

//
// Generic element-wise subtraction method
//
func DotSub(arg1 Data, arg2 Data) (diff Data) {
	diff = Sub(arg1, arg2)
	return diff
}

//
// Generic subtraction method
//
func Sub(arg1 Data, arg2 Data) (diff Data) {
	flags := _argBitmask(arg1, arg2)

	switch flags {
	case ARG1_MATRIX | ARG2_MATRIX:
		LogDebug("SubMM")
		diff = _subMM(arg1.(Matrix), arg2.(Matrix))
	case ARG1_VECTOR | ARG2_VECTOR:
		LogDebug("SubVV")
		diff = _subVV(arg1.(Vector), arg2.(Vector))
	case ARG1_MATRIX | ARG2_VECTOR:
		LogDebug("SubMV")
		diff = _subMV(arg1.(Matrix), arg2.(Vector))
	case ARG1_VECTOR | ARG2_MATRIX:
		LogDebug("SubVM")
		diff = _subVM(arg1.(Vector), arg2.(Matrix))
	case ARG1_MATRIX | ARG2_SCALAR:
		LogDebug("SubMS")
		diff = _subMS(arg1.(Matrix), arg2.(float64))
	case ARG1_SCALAR | ARG2_MATRIX:
		LogDebug("SubSM")
		diff = _subSM(arg1.(float64), arg2.(Matrix))
	case ARG1_VECTOR | ARG2_SCALAR:
		LogDebug("SubVS")
		diff = _subVS(arg1.(Vector), arg2.(float64))
	case ARG1_SCALAR | ARG2_VECTOR:
		LogDebug("SubSV")
		diff = _subSV(arg1.(float64), arg2.(Vector))
	default:
		LogError("Unhandled argument type / combination")
	}

	return diff
}

//
// Creates a matrix of the differences of two matricies.
//
func _subMM(matrix Matrix, matrix2 Matrix) (diff Matrix) {
	rows, cols := Size(matrix)
	diff = NewMatrix(rows, cols)

	for i, row := range matrix {
		for j, val := range row {
			diff[i][j] = val - matrix2[i][j]
		}
	}

	return diff
}

//
// Creates a vector of the differences of two vectors.
//
func _subVV(vector Vector, vector2 Vector) (diff Vector) {
	_, cols := Size(vector)
	diff = NewVector(cols)

	for i, val := range vector {
		diff[i] = val - vector2[i]
	}

	return diff
}

//
// Creates a vector of the differences of two vectors.
//
func _subMV(matrix Matrix, vector Vector) (diff Matrix) {
	rows, cols := Size(matrix)
	rows2, cols2 := Size(vector)

	if cols != cols2 {
		LogErrorf("error: operator -: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
		return diff
	}

	diff = NewMatrix(rows, cols)

	for i, row := range matrix {
		for j, val := range row {
			diff[i][j] = val - vector[j]
		}
	}

	return diff
}

//
// Creates a vector of the differences of two vectors.
//
func _subVM(vector Vector, matrix Matrix) (diff Matrix) {
	rows, cols := Size(vector)
	rows2, cols2 := Size(matrix)

	if cols != cols2 {
		LogErrorf("error: operator -: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
		return diff
	}

	diff = NewMatrix(rows2, cols2)

	for i, row := range matrix {
		for j, val := range row {
			diff[i][j] = vector[j] - val
		}
	}

	return diff
}

//
// Creates a matrix of the difference of a matrix and a scalar.
//
func _subMS(matrix Matrix, scalar float64) (diff Matrix) {
	rows, cols := Size(matrix)

	diff = NewMatrix(rows, cols)

	for i, row := range matrix {
		for j, val := range row {
			diff[i][j] = val - scalar
		}
	}

	return diff
}

//
// Creates a matrix of the difference of a scalar and a matrix.
//
func _subSM(scalar float64, matrix Matrix) (diff Matrix) {
	rows, cols := Size(matrix)

	diff = NewMatrix(rows, cols)

	for i, row := range matrix {
		for j, val := range row {
			diff[i][j] = scalar - val
		}
	}

	return diff
}

//
// Creates a matrix of the difference of a matrix and a scalar.
//
func _subVS(vector Vector, scalar float64) (diff Vector) {
	_, cols := Size(vector)

	diff = NewVector(cols)

	for i, val := range vector {
		diff[i] = val - scalar
	}

	return diff
}

//
// Creates a matrix of the difference of a scalar and a matrix.
//
func _subSV(scalar float64, vector Vector) (diff Vector) {
	_, cols := Size(vector)

	diff = NewVector(cols)

	for i, val := range vector {
		diff[i] = val - scalar
	}

	return diff
}

//
// Generic element-wise multiplication method
//
func DotMul(arg1 Data, arg2 Data) (prod Data) {
	flags := _argBitmask(arg1, arg2)

	switch flags {
	case ARG1_MATRIX | ARG2_MATRIX:
		LogDebug("MulMM")
		prod = _dotMulMM(arg1.(Matrix), arg2.(Matrix))
	case ARG1_VECTOR | ARG2_VECTOR:
		LogDebug("MulVV")
		prod = _mulVV(arg1.(Vector), arg2.(Vector))
	case ARG1_MATRIX | ARG2_VECTOR:
		LogDebug("MulMV")
		rows, cols := Size(arg1)
		rows2, cols2 := Size(arg2)
		LogErrorf("error: operator *: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
	case ARG1_VECTOR | ARG2_MATRIX:
		LogDebug("MulVM")
		prod = _mulVM(arg1.(Vector), arg2.(Matrix))
	case ARG1_MATRIX | ARG2_SCALAR:
		LogDebug("MulMS")
		prod = _mulSM(arg2.(float64), arg1.(Matrix))
	case ARG1_SCALAR | ARG2_MATRIX:
		LogDebug("MulSM")
		prod = _mulSM(arg1.(float64), arg2.(Matrix))
	case ARG1_VECTOR | ARG2_SCALAR:
		LogDebug("MulVS")
		prod = _mulSV(arg2.(float64), arg1.(Vector))
	case ARG1_SCALAR | ARG2_VECTOR:
		LogDebug("SV")
		prod = _mulSV(arg1.(float64), arg2.(Vector))
	default:
		LogError("Unhandled argument type / combination")
	}

	return prod
}

//
// Creates a matrix of the products of two matricies.
//
func _dotMulMM(matrix Matrix, matrix2 Matrix) (prod Matrix) {
	rows, cols := Size(matrix)
	rows2, cols2 := Size(matrix2)

	if rows != rows || cols != cols2 {
		LogErrorf("error: operator *: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
		return prod
	}

	prod = NewMatrix(rows, cols)

	// matrix * matrix multiplication => col[0] * row[0] + col[1] * row[1]
	for i, row := range matrix {
		for j, val := range row {
			prod[i][j] = val * matrix2[i][j]
		}
	}

	return prod
}

//
// Generic multiplication method
//
func Mul(arg1 Data, arg2 Data) (prod Data) {
	flags := _argBitmask(arg1, arg2)

	switch flags {
	case ARG1_MATRIX | ARG2_MATRIX:
		LogDebug("MulMM")
		prod = _mulMM(arg1.(Matrix), arg2.(Matrix))
	case ARG1_VECTOR | ARG2_VECTOR:
		LogDebug("MulVV")
		prod = _mulVV(arg1.(Vector), arg2.(Vector))
	case ARG1_MATRIX | ARG2_VECTOR:
		LogDebug("MulMV")
		rows, cols := Size(arg1)
		rows2, cols2 := Size(arg2)
		LogErrorf("error: operator *: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
	case ARG1_VECTOR | ARG2_MATRIX:
		LogDebug("MulVM")
		prod = _mulVM(arg1.(Vector), arg2.(Matrix))
	case ARG1_MATRIX | ARG2_SCALAR:
		LogDebug("MulMS")
		prod = _mulSM(arg2.(float64), arg1.(Matrix))
	case ARG1_SCALAR | ARG2_MATRIX:
		LogDebug("MulSM")
		prod = _mulSM(arg1.(float64), arg2.(Matrix))
	case ARG1_VECTOR | ARG2_SCALAR:
		LogDebug("MulVS")
		prod = _mulSV(arg2.(float64), arg1.(Vector))
	case ARG1_SCALAR | ARG2_VECTOR:
		LogDebug("SV")
		prod = _mulSV(arg1.(float64), arg2.(Vector))
	default:
		LogError("Unhandled argument type / combination")
	}

	return prod
}

//
// Creates a matrix of the products of two matricies.
//
func _mulMM(matrix Matrix, matrix2 Matrix) (prod Matrix) {
	rows, cols := Size(matrix)
	rows2, cols2 := Size(matrix2)

	if cols != rows2 {
		LogErrorf("error: operator *: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
		return prod
	}

	prod = NewMatrix(len(matrix), 1)

	// matrix * matrix multiplication => col[0] * row[0] + col[1] * row[1]
	for i, row := range matrix {
		for j, val := range row {
			prod[i][0] += val * matrix2[j][0]
		}
	}

	return prod
}

//
// Creates a vector of the product of two vectors.
//
func _mulVV(vector Vector, vector2 Vector) (prod Vector) {
	rows, cols := Size(vector)
	rows2, cols2 := Size(vector2)

	if cols != cols2 {
		LogErrorf("error: operator *: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
		return prod
	}

	prod = NewVector(cols)

	for i, val := range vector {
		prod[i] = val * vector2[i]
	}

	return prod
}

//
// Creates a matrix of the products of a value and matrix.
//
func _mulVM(vector Vector, matrix Matrix) (prod Vector) {
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
func _mulSM(factor float64, matrix Matrix) (prod Matrix) {
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
// Creates a vector of the products of a vector and a value.
//
func _mulSV(factor float64, vector Vector) (prod Vector) {
	prod = NewVector(len(vector))

	for i, val := range vector {
		prod[i] = val * factor
	}

	return prod
}

//
// Generic element-wise division method
//
func DotDiv(arg1 Data, arg2 Data) (quot Data) {
	quot = Div(arg1, arg2)
	return quot
}

//
// Generic division method
//
func Div(arg1 Data, arg2 Data) (quot Data) {
	flags := _argBitmask(arg1, arg2)

	switch flags {
	case ARG1_MATRIX | ARG2_MATRIX:
		LogDebug("DivMM")
		quot = _divMM(arg1.(Matrix), arg2.(Matrix))
	case ARG1_VECTOR | ARG2_VECTOR:
		LogDebug("DivVV")
		rows, cols := Size(arg1)
		rows2, cols2 := Size(arg2)
		LogErrorf("error: operator /: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
	case ARG1_MATRIX | ARG2_VECTOR:
		LogDebug("DivMV")
		rows, cols := Size(arg1)
		rows2, cols2 := Size(arg2)
		LogErrorf("error: operator /: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
	case ARG1_VECTOR | ARG2_MATRIX:
		LogDebug("DivVM")
		quot = _divVM(arg1.(Vector), arg2.(Matrix))
	case ARG1_MATRIX | ARG2_SCALAR:
		LogDebug("DivMS")
		quot = _divMS(arg1.(Matrix), arg2.(float64))
	case ARG1_SCALAR | ARG2_MATRIX:
		LogDebug("DivSM")
		rows, cols := Size(arg1)
		rows2, cols2 := Size(arg2)
		LogErrorf("error: operator /: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
	case ARG1_VECTOR | ARG2_SCALAR:
		LogDebug("DivVS")
		quot = _divVS(arg1.(Vector), arg2.(float64))
	case ARG1_SCALAR | ARG2_VECTOR:
		LogDebug("DivSV")
		rows, cols := Size(arg1)
		rows2, cols2 := Size(arg2)
		LogErrorf("error: operator /: nonconformant arguments (op1 is %dx%d, op2 is %dx%d)\n", rows, cols, rows2, cols2)
	default:
		LogError("Unhandled argument type / combination")
	}

	return quot
}

//
// Creates a matrix of the quotients of a matrix and a value.
//
func _divMM(matrix Matrix, matrix2 Matrix) (quot Matrix) {
	rows, cols := Size(matrix)
	quot = NewMatrix(rows, cols)

	for i, row := range matrix {
		for j, val := range row {
			quot[i][j] = val * math.Pow(matrix2[i][j], -1)
		}
	}

	return quot
}

//
// Creates a matrix of the quotients of a matrix and a value.
//
func _divMV(matrix Matrix, vector Vector) (quot Matrix) {
	rows, cols := Size(matrix)
	quot = NewMatrix(rows, cols)

	// TBD

	return quot
}

//
// Creates a matrix of the quotients of a matrix and a value.
//
func _divVM(vector Vector, matrix Matrix) (quot Matrix) {
	rows, cols := Size(matrix)
	quot = NewMatrix(rows, cols)

	// This is very complicated

	return quot
}

//
// Creates a matrix of the quotients of a matrix and a value.
//
func _divMS(matrix Matrix, divisor float64) (quot Matrix) {
	rows, cols := Size(matrix)
	quot = NewMatrix(rows, cols)

	for i, row := range matrix {
		for j, val := range row {
			quot[i][j] = val / divisor
		}
	}

	return quot
}

//
// Creates a vector of the quotients of a vector and a value.
//
func _divVS(vector Vector, divisor float64) (quot Vector) {
	_, cols := Size(vector)
	quot = NewVector(cols)

	for i, val := range vector {
		quot[i] = val / divisor
	}

	return quot
}
