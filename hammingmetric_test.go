/*
 * Copyright (c) 2020 Engin Yöyen.
 * Use of this source code is governed by an MIT
 * license that can be found in the LICENSE file.
 */

package strmetric

import "testing"

func TestHammingMetric(t *testing.T) {
	var testCases = []struct {
		a        string
		b        string
		distance int
	}{
		{"abc", "abc", 0},
		{"0123", "0123", 0},
		{"0123", "0abc", 3},
		{"table", "chair", 5},
		{"101101", "100101", 1},
		{"11011001", "10011101", 2},
	}

	for _, test := range testCases {
		distance, _ := HammingMetric(test.a, test.b)
		if distance != test.distance {
			t.Errorf("Expected %d, got %d, for inputs:  %s & %s", test.distance, distance, test.a, test.b)
		}
	}
}
