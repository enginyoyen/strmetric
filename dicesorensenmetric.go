/*
 * Copyright (c) 2020 Engin Yöyen.
 * Use of this source code is governed by an MIT
 * license that can be found in the LICENSE file.
 */

package strmetric

import (
	"fmt"
	"math"
)

//Dice-Sørensen algorithm is used to gauge the similarity of two samples.
//See http://en.wikipedia.org/wiki/Dice%27s_coefficient
func DiceSorensonMetric(a, b string) (float64, error) {
	if len(a) < 2 || len(b) < 2 {
		return math.NaN(), fmt.Errorf("length of the strings are less than 2, it is not possible to compare a=%d, b=%d", len(a), len(b))
	}

	if a == b {
		return 1.00, nil
	}

	aBigramed := *nGrams(a, 2)
	bBigramed := *nGrams(b, 2)

	//total number of bi-grams of both inputs
	var maxInput = len(aBigramed) + len(bBigramed)
	var matchingPairs = 0
	for _, str := range aBigramed {
		if isMatch(str, bBigramed) {
			matchingPairs++
		}
	}
	// double the count og matchingPairs as bi-grams occurs once in each parameter(a,b)
	return float64(matchingPairs*2) / float64(maxInput), nil
}

//check whether given bi-gram exist in the provided slice
func isMatch(biGram string, source []string) bool {
	for i, s := range source {
		if biGram == s {
			source = append(source[:i], source[i+1:]...)
			return true
		}
	}
	return false
}

// Returns contiguous sequence of n items from given string s
func nGrams(s string, n int) *[]string {
	var result []string
	for i := 0; i < len(s)-n+1; i++ {
		result = append(result, s[i:i+n])
	}
	return &result
}
