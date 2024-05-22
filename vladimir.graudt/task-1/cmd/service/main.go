package main

import (
  "fmt"
  "os"
)

func err(e string) {
  fmt.Println("Error:", e)
  os.Exit(1)
}

func readNum() float32 {
  var num float32
  if n, _ := fmt.Scanln(&num); n == 0 {
    /* flush stdin */
    var dummy string
    fmt.Scanln(&dummy)
    err("not a valid number")
  }
  return num
}

func isValidOp(op string) bool {
  return op == "+" || op == "-" || op == "*" || op == "/"
}


func main() {
  var op string

  fmt.Print("Enter first operand: ")
  x := readNum()

  fmt.Print("Enter operation (+ - * /): ")
  fmt.Scan(&op)

  if !isValidOp(op) {
    err("invalid operation, must be +, -, *, /")
  }

  fmt.Print("Enter second operand: ")
  y := readNum()

  switch op {
  case "+": fmt.Println("Result:", x + y)
  case "-": fmt.Println("Result:", x - y)
  case "*": fmt.Println("Result:", x * y)
  case "/":
    if y == 0 {
      err("division by zero")
    } else {
      fmt.Println("Result:", x / y)
    }
  }
}
