package rpn

import (
	"fmt"
	"testing"
)

func TestArithmetic(t *testing.T) {
	fmt.Println("TestArithmetic: ")
	ins := []string{
		"1+(2-3)",
		"1+(2-3)*4",
		"1+4*(2-3)",
		"(0)-1+(1.5+1)+2*3/4+(((1+2)*4))",
	}
	for _, in := range ins {
		ast, err := Parse(in)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Println(in, "\n\t", ast.Rpn())
	}
}
