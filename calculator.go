package main

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

var (
  x, y, result float64 = 0, 0, 0
  operator string = ""
  scanner = bufio.NewScanner(os.Stdin)
)

func calculate(x float64, operator string, y float64) float64 {
  switch operator {
    case "+": result = x + y
    case "-": result = x - y
    case "*": result = x * y
    case "/": result = x / y
  }

  return result
}

func show(result float64) {
  fmt.Println(result)
}

func main() {
  for {
    fmt.Println("pls insert first value")
    scanner.Scan()
    i, err1 := strconv.ParseFloat(scanner.Text(), 64)
    if err1 == nil {
      x = i
    }

    fmt.Println("pls insert operator")
    scanner.Scan()
    operator = scanner.Text()

    fmt.Println("pls insert second value")
    scanner.Scan()
    j, err2 := strconv.ParseFloat(scanner.Text(), 64)
    if err2 == nil {
      y = j
    }

    show(calculate(x, operator, y))
  }
}
