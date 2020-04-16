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

// Jaro Similarity is the edit distance between two strings.
// See https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance
func JaroMetric(a, b string) (float64, error) {
	if len(a) == 0 || len(b) == 0 {
		return math.NaN(), fmt.Errorf("length of the strings can not be equal to 0 a=%d, b=%d", len(a), len(b))
	}
	if a == b {
		return 1.0, nil
	}
	// compute the window of possible matches
	matchDistance := (max(len(a), len(b)) / 2) - 1

	primary := make([]bool, len(a))
	secondary := make([]bool, len(b))

	matches := 0
	transportation := 0

	for i := 0; i < len(a); i++ {
		start := max(0, i-matchDistance)
		end := min(len(b), i+matchDistance+1)

		for j := start; j < end; j++ {
			// if it is already processed or if it is not equal continue
			if secondary[j] || a[i] != b[j] {
				continue
			}
			primary[i], secondary[j] = true, true
			matches++
			break
		}
	}
	// if there are 0 matches, no need further computation
	if matches == 0 {
		return 0.0, nil
	}

	k := 0
	for i := 0; i < len(a); i++ {
		if !primary[i] {
			continue
		}
		for !secondary[k] {
			k++
		}
		if a[i] != b[k] {
			transportation++
		}
		k++
	}
	score := score(matches, transportation, len(a), len(b))
	return score, nil
}

// score calculated as
// score = (1/3) * {(m/s1) + (m/s1) + (m-t)/m }
// m is the number of matching characters
// t is half the number of transpositions
// where |s1| and |s2| is the length of string s1 and s2 respectively
func score(matches, transportation, lengthOfA, lengthOfB int) float64 {
	transportation /= 2 // half the number of transpositions
	matchingCharsOfA := float64(matches) / float64(lengthOfA)
	matchingCharsOfB := float64(matches) / float64(lengthOfB)
	transposition := (float64(matches) - float64(transportation)) / float64(matches)
	return (matchingCharsOfA + matchingCharsOfB + transposition) / 3.0
}

