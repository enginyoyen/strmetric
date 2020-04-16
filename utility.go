/*
 * Copyright (c) 2020 Engin Yöyen.
 * Use of this source code is governed by an MIT
 * license that can be found in the LICENSE file.
 */

package strmetric

// Returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Returns the greater of two integers
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
