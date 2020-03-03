package main

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// User contains user information
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {
	validate = validator.New()

	validateStruct()
	validateVariable()
}

func validateStruct() {
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Println("error", err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			log.Println(err.Namespace())
			log.Println(err.Field())
			log.Println(err.StructNamespace())
			log.Println(err.StructField())
			log.Println(err.Tag())
			log.Println(err.ActualTag())
			log.Println(err.Kind())
			log.Println(err.Type())
			log.Println(err.Value())
			log.Println(err.Param())
			log.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}

	// save user to database
}

func validateVariable() {
	myEmail := "joeybloggs.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		log.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}
