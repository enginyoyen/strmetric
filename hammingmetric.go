package strmetric

import (
	"fmt"
)

// The Hamming distance is between two strings of equal length is the number of positions at which the corresponding
// symbols are different. It measures the minimum number of substitutions required to change one string
// into the other, or the minimum number of errors that. See https://en.wikipedia.org/wiki/Hamming_distance
func HammingMetric(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("Two strings must be at equal length for Hamming distance calculations. They are not :  a=%d, b=%d", len(a), len(b))
	}
	if a == b {
		return 0, nil
	}

	var distance = 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
