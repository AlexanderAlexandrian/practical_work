package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----Проверить, что есть совпадающие числа-----")

	fmt.Println("Введите первое число:")
	var a int
	fmt.Scan(&a)

	fmt.Println("Введите второе число:")
	var b int
	fmt.Scan(&b)

	fmt.Println("Введите третье число:")
	var c int
	fmt.Scan(&c)

	if a == b || a == c || b == c {
		fmt.Println("Есть совпадения в числах")
	}
}
