// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	"math"
)

//
// This set of statistics functions was created for use in a
// larger project. They are primarily focused on simplicity
// and performance at the cost of safety. There are no plans
// to add safety checks at this time, it's left to the caller
// to pass in sane values ... or else!
//

//
// Round float to precision decimals
//
func Round(v float64, p int) (result float64) {
	shift := math.Pow(10, float64(p))

	if v < 0 {
		result = math.Ceil(v*shift-0.5) / shift
	} else {
		result = math.Floor(v*shift+0.5) / shift
	}

	return result
}

//
// Median of a vector
//
func Median(v *[]float64) (result float64) {
	n := len(*v) / 2
	result = ((*v)[n] + (*v)[n-1]) / 2

	return result
}

//
// Sum of a vector
//
func Sum(v *[]float64) (result float64) {
	for _, v := range *v {
		result += v
	}

	return result
}

//
// Mean of a vector
//
func Mean(v *[]float64) (result float64) {
	result = Sum(v) / float64(len(*v))

	return result
}

//
// Mean excluding the top and bottom X% of outliers
//
func SampleMean(v *[]float64, trim float64) (r float64) {
	count := len(*v)
	low := int(float64(count) * trim)
	high := int(float64(count) * (1.0 - trim))
	for _, v := range (*v)[low : high-1] {
		r += v
	}
	r /= float64(high - low)

	return r
}

//
// Mode of a vector
// NOTE: This was the most performant
//   implmentation after many tests
//   the downside is the space
//   complexity is max O(3n)
//
func Mode(v *[]float64) (result float64) {
	n := len(*v)
	val := make([]int, n)
	valtop := -1

	ht := make(map[float64]*int)

	maxoccurs := 0

	var nval *int

	for _, key := range *v {
		nval = ht[key]
		if nval == nil {
			valtop += 1
			nval = &val[valtop]
			ht[key] = nval
		}
		*nval += 1
		if *nval >= maxoccurs {
			if *nval == maxoccurs && key > result {
				continue
			}
			result = key
			maxoccurs = *nval
		}
	}

	return result
}

//
// Standard deviation of a vector
//
func StandardDeviation(v *[]float64) (result float64) {
	n := float64(len(*v))
	m := Mean(v)

	for _, v := range *v {
		result += math.Pow(v-m, 2) / n
	}

	result = math.Sqrt(result)

	return result
}

//
// Confidence interval of vector
//
func ConfidenceInterval(v *[]float64, cv float64) (lower float64, upper float64) {
	meanV := Mean(v)
	bounds := (cv * StandardDeviation(v) / math.Sqrt(float64(len(*v))))

	lower = meanV - bounds
	upper = meanV + bounds

	return lower, upper
}

//
// Coefficient of correlation
//
func CoefficientOfCorrelation(x *[]float64, y *[]float64) float64 {
	_, _, r := _coefficientOfCorrelation(x, y)
	return r
}

//
// Internal: We wrap this so we can reuse the std dev and ci in Slope
//
func _coefficientOfCorrelation(x *[]float64, y *[]float64) (stdDevX float64, stdDevY float64, r float64) {
	n := len(*x)

	meanX := Mean(x)
	meanY := Mean(y)

	stdDevX = StandardDeviation(x)
	stdDevY = StandardDeviation(y)

	XxY := 0.0
	for i := 0; i < n; i++ {
		XxY += ((*x)[i] - meanX) * ((*y)[i] - meanY)
	}

	r = (XxY / (stdDevX * stdDevY)) / float64(n-1)

	return stdDevX, stdDevY, r
}

func Slope(x *[]float64, y *[]float64) (r float64) {
	stdDevX, stdDevY, c := _coefficientOfCorrelation(x, y)

	r = c * (stdDevY / stdDevX)

	return r
}
