// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	"math"
	"math/rand"
	"time"
)

//
// Constructs a set filled with random float whole values with
// range min <= value <= max
//
func RandomWholeFloatSet(l int, min float64, max float64) []float64 {
	r := make([]float64, l)

	s := max - min + 1

	rand.Seed(time.Now().UTC().UnixNano())
	for i := range r {
		r[i] = float64(int(rand.Float64()*s + min))
	}

	return r
}

//
// Constructs a set filled with random float values with
// range min <= value <= max
//
func RandomFloatSet(l int, min float64, max float64) []float64 {
	r := make([]float64, l)

	s := max - min + 1

	rand.Seed(time.Now().UTC().UnixNano())
	for i := range r {
		v := rand.Float64()*s + min
		if v > max {
			// choose floor in the event we wind up with
			// a fractional over max
			v = math.Floor(v)
		}
		r[i] = v
	}

	return r
}

//
// Returns two sets of randomly selected values, a training set and testing set
//
func RandomTrainTestSets(inarray []float64, percent float64) (train []float64, test []float64) {
	count := float64(len(inarray))

	// round to fit the set size accurately
	traincount := int(math.Floor(count*percent + .5))
	train = make([]float64, traincount)
	traintop := 0

	testcount := int(count) - traincount
	test = make([]float64, testcount)
	testtop := 0

	rand.Seed(time.Now().UTC().UnixNano())

	// the logic guarantees we will have two arrays of
	// the expected sizes
	for _, v := range inarray {
		if rand.Float64() <= percent && traintop < traincount {
			train[traintop] = v
			traintop += 1
		} else if testtop < testcount {
			test[testtop] = v
			testtop += 1
		} else {
			train[traintop] = v
			traintop += 1
		}
	}

	return train, test
}

//
// Returns three sets of randomly selected values, a training set, validation set and testing set
//
func RandomTrainValidationTestSets(inarray []float64, validationpercent float64, testpercent float64) (train []float64, validation []float64, test []float64) {
	count := float64(len(inarray))

	trainpercent := 1.0 - validationpercent - testpercent

	// round to fit the set size accurately
	traincount := int(math.Floor(count*trainpercent + .5))
	train = make([]float64, traincount)
	traintop := 0

	validationcount := int(math.Floor(count*validationpercent + .5))
	validation = make([]float64, validationcount)
	validationtop := 0

	testcount := int(count) - traincount - validationcount
	test = make([]float64, testcount)
	testtop := 0

	rand.Seed(time.Now().UTC().UnixNano())

	// the logic guarantees we will have three arrays of
	// the expected sizes
	for _, v := range inarray {
		r := rand.Float64()
		if r <= trainpercent && traintop < traincount {
			train[traintop] = v
			traintop += 1
		} else if r > trainpercent && r <= trainpercent+validationpercent && validationtop < validationcount {
			validation[validationtop] = v
			validationtop += 1
		} else if testtop < testcount {
			test[testtop] = v
			testtop += 1
		} else if traintop < traincount {
			train[traintop] = v
			traintop += 1
		} else {
			validation[validationtop] = v
			validationtop += 1
		}
	}

	return train, validation, test
}
