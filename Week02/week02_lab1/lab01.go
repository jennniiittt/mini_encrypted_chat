package main

import "fmt"

func main() {
	var a, b int

	//take two integer inputs
    fmt.Print("Enter the first integer: ")
	fmt.Scanln(&a)
	fmt.Print("Enter the second integer: ")
	fmt.Scanln(&b)

	fmt.Println("=====Values=====")
	fmt.Printf("a = %d, b = %d\n", a, b)

	
	// Assignment operations
    a = b
    fmt.Printf("\na = b  ➜ a = %d\n", a)

    a += b
    fmt.Printf("a += b ➜ a = %d\n", a)

    a -= b
    fmt.Printf("a -= b ➜ a = %d\n", a)

    a *= b
    fmt.Printf("a *= b ➜ a = %d\n", a)

	a /= b
	fmt.Printf("a /= b ➜ a = %d\n", a)

}
