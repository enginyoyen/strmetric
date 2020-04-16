/*
 * Copyright (c) 2020 Engin Yöyen.
 * Use of this source code is governed by an MIT
 * license that can be found in the LICENSE file.
 */

package strmetric

import (
	"math"
	"testing"
)

func TestDiceSorensonMetric(t *testing.T) {
	var testCases = []struct {
		a          string
		b          string
		similarity float64
	}{
		{"identical strings", "identical strings", 1.00},
		{"ACCESSARY", "ACCESSORY", 0.75},
		{"world", "arlon", 0.25},
		{"night", "nacht", 0.25},
	}

	for _, test := range testCases {
		similarity, _ := DiceSorensonMetric(test.a, test.b)
		if similarity != test.similarity {
			t.Errorf("Expected %f, got %f, for inputs:  %s & %s", test.similarity, similarity, test.a, test.b)
		}
	}
}

func TestDiceSorensonMetricError(t *testing.T) {
	var testCases = []struct {
		a          string
		b          string
		similarity float64
	}{
		{"", "", math.NaN()},
		{"", "ac", math.NaN()},
	}

	for _, test := range testCases {
		similarity, err := DiceSorensonMetric(test.a, test.b)
		if err == nil {
			t.Errorf("Expected error on empty strings but got similatiry result as  %f", similarity)
		}
	}
}
