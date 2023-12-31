package main

import (
  "fmt"
  "math"
  "errors"
)

func main() {
  fmt.Printf("Hello World\n")

  eatInPrice, takeOutPrice := taxPrice(float64(150))
  fmt.Printf("use takeOutPrice method\n")
  fmt.Printf("eatIn price %.0f\n", eatInPrice)
  fmt.Printf("takeOutPrice %.0f\n", takeOutPrice)

  taxType := "eatIn"
  // taxType := "takeOut"
  // taxType := ""
  price, err := taxPriceByTaxType(taxType, float64(100))
  if err == nil {
    fmt.Printf("use taxPriceByTaxType method\n")
    fmt.Printf(taxType + " price %.0f\n", price)
  } else {
    fmt.Println("エラー:", err)
  }
}

const (
	EAT_IN_TAX = 0.1
	TAKE_OUT_TAX = 0.08
)

func taxPrice(price float64) (float64, float64) {
 return math.Floor(price + price * EAT_IN_TAX), math.Floor(price + price * TAKE_OUT_TAX)
}

func taxPriceByTaxType(taxType string, price float64) (float64, error) {
  fmt.Printf("taxType：" + taxType + "\n")
  switch taxType {
	case "eatIn":
		return math.Floor(price + price * EAT_IN_TAX), nil
	case "takeOut":
		return math.Floor(price + price * TAKE_OUT_TAX), nil
	default:
    return 0, errors.New("税タイプが指定されていません")
  }
}
