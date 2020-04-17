/*
 * Copyright (c) 2020 Engin Yöyen.
 * Use of this source code is governed by an MIT
 * license that can be found in the LICENSE file.
 */

package strmetric

import "testing"

func TestLevenshteinMetric(t *testing.T) {
	var testCases = []struct {
		src      string
		trg      string
		distance int
	}{
		{"identical", "identical", 0},
		{"kitten", "sitting", 3},
		{"saturday", "sunday", 3},
		{"JELLYFISH", "SMELLYFISH", 2},
	}

	for _, test := range testCases {
		distance, _ := LevenshteinMetric(test.src, test.trg)
		if distance != test.distance {
			t.Errorf("Expected %d, got %d, for inputs:  %s & %s", test.distance, distance, test.src, test.trg)
		}
	}

}
