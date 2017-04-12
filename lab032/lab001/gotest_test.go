package lab001

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除法函数测试没通过")
	} else {
		t.Log("第一个测试通过了")
	}
}
func Test_Division_2(t *testing.T) {
	t.Error("就是不通过")
}
