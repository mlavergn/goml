// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	"math"
)

// Eulers number (1 + 1/1 + 1/(1*2) + 1/(1*2*3) + ...)
const e float64 = 2.71828182845904523536028747135266249775724709369995

//
// LR Cost Function
// J = (1 / (2 * m)) * sum((X * theta - y) .^ 2)
//
func ComputeLRCost(X Matrix, y Matrix, theta Matrix) (J float64) {
	m := len(y)

	h := 0.0
	for row := 0; row < m; row++ {
		// matrix multiplication => col[0] * row[0] + col[1] * row[1]
		hX := 0.0
		for col := 0; col < int(len(X[0])); col++ {
			hX += X[row][col] * theta[col][0]
		}
		h += math.Pow((hX - y[row][0]), 2)
	}

	J = (1.0 / (2.0 * float64(m))) * h

	return J
}

//
// theta = theta - ((alpha * ((theta' * X') - y')) * X / m)'
//
func GradientDescent(X Matrix, y, intheta Matrix, alpha float64, num_iters int) (theta Matrix, J_history Vector) {
	J_history = NewEmptyVector()

	theta = intheta

	m := len(y)
	for i := 0; i < num_iters; i++ {
		Jc := ComputeLRCost(X, y, theta)

		tX := Mul(T(theta), T(X))
		tXy := Sub(tX, T(y))
		atXy := Mul(alpha, tXy)
		atXyX := Mul(atXy, X)
		atXyXm := Div(atXyX, float64(m))
		theta = Sub(theta, T(atXyXm)).(Matrix)

		Jn := ComputeLRCost(X, y, theta)

		if Jn == Jc || Jn > Jc {
			break
		}

		J_history = append(J_history, Jn)
	}

	return theta, J_history
}

//
// result keeps the same type as z
// g = 1 ./ (1 + e.^-z)
//
func Sigmoid(z Data) (g Data) {
	switch z.(type) {
	case Matrix, [][]float64:
		row, col := Size(z)
		g = NewMatrix(row, col)
		for i, row := range z.(Matrix) {
			for j, val := range row {
				g.([][]float64)[i][j] = 1 / (1 - math.Pow(e, -val))
			}
		}
	case Vector, []float64:
		_, col := Size(z)
		g = NewVector(col)
		for i, val := range z.(Vector) {
			g.([]float64)[i] = 1 / (1 - math.Pow(e, -val))
		}
	case float64:
		// . for each
		g = 1 / (1 - math.Pow(e, -z.(float64)))
	default:
		// unhandled
	}

	return g
}

//
// Calculate the accuracy evaluation metric
// accuracy = P1A1 + P0A0 / (P1A1 + P1A0 + P0A1 + P0A0)
//
func Accuracy(ap [][]int) (accuracy float64) {
	accuracy = float64(ap[0][0]+ap[1][1]) / float64(ap[0][0]+ap[0][1]+ap[1][0]+ap[1][1])

	return accuracy
}

//
// Calculate the precision evaluation metric
// precision = P1A1 / (P1A1 + P1A0)
//
func Precision(ap [][]int) (precision float64) {
	precision = float64(ap[0][0]) / float64(ap[0][0]+ap[0][1])

	return precision
}

//
// Calculate the recall evaluation metric
// recall = P1A1 / (P1A1 + P1A1)
//
func Recall(ap [][]int) (recall float64) {
	recall = float64(ap[0][0]) / float64(ap[0][0]+ap[1][0])

	return recall
}

//
// Calculate the F1 score evaluation metric
// Determines the appropriate threshold to use for precision / recall
// in an attempt to maximize the F1 score
//
func F1Score(ap [][]int) (f1 float64) {
	precision := Precision(ap)
	recall := Recall(ap)

	f1 = (2.0 * precision * recall) / (precision + recall)

	return f1
}
