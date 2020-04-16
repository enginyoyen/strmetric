/*
 * Copyright (c) 2020 Engin Yöyen.
 * Use of this source code is governed by an MIT
 * license that can be found in the LICENSE file.
 */

package strmetric

// Jaro-Winkler Similarity is the edit distance between two strings
// and uses prefix scale to get accurate results when they share common prefix
// See https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance
func JaroWinklerMetric(a, b string) (float64, error) {
	result, err := JaroMetric(a, b)
	if err != nil {
		return result, err
	}

	if result == 0.0 || result == 1.0 {
		return result, nil
	}

	prefix := 0
	if result > 0.7 {
		minStrLength := min(len(a), len(b))

		for i := 0; i < minStrLength; i++ {
			// only prefix of the string has tobe taken into an account
			// therefore as soon as there is no match break the loop
			if a[i] == b[i] {
				prefix++
			} else {
				break
			}
		}
		// reduce the prefix number to 4 if it is greater
		// max of 4 characters are allowed in prefix
		prefix = min(4, prefix)

		result = result + 0.1*float64(prefix)*(1.0-result)
	}
	return result, nil
}
