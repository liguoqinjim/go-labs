package main

import "fmt"

func main() {
	fmt.Printf("9.8249   =>   %0.2f(四舍)\n", 9.8249)
	fmt.Printf("9.82671  =>   %0.2f(六入)\n", 9.82671)
	fmt.Printf("9.8351   =>   %0.2f(五后非零就进一)\n", 9.8351)
	fmt.Printf("9.82501  =>   %0.2f(五后非零就进一)\n", 9.82501)
	fmt.Printf("9.8250   =>   %0.2f(五后为零看奇偶，五前为偶应舍去)\n", 9.8250)
	fmt.Printf("9.8350   =>   %0.2f(五后为零看奇偶，五前为奇要进一)\n", 9.8350)
	fmt.Printf("9.8350   =>   %0.0f(六入)\n", 9.8350)

	//%0.2f的这个0表示width 表示输出的长度
	fmt.Printf("\n\n")
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%0.2f|%0.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) //'-'表示靠左
}
