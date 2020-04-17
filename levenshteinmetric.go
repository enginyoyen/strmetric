/*
 * Copyright (c) 2020 Engin Yöyen.
 * Use of this source code is governed by an MIT
 * license that can be found in the LICENSE file.
 */

package strmetric

import (
	"fmt"
	"strings"
)

// Levenshtein distance is a string metric for measuring the difference between two sequences.
// Used for spell checkers, error detection/correction in optical character recognition, etc.
// See http://en.wikipedia.org/wiki/Levenshtein_distance
func LevenshteinMetric(src, trg string) (int, error) {
	if len(src) == 0 || len(trg) == 0 {
		return -1, fmt.Errorf("length of the strings can not be equal to 0 a=%d, b=%d", len(src), len(trg))
	}

	if src == trg {
		return 0, nil
	}
	//lowercase for comparision
	_src := strings.ToLower(src)
	_trg := strings.ToLower(trg)
	srcLength := len(src)
	trgLength := len(trg)

	// create the matrix
	costMatrix := make([][]int, srcLength+1)
	for i := range costMatrix {
		costMatrix[i] = make([]int, trgLength+1)
	}

	// add source
	for i := range costMatrix {
		costMatrix[i][0] = i
	}
	// add target
	for j := range costMatrix[0] {
		costMatrix[0][j] = j
	}

	substitutionCost := 0
	// calculate the substitutionCost
	for i := 1; i <= srcLength; i++ {
		for j := 1; j <= trgLength; j++ {
			if _src[i-1] == _trg[j-1] {
				substitutionCost = 0
			} else {
				substitutionCost = 1
			}
			deletion := costMatrix[i-1][j] + 1                      // deletion (left)
			insertion := costMatrix[i][j-1] + 1                     // insertion (up)
			substitution := costMatrix[i-1][j-1] + substitutionCost // substitution (left-up)
			costMatrix[i][j] = min(min(deletion, insertion), substitution)
		}
	}
	result := costMatrix[srcLength][trgLength]
	return result, nil
}
