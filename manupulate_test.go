package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceText(t *testing.T) {
	assert := assert.New(t)
	assert.Equal([]string{"a"}, sliceText("a"))
	assert.Equal([]string{"a", "b"}, sliceText("ab"))
	assert.Equal([]string{"a", "b", "c"}, sliceText("abc"))
}

func TestRemoveDuplicate(t *testing.T) {
	assert := assert.New(t)
	results := removeDuplicate([]string{"a", "a", "a", "ab", "ab", "aab", "aab", "aaab", "aaab"})
	expect := []string{"a", "ab", "aab", "aaab"}
	assert.Len(results, len(expect))
	assert.ElementsMatch(results, expect)
}

func TestManipulate(t *testing.T) {
	testCases := []struct {
		name   string
		input  []string
		expect []string
	}{
		{
			name:   "a",
			input:  sliceText("a"),
			expect: []string{"a"},
		},
		{
			name:   "ab",
			input:  sliceText("ab"),
			expect: []string{"ab", "ba"},
		},
		{
			name:   "abc",
			input:  sliceText("abc"),
			expect: []string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
		{
			name:   "aabb",
			input:  sliceText("aabb"),
			expect: []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
		{
			name:   "aaabb",
			input:  sliceText("aaabb"),
			expect: []string{"aaabb", "aabab", "aabba", "abaab", "ababa", "abbaa", "baaab", "baaba", "babaa", "bbaaa"},
		},
		{
			name:   "abcd",
			input:  sliceText("abcd"),
			expect: []string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb", "bacd", "badc", "bcad", "bcda", "bdac", "bdca", "cabd", "cadb", "cbad", "cbda", "cdab", "cdba", "dabc", "dacb", "dbac", "dbca", "dcab", "dcba"},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			results := Manipulate(tc.input)
			require.Len(t, results, len(tc.expect))
			require.ElementsMatch(t, results, tc.expect)
		})
	}
}
