package main

import "fmt"

// Function to demonstrate Bitwise AND
func myAND(a, b int) {
	fmt.Printf("a & b  = %d\n", a&b)
}

// Function to demonstrate Bitwise OR
func myOR(a, b int) {
	fmt.Printf("a | b  = %d\n", a|b)
}

// Function to demonstrate Bitwise XOR
func myXor(a, b int) {
	fmt.Printf("a ^ b  = %d\n", a^b)
}

// Function to demonstrate Bitwise NOT
func myNOT(a, b int) {
	fmt.Printf("^a = %d\n", ^a)
	fmt.Printf("^b = %d\n", ^b)
}

// Function to demonstrate Bitwise Shift
func myShift(a int) {
	fmt.Printf("a << 1 = %d\n", a<<1)
	fmt.Printf("a >> 1 = %d\n", a>>1)
}

// Function to demonstrate Assignment Operators
func myAssignment(a, b int) {
	fmt.Println("\n--- Assignment Operations ---")

	a += b
	fmt.Printf("a += b ➜ %d\n", a)

	a -= b
	fmt.Printf("a -= b ➜ %d\n", a)

	a *= b
	fmt.Printf("a *= b ➜ %d\n", a)

	if b != 0 {
		a /= b
		fmt.Printf("a /= b ➜ %d\n", a)

		a %= b
		fmt.Printf("a %%= b ➜ %d\n", a)
	}
}

func main() {
	var a, b int

	fmt.Print("Enter first integer (a): ")
	fmt.Scanln(&a)
	fmt.Print("Enter second integer (b): ")
	fmt.Scanln(&b)

	fmt.Println("\n--- Bitwise Operations ---")
	myAND(a, b)
	myOR(a, b)
	myXor(a, b)
	myNOT(a, b)
	myShift(a)

	myAssignment(a, b)
}
