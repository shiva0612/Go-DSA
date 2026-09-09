package main

import "strings"

func main() {
	isValid("[]")
}
func isValid(a string) bool {

	stack := []int{}
	openb, closeb := "([{", ")]}"

	for _, v := range a {
		if index := strings.IndexRune(openb, v); index != -1 {
			stack = append(stack, index)
		} else {
			if len(stack) == 0 {
				return false
			}
			index := strings.IndexRune(closeb, v)
			if stack[len(stack)-1] != index {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
