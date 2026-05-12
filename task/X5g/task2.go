package main

import "fmt"

func reverse(sw []int) {
	l, r := 0, len(sw)-1

	for l < r {
		sw[l], sw[r] = sw[r], sw[l]
		l++
		r--
	}
}

func reverse2(sw []int) {
	for i, j := 0, len(sw)-1; i < j; i, j = i+1, j-1 {
		sw[i], sw[j] = sw[j], sw[i]
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	reverse(arr)

	fmt.Println(arr)
}
