package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortWords(s string) string {
	var words = []string{}

	var n, initial int
	A := rune('A')
	Z := rune('Z')
	for i, c := range s {
		if c >= A && c <= Z {
			n++
			if n%2 == 0 {
				words = append(words, s[initial:i+1])
				initial = i + 1
			}
		}
	}

	sort.Slice(words, func(i, j int) bool {
		return strings.ToLower(words[i]) < strings.ToLower(words[j])
	})

	return strings.Join(words, "")
}

func main() {
	var s string
	_, e := fmt.Scanf("%s", &s)
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(SortWords(s))
}
