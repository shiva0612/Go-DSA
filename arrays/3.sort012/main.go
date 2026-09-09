package main

func main() {
	ip := []int{2, 0, 2, 1, 1, 0}
	sort012(ip)
}

func sort012(a []int) {

	low, mid := 0, 0
	high := len(a) - 1

	for mid <= high {
		switch a[mid] {
		case 0:
			a[low], a[mid] = a[mid], a[low]
			low += 1
			mid += 1
		case 1:
			mid += 1
		case 2:
			a[mid], a[high] = a[high], a[mid]
			high -= 1
		}
	}
}
