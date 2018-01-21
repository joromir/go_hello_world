package main

import "fmt"

func hi() {
    fmt.Printf("Lalallalal\n")
}

func my_sum(operand_one int, operant_two int) int {
    return operand_one + operant_two
}

func main() {
    hi()

    fmt.Print(my_sum(12, 12))
    fmt.Printf("\n\n\n")
}
