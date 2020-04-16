/*
 * Copyright (c) 2020 Engin Yöyen.
 * Use of this source code is governed by an MIT
 * license that can be found in the LICENSE file.
 */

package strmetric

import (
	"fmt"
	"testing"
)

func TestJaroMetric(t *testing.T) {
	var testCases = []struct {
		a        string
		b        string
		distance float64
	}{
		{"identical", "identical", 1.00},
		{"martha", "marhta", 0.944444},
		{"DWAYNE", "DUANE", 0.822222},
		{"JELLYFISH", "SMELLYFISH", 0.896296},
	}

	for _, test := range testCases {
		distance, _ := JaroMetric(test.a, test.b)
		if fmt.Sprintf("%.6f", distance) != fmt.Sprintf("%.6f", test.distance) {
			t.Errorf("Expected %f, got %f, for inputs:  %s & %s", test.distance, distance, test.a, test.b)
		}
	}
}
