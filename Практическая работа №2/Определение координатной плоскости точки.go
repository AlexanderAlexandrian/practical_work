package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----Определение координатной плоскости точки-----")

	fmt.Println("Введите значение X:")
	var x int
	fmt.Scan(&x)

	fmt.Println("Введите значение Y:")
	var y int
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		fmt.Println("Точка принадлежит к первой координатной четверти.")
	} else if x > 0 && y < 0 {
		fmt.Println("Точка принадлежит ко второй координатной четверти.")
	} else if x < 0 && y < 0 {
		fmt.Println("Точка принадлежит к третьей координатной четверти.")
	} else if x < 0 && y > 0 {
		fmt.Println("Точка принадлежит к четвертой координатной четверти.")
	} else {
		fmt.Println("Точка находится на оси координат")
	}

}
