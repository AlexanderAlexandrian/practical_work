package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("-----Решение квадратного уравнения-----")

	fmt.Println("Введите коэфициент a:")
	var a float64
	fmt.Scan(&a)

	fmt.Println("Введите коэфициент b:")
	var b float64
	fmt.Scan(&b)

	fmt.Println("Введите коэфициент c:")
	var c float64
	fmt.Scan(&c)

	D := math.Pow(b, 2) - 4*a*c

	if D > 0 {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println("Имеет два различных корня:", x1, "и", x2)
	} else if D == 0 {
		x := -b / (2 * a)
		fmt.Println("Имеет один корень:", x)
	} else {
		fmt.Println("Корней нет.")
	}
}
