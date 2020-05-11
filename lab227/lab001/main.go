package main

import (
	"github.com/iancoleman/strcase"
	"log"
)

var (
	//snakeCase = "any_kind_of_string"
	//CamelCase = "AnyKindOfString"

	s = "AnyKind of_string"
)

func main() {
	//snakeCase
	log.Println(strcase.ToSnake(s))

	//camelcase
	log.Println(strcase.ToCamel(s))
	log.Println(strcase.ToLowerCamel(s))
}

//例子
//ToSnake(s)	any_kind_of_string
//ToSnakeWithIgnore(s, '.')	any_kind.of_string
//ToScreamingSnake(s)	ANY_KIND_OF_STRING
//ToKebab(s)	any-kind-of-string
//ToScreamingKebab(s)	ANY-KIND-OF-STRING
//ToDelimited(s, '.')	any.kind.of.string
//ToScreamingDelimited(s, '.', '', true)	ANY.KIND.OF.STRING
//ToScreamingDelimited(s, '.', ' ', true)	ANY.KIND OF.STRING
//ToCamel(s)	AnyKindOfString
//ToLowerCamel(s)	anyKindOfString
