package main

import "fmt"

func hi() {
    fmt.Printf("Lalallalal\n")
}

func my_sum(operand_one int, operant_two int) int {
    return operand_one + operant_two
}

func factorial(number int) int {
  if number == 1 {
    return 1
  } else {
    return number * factorial(number - 1)
  }
}

func fibonacci(number int) int {
  if number == 1 || number == 2 {
    return 1
  } else {
    return fibonacci(number - 1) + fibonacci(number - 2)
  }
}

func main() {
    hi()

    fmt.Print(my_sum(12, 12))
    fmt.Printf("\n\n\n")

    fmt.Print(factorial(5))
    fmt.Printf("\n\n\n")

    fmt.Print(factorial(10))
    fmt.Printf("\n\n\n")

    fmt.Print(fibonacci(6))
    fmt.Printf("\n\n\n")
}
