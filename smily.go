package main

import "slices"

var eyes = []rune{':', ';'}
var noses = []rune{'-', '~'}
var mouths = []rune{')', 'D'}

func validSmily(emoticon string) bool {
	sliceEmo := []rune(emoticon)
	switch len(sliceEmo) {
	case 2:
		return slices.Contains(eyes, sliceEmo[0]) &&
			slices.Contains(mouths, sliceEmo[1])
	case 3:
		return slices.Contains(eyes, sliceEmo[0]) &&
			slices.Contains(noses, sliceEmo[1]) &&
			slices.Contains(mouths, sliceEmo[2])
	}

	return false
}

func CountSmilyFace(text []string) int {
	count := 0
	for _, v := range text {
		if validSmily(v) {
			count++
		}
	}
	return count
}
