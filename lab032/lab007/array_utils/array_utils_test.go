package array_utils

import (
	"fmt"
	"os"
	"testing"
)

//TestMain会在下面所有测试方法执行开始前先执行，一般用于初始化资源和执行完后释放资源
func TestMain(m *testing.M) {
	fmt.Println("初始化资源")
	result := m.Run() //运行go的测试，相当于调用main方法
	fmt.Println("释放资源")
	os.Exit(result) //退出程序
}

//单元测试
func TestFindMaxSeqSum(t *testing.T) {
	sum := FindMaxSeqSum([]int{1, 3, -9, 6, 8, -19})
	if sum == 14 {
		t.Log("successful")
	} else {
		t.Error("failed")
	}
}

//基准测试
func BenchmarkFindMaxSeqSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindMaxSeqSum([]int{1, 3, -9, 6, 8, -19})
	}
}

//这个验证的是FindMaxSeqSum方法控制台输出的max和OutPut后面的14是否一致，如果相同，则表示验证通过，否则测试用例失败
func ExampleFindMaxSeqSum() {
	FindMaxSeqSum([]int{1, 3, -9, 6, 8, -19})
	// OutPut: 14
}
