package m2

import (
	"fmt"

	"github.com/henrylee2cn/cfgo"
)

type T2 struct {
	X string
	Y []string `yaml:",flow"`
	Z []int
	N bool
}

func (t *T2) Reload(bind cfgo.BindFunc) error {
	fmt.Println("module_2: T2 reload do some thing...")
	return bind()
}

func init() {
	structPtr2 := &T2{
		X: "xxx",                   //default value
		Y: []string{"x", "y", "z"}, //default value
		Z: []int{1, 2, 3},          //default value
	}
	{
		c := cfgo.MustGet("config/config.yaml")
		c.MustReg("section2", structPtr2)
	}
	// or
	// cfgo.MustReg("section2", structPtr2)
	fmt.Printf("structPtr2(config/config.yaml): %+v\n\n", structPtr2)
}
