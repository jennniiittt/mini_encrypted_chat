package main

import "fmt"

func main() {
    var a, b int

    // Take two integer inputs
    fmt.Print("Enter first number (a): "); fmt.Scanln(&a)
    fmt.Print("Enter second number (b): "); fmt.Scanln(&b)

    // Logical operators demonstration
    fmt.Println("\n====== Logical Operations ======")

    bothPositive := (a > 0) && (b > 0)
    fmt.Printf("Both a and b are positive: %t\n", bothPositive)

    oneGreater := (a > b) || (b > a)
    fmt.Printf("At least one is greater than the other: %t\n", oneGreater)

    notEqual := !(a == b)
    fmt.Printf("a and b are not equal: %t\n", notEqual)
}
