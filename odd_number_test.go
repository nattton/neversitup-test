package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindOddNumber(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		expect int
	}{
		{
			name:   "[7]",
			input:  []int{7},
			expect: 7,
		},
		{
			name:   "[0]",
			input:  []int{0},
			expect: 0,
		},
		{
			name:   "[0,1,0,1,0]",
			input:  []int{0, 1, 0, 1, 0},
			expect: 0,
		},
		{
			name:   "[1,2,2,3,3,3,4,3,3,3,2,2,1]",
			input:  []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			expect: 4,
		},
		{
			name:   "[1,2,2,3,3,3,4,3,3,3,2,2,1,4,5]",
			input:  []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1, 4, 5},
			expect: 5,
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			result := FindOddNumber(tc.input)
			require.Equal(t, result, tc.expect)
		})
	}
}
