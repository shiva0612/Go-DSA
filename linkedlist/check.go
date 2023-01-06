package linkedlist

import "fmt"

func Check_revereBySize() {

	a := arrayTOList("13")
	a.revereBySize(3)
}
func Check_checkIfPalindrome() {
	a := arrayTOList("1241")
	ans := a.checkIfPalindrome2()
	fmt.Println("ans = ", ans)
}

func Check_mergeList() {
	a, b := arrayTOList("234"), arrayTOList("239")
	a.printlist()
	b.printlist()
	merged := mergeList(a, b)
	merged.printlist()
}

func Check_addNumbers() {
	a, b := arrayTOList("234"), arrayTOList("239")
	a.printlist()
	b.printlist()
	ans := addNumbers(a, b)
	ans.printlist()
}

func Check_rotateByK() {
	a := arrayTOList("1")
	a.rotateByK(3)
	a.printlist()
}
