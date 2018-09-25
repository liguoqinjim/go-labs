// Package lab001
//
// This is a test of godoc
package lab001

import "log"

// 这是一个常数
const TEST_MODE = true

// this is a struct
// Student
type Student struct {
	Sid   int
	Sname string
}

// this is student's function,
// function for do the homework
func (s *Student) Homework() {
	log.Println("student do the homework")
}

// B

// 1223Hello
//B
// A function say hello
//C
func Hello() {

}

// BUG(tom): 我是bug说明

// Hi, this is a hi function
func Hi() {

}

// BUG(kimi): 我是bug注释2
// Bonjour ...也是pkg注释
func Bonjour() {

}
