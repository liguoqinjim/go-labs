package array_utils

import "fmt"

//求最大子序列和 （就是说子序列加起来和最大）
func FindMaxSeqSum(array []int) int {
	SeqSum := make([]int, 0) // 存储子序列和
	// 初始子序列和为 数组下标为0的值
	SeqSum = append(SeqSum, array[0])
	for i := 1; i < len(array); i++ {
		if array[i] > SeqSum[i-1]+array[i] {
			SeqSum = append(SeqSum, array[i])
		} else {
			SeqSum = append(SeqSum, SeqSum[i-1]+array[i])
		}
	}
	max := SeqSum[0]
	for j := 1; j < len(SeqSum); j++ {
		if SeqSum[j] > SeqSum[j-1] {
			max = SeqSum[j]
		}
	}
	fmt.Println(max) //打印结果
	return max
}
