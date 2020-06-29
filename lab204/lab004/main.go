package main

import (
	"github.com/go-playground/validator/v10"
	"log"
)

type Price struct {
	PriceMin int `validate:"gte=10"`
	PriceMax int `validate:"gtfield=PriceMin"`
}

func main() {
	validate := validator.New()

	p := &Price{
		PriceMin: 15,
		PriceMax: 16,
	}
	if err := validate.Struct(p); err != nil {
		log.Printf("validate error:%v", err)
	}

	log.Println("validate success")
}
