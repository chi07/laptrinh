package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64 = 1
	var b float64 = 2
	var c float64 = -3
	var delta float64

	delta = b*b - 4*a*c

	if delta < 0 {
		fmt.Println("Phương trình vô nghiệm")
	} else if delta == 0 {
		x := -b / (2 * a)
		fmt.Println("Phương trình có nghiệm duy nhất: ", x)
	} else {
		x1 := (-b + math.Sqrt(delta)) / 2 * a
		x2 := (-b - math.Sqrt(delta)) / 2 * a
		fmt.Println("Phương trình có 2 nghiệm phân biệt x1 =", x1, " và x2 =", x2)
	}
}
