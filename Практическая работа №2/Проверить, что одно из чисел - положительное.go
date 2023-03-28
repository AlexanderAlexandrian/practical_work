package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----Проверить,что одно из чисел - положительное----- ")

	fmt.Println("Введите первое число:")
	var a int
	fmt.Scan(&a)

	fmt.Println("Введите второе число:")
	var b int
	fmt.Scan(&b)

	fmt.Println("Введите третье число:")
	var c int
	fmt.Scan(&c)

	if a > 0 || b > 0 || c > 0 {
		fmt.Println("Присутствует положительное число")
	} else {
		fmt.Println("Отсутствует положительное число")
	}
}
