package main

import (
	"fmt"
	"sort"
	"strings"
)

func permutations(str []string) []string {
	var results []string
	shuffle(str, 0, &results)
	results = removeDuplicate(results)
	return results
}

func shuffle(str []string, index int, results *[]string) {
	if index == len(str)-1 {
		*results = append(*results, strings.Join(str, ""))
		return
	}

	for i := index; i < len(str); i++ {
		str[index], str[i] = str[i], str[index]
		shuffle(str, index+1, results)
		str[index], str[i] = str[i], str[index]
	}
}

func removeDuplicate(str []string) []string {
	sort.Strings(str)
	size := len(str)
	j := 1
	for i := 1; i < size; i++ {
		if str[i] != str[j-1] {
			str[j] = str[i]
			j++
		}
	}
	return str[:j]
}

func sliceText(text string) []string {
	return strings.Split(text, "")
}

func Manipulate(text []string) []string {
	fmt.Println("text: ", text)
	results := permutations(text)

	fmt.Println("results:", results)
	return results
}
