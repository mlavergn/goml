// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

//
// Pretty prints a float value by removing trailing
// zeros and presenting as an integer if whole
//
func _formatFloat(v float64) string {
	r := fmt.Sprintf("%.3f", v)
	r = strings.Trim(r, "0.")

	// don't send back an empty string if we end up
	// with a zero through rounding and/or input
	if len(r) == 0 {
		r = "0"
	}

	return r
}

//
// Pretty prints the duration as the number of ns/us/ms
//
func _formatDuration(elapsed float64) string {
	var r string

	ns := elapsed
	if ns < 1000.0 {
		r = fmt.Sprintf("%sns", _formatFloat(ns))
	} else {
		us := ns / 1000.0
		if us < 1000.0 {
			r = fmt.Sprintf("%sus", _formatFloat(us))
		} else {
			ms := us / 1000.0
			r = fmt.Sprintf("%sms", _formatFloat(ms))
		}
	}

	return r
}

//
// Wrapper for benchmarking methods
//
type Benchmark struct {
	readno int
	reads  []float64
	timer  time.Time
}

//
// Constructor
//
func NewBenchmark(iters int) *Benchmark {
	r := &Benchmark{reads: make([]float64, iters), timer: time.Now()}

	return r
}

//
// Marks a measurement
//
func (self *Benchmark) Measure() {
	e := time.Since(self.timer)
	self.reads[self.readno] = float64(e)
	self.readno += 1
	self.timer = time.Now()
}

//
// Resets for a new measurement
//
func (self *Benchmark) Reset() {
	self.timer = time.Now()
	self.readno = 0
}

//
// Excludes the top and bottom 10% of outliers
//
func _SampleMean(v *[]float64) (r float64) {
	count := len(*v)
	low := int(float64(count) * 0.1)
	high := int(float64(count) * 0.9)
	for _, v := range (*v)[low : high-1] {
		r += v
	}
	r /= float64(high - low)

	return r
}

//
// Logs the benchmark trimmed mean runtime
//
func (self *Benchmark) Report(label string) {
	r := 0.0
	sort.Float64s(self.reads)
	if self.readno <= 2 {
		r = self.reads[0]
	} else {
		// ignore the outliers
		r = _SampleMean(&self.reads)
	}

	fmt.Printf("%10s %s\n", _formatDuration(r), label)
}
