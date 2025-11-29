package main

import (
	"fmt"
)

// Functions for each operation
func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func mod(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("modulo by zero")
	}
	return a % b, nil
}

func main() {
	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("1) Add")
		fmt.Println("2) Sub")
		fmt.Println("3) Mul")
		fmt.Println("4) Div")
		fmt.Println("5) Mod")
		fmt.Println("6) Exit")

		var choice int
		fmt.Print("Choose: ")
		fmt.Scanln(&choice)

		if choice == 6 {
			fmt.Println("Exiting calculator... Goodbye!")
			break
		}

		var a, b int
		fmt.Print("Enter value A: ")
		fmt.Scanln(&a)
		fmt.Print("Enter value B: ")
		fmt.Scanln(&b)

		switch choice {
		case 1:
			fmt.Printf(">>>>> Result: %d\n", add(a, b))
		case 2:
			fmt.Printf(">>>>> Result: %d\n", sub(a, b))
		case 3:
			fmt.Printf(">>>>> Result: %d\n", mul(a, b))
		case 4:
			result, err := div(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf(">>>>> Result: %d\n", result)
			}
		case 5:
			result, err := mod(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("\n>>>>> Result: %d\n", result)
			}
		default:
			fmt.Println("Invalid option! Please choose a number between 1 and 6.")
		}
	}
}
