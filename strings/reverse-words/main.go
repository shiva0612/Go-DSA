package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "my name is shiva"
	reverse1(a)
	reverse2(a)
	reverse3(a)
}

func reverse1(a string) {
	s := strings.Fields(a)
	l, h := 0, len(s)-1
	for l <= h {
		s[l], s[h] = s[h], s[l]
		l++
		h--
	}
	fmt.Println(s)
}

func reverse2(a string) {
	stack := []string{}
	temp := ""
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == byte(' ') {
			stack = append(stack, temp)
			temp = ""
			continue
		}
		temp = string(a[i]) + temp
	}
	if temp != "" {
		stack = append(stack, temp)
	}
	fmt.Println(stack)
}

func reverse3(a string) {
	var temp, ans string
	for _, v := range a {
		if v != rune(' ') {
			temp += string(v)
		} else {

			ans = temp + " " + ans
			temp = ""
		}
	}
	if temp != "" {
		ans = temp + " " + ans

	}
	fmt.Println(ans)
}
