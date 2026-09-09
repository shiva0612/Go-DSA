package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 1
	close(ch)
	for range 3 {
		data, ok := <-ch
		if ok {
			fmt.Println(data)
		} else {
			fmt.Println("NOT OK")
		}
	}
}
