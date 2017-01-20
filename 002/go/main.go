package main

import "fmt"

func main() {
	var num1, num2, total int = 1, 2, 0

	for num1 <= 4000000 {
		temp := num2
		num2 += num1
		num1 = temp
		if num1%2 == 0 {
			total += num1
		}
	}

	fmt.Println(total)
}
