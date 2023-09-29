package main

import (
	"testing"

	"github.com/go-playground/assert"
)

func TestBS(t *testing.T) {
	// assert.Equal(t, first_last_occurance([]int{3, 4, 13, 13, 13, 20, 40}, 13), []int{2, 4})
	// assert.Equal(t, count_occurances([]int{6, 7, 8, 1, 2, 3, 4, 5}, 13), 3)
	// assert.Equal(t, rotated_at([]int{6, 7, 8, 1, 2, 3, 4}), []int{3, 1})
	// assert.Equal(t, find_single_element([]int{1, 1, 2, 2, 3, 4, 4, 5, 5}), 3)
	// assert.Equal(t, find_peak_bs([]int{1, 2, 3, 2, 1}), 3)
	// assert.Equal(t, sq_root(37), 6)
	// assert.Equal(t, nthRoot(4, 16), 2)
	// assert.Equal(t, minRateToEatBananas([]int{3, 6, 7, 11}, 8), 4)
	// assert.Equal(t, noOfBoq_bs([]int{7, 7, 7, 7, 13, 11, 12, 7}, 2, 3), 12)
	assert.Equal(t, min_ship_days([]int{5, 4, 5, 2, 3, 4, 5, 6}, 5), 9)
}
