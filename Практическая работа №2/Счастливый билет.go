package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----Счастливый билет-----")

	fmt.Println("Введите номер билета:")
	var number int
	fmt.Scan(&number)

	a := number / 100 / 10
	b := number / 100 % 10
	c := number % 100 / 10
	d := number % 100 % 10

	if number < 1000 || number > 9999 {
		fmt.Println("Вы ввели неверный номер билета. Повторите попытку.")
	} else if a == d && b == c {
		fmt.Println("Ваш билет зеркальный")
	} else if a+b == c+d {
		fmt.Println("Ваш билет счастливый.")
	} else {
		fmt.Println("Ваш билет обычный.")
	}
}
