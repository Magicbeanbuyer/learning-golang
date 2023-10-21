package main

import (
    "fmt"

    "github.com/shopspring/decimal"
)

func main() {
    amount, _ := decimal.NewFromString("1.2")
    fmt.Println(amount)
}