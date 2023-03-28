package main

import (
	"fmt"
)

func main() {
	fmt.Println("-----Игра «Угадай число»-----")

	fmt.Println("Загадайте число от 1 до 10.")

	fmt.Println("Ваше число больше, меньше или равно 3?")
	var answer string
	fmt.Scan(&answer)

	if answer == "меньше" {
		fmt.Println("Ваше число меньше или равно 2?.")
		fmt.Scan(&answer)
		if answer == "меньше" {
			fmt.Println("Ваше число 1")
		} else if answer == "равно" {
			fmt.Println("Ваше число 2.")
		}
	} else if answer == "равно" {
		fmt.Println("Ваше число 3.")
	} else {
		fmt.Println("Ваше число больше, меньше или равно 5?")
		fmt.Scan(&answer)
		if answer == "меньше" {
			fmt.Println("Ваше число 4.")
		} else if answer == "равно" {
			fmt.Println("Ваше число 5.")
		} else {
			fmt.Println("Ваше число больше, меньше или равно 7?")
			fmt.Scan(&answer)
			if answer == "меньше" {
				fmt.Println("Ваше число 6.")
			} else if answer == "равно" {
				fmt.Println("Ваше число 7.")
			} else {
				fmt.Println("Ваше число больше, меньше или равно 9?")
				fmt.Scan(&answer)
				if answer == "меньше" {
					fmt.Println("Ваше число 8.")
				} else if answer == "равно" {
					fmt.Println("Ваше число 9.")
				} else {
					fmt.Println("Ваше число 10.")
				}
			}
		}
	}
}
