package main

import "regexp"

var cleanPattern = regexp.MustCompile("[^<>+-.,\\[\\]]")

func CleanInput(input string) string {
	return cleanPattern.ReplaceAllString(input, "")
}

func Validate(input string) bool {
	height := 0
	bmap := map[rune]int{'[': 1, ']': -1,}
	for _, r := range input {
		if r != '[' && r != ']' {
			continue
		}
		height = height + bmap[r]

		if height < 0 {
			return false
		}
	}
	return height == 0
}
