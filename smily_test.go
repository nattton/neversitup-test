package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountSmilyFace(t *testing.T) {
	testCases := []struct {
		name   string
		input  []string
		expect int
	}{
		{
			name:   "[':)', ';(', ';}', ':-D']",
			input:  []string{":)", ";(", ";}", ":-D"},
			expect: 2,
		},
		{
			name:   "[';D', ':-(', ':-)', ';~)']",
			input:  []string{";D", ":-(", ":-)", ";~)"},
			expect: 3,
		},
		{
			name:   "[';]', ':[', ';*', ':$', ';-D']",
			input:  []string{";]", ":[", ";*", ":$", ";-D"},
			expect: 1,
		},
		{
			name:   "[':-)', ':~D', ';)', ':~)', ';-D']",
			input:  []string{":-)", ":~D", ";)", ":~)", ";-D"},
			expect: 5,
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			result := CountSmilyFace(tc.input)
			require.Equal(t, tc.expect, result)
		})
	}
}
