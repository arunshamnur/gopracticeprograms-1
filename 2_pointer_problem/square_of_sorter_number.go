package main

import (
	"fmt"
	"math"
)

func main() {
	var a = []int{-7, -6, -5, 2, 3, 8}
	var i, j = 0, len(a) - 1
	var length = len(a) - 1
	res := make([]int, len(a))
	for i < j {
		if math.Abs(float64(a[i])) > math.Abs(float64(a[j])) {
			res[length] = a[i] * a[i]
			i++
		} else {
			res[length] = a[j] * a[j]
			j--
		}
		length--
	}
	fmt.Println(res)
}
