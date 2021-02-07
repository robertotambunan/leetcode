package main

import "log"

func main() {
	log.Println(isValid("()"))
}

func isValid(s string) bool {
	openBracket := make(map[string]string)
	openBracket["("] = ")"
	openBracket["{"] = "}"
	openBracket["["] = "]"
	var stack string

	for _, v := range s {
		if _, ok := openBracket[string(v)]; ok {
			stack = push(stack, string(v))
		} else {
			var popItem string
			stack, popItem = pop(stack)
			if pair, ok := openBracket[popItem]; ok && string(v) == pair {
				continue
			} else {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func push(stack string, newChar string) string {
	return stack + newChar
}
func pop(stack string) (currentStack string, item string) {
	item = stack[len(stack)-1:]
	currentStack = stack[:len(stack)-1]
	return
}
