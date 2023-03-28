package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----Сумма без сдачи-----")

	fmt.Println("Введите стоимость товара:")
	var price int
	fmt.Scan(&price)

	fmt.Println("Введите номинал первой монеты:")
	var a int
	fmt.Scan(&a)

	fmt.Println("Введите номинал второй монеты:")
	var b int
	fmt.Scan(&b)

	fmt.Println("Введите номинал третьей монеты:")
	var c int
	fmt.Scan(&c)

	if price == a || price == b || price == c || price == a+b || price == a+b+c || price == a+c || price == b+c {
		fmt.Println("Товар может быть оплачен без сдачи.")
	} else if price > a+b+c {
		fmt.Println("У вас недостаточно средств.")
	} else {
		fmt.Println("Товар не может быть оплачен без сдачи")
	}
}
