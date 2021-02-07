package main

import "log"

func main() {
	log.Println(longestCommonPrefix([]string{"anable", "anagram", "anonym"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	checkerStrings := strs[0]
	for k, v := range strs {
		if len(v) < len(checkerStrings) && k > 0 {
			checkerStrings = v
		}
	}
	for i := len(checkerStrings); i >= 0; i-- {
		counter := 0
		for _, v := range strs {
			if v[:i] != checkerStrings[:i] {
				break
			} else {
				counter++
			}
		}
		if counter == len(strs) {
			return checkerStrings[:i]
		}
	}
	return ""
}
