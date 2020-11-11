package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"log"
)

func main() {
	//demo()
	sample()
}

func demo() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromInt(3)

	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875
}

func sample() {
	price1 := decimal.NewFromFloat(float64(6) / 100)
	log.Println("price1=", price1)

	price2 := decimal.NewFromFloat(float64(1234) / 100)
	log.Println("price2=", price2)

	total := price1.Add(price2)
	log.Println("total=", total)
	total02 := total.Mul(decimal.NewFromInt(100))
	log.Println("total02=", total02)
	log.Println(total02.IntPart())
}
