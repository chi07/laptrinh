package main

import (
	"fmt"
	"math"
)

func main() {
	n := 200
	K := 10

	a := math.Sqrt(float64(n))
	fmt.Println(a)

	for i := 0; i < n; i++ {
		for j := 1; j < K; j++ {
			fmt.Println("i * j = ", i*j)
		}
	}

	m := 3
	y := 2022

	m = 5
	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Printf("thang %v co 31 ngay", m)
	case 4, 6, 9, 11:
		fmt.Printf("thang %v co 30 ngay", m)
	case 2:
		if y == 0 {
			fmt.Printf("hay nhap vao nam ")
		} else {
			if y != 0 && y%4 == 0 {
				fmt.Printf("thang %v nam %v co 29 ngay", m, y)
			} else {
				fmt.Printf("thanng %v nam %v co 28 ngay ", m, y)
			}
		}
	default:
		fmt.Println("thang nay khong hop le")
	}
}
