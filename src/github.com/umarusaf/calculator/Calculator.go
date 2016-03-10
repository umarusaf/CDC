package main

import (
	"fmt"
	"os"
)

func add(x, y int) int {
	return x + y
}
func subtract(x, y int) int {
	return x - y
}
func multiply(x, y int) int {
	return x * y
}
func division(x, y int) int {
	var result int
	if y != 0 {
		result = x / y
	} else {
		result = 0
	}
	return result
}
func main() {
	var x, y, choice int
	for {
		fmt.Printf("Enter x:")
		fmt.Scan(&x)
		fmt.Printf("Enter y:")
		fmt.Scan(&y)
		fmt.Printf("Press 1 to perform addition\nPress 2 to perform subtract\nPress 3 to perform multiplication\nPress 4 to perform division\nPress 5 to exit\n")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Printf("Result:%d\n", add(x, y))
		case 2:
			fmt.Printf("Result:%d\n", subtract(x, y))
		case 3:
			fmt.Printf("Result:%d\n", multiply(x, y))
		case 4:
			{
				if y != 0 {
					fmt.Printf("Result:%d\n", division(x, y))
				} else {
					fmt.Printf("Invalid Input.\n")
				}
			}
		case 5:
			os.Exit(0)

		default:
			fmt.Printf("Invalid Input.\n")

		}
		fmt.Scanln()
	}

}
