// Copyright 2016, Marc Lavergne <mlavergn@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package goml

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
)

//
// Loads a matlab / comma delimited dataset into a matrix.
// Assumptions: float64
//
func Load(filePath string) (matrix Matrix) {
	// read the first line to introspect the column count
	matrix = nil

	usr, err := user.Current()
	if err == nil {
		path := usr.HomeDir + "/" + filePath

		_, err = os.Stat(path)
		if os.IsNotExist(err) {
			return matrix
		}

		bytes, err := ioutil.ReadFile(path)
		rawData := strings.TrimSpace(string(bytes))

		if err == nil {
			lines := strings.Split(rawData, "\n")
			rows := len(lines)
			cols := 0
			for row, line := range lines {
				data := strings.Split(line, ",")
				if cols == 0 {
					// first pass
					cols = len(data)
					matrix = NewMatrix(rows, cols)
				}
				for col, val := range data {
					fval, _ := strconv.ParseFloat(val, 64)
					matrix[row][col] = fval
				}
			}
		} else {
			log.Fatal(err)
		}
	}

	return matrix
}
