package linkedlist

import (
	"fmt"
)

func Check_revereBySize() {

	a := stringTOLList("13")
	a.revereBySize(3)
}
func Check_checkIfPalindrome() {
	a := stringTOLList("1241")
	ans := a.checkIfPalindrome2()
	fmt.Println("ans = ", ans)
}

func Check_mergeList() {
	a, b := stringTOLList("234"), stringTOLList("239")
	a.printlist()
	b.printlist()
	merged := mergeList(a, b)
	merged.printlist()
}

func Check_addNumbers() {
	a, b := stringTOLList("234"), stringTOLList("239")
	a.printlist()
	b.printlist()
	ans := additionOfNumbers(a, b)
	ans.printlist()
}

func Check_rotateByK() {
	a := stringTOLList("1")
	a.rotateByK(3)
	a.printlist()
}
