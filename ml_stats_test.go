// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	"testing"
)

var testA = []float64{15, 12, 8, 8, 7, 7, 7, 6, 5, 3}
var testB = []float64{10, 25, 17, 11, 13, 17, 20, 13, 9, 15}

func _testAB(t *testing.T, aIn float64, aExp float64, bIn float64, bExp float64) {
	if Round(aIn, 6) != aExp || Round(bIn, 6) != bExp {
		t.Errorf("%f / %f vs expected %f / %f", aIn, bIn, aExp, bExp)
	}
}

func _testA(t *testing.T, aIn float64, aExp float64) {
	if Round(aIn, 6) != aExp {
		t.Errorf("%f vs expected %f", aIn, aExp)
	}
}

func TestMedian(t *testing.T) {
	_testAB(t, Median(&testA), 7.0, Median(&testB), 15.0)
}

func TestSum(t *testing.T) {
	_testAB(t, Sum(&testA), 78.0, Sum(&testB), 150.0)
}

func TestMean(t *testing.T) {
	_testAB(t, Mean(&testA), 7.8, Mean(&testB), 15.0)
}

func TestSampleMean(t *testing.T) {
	_testAB(t, SampleMean(&testA, 0.1), 6.875, SampleMean(&testB, 0.1), 14.5)
}

func TestMode(t *testing.T) {
	_testAB(t, Mode(&testA), 7.0, Mode(&testB), 13.0)
}

func TestStandardDeviation(t *testing.T) {
	_testAB(t, StandardDeviation(&testA), 3.249615, StandardDeviation(&testB), 4.669047)
}

func TestConfidenceInterval(t *testing.T) {
	lowA, upA := ConfidenceInterval(&testA, .95)
	lowB, upB := ConfidenceInterval(&testB, .95)

	_testAB(t, lowA, 6.823762, upA, 8.776238)
	_testAB(t, lowB, 13.597342, upB, 16.402658)
}

func TestCoefficientOfCorrelation(t *testing.T) {
	_testA(t, CoefficientOfCorrelation(&testA, &testB), 0.161109)
}

func TestSlope(t *testing.T) {
	_testA(t, Slope(&testA, &testB), 0.231481)
}
