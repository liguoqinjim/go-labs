package main

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"
	"log"
)

func main() {
	list := arraylist.New()
	list.Add("a")
	list.Add("c", "b")

	//a,c,b
	log.Println(list)

	//排序
	list.Sort(utils.StringComparator)
	//a,b,c
	log.Println("sort:", list)

	//查找
	v1, ok := list.Get(0)
	log.Println("Get(0):", v1, ok)

	v1, ok = list.Get(100)
	log.Println("Get(100):", v1, ok)

	//包含
	isContained := list.Contains("a")
	log.Println("Contains a:", isContained)
	isContained = list.Contains("a", "b")
	log.Println("Contains a,b:", isContained)
	isContained = list.Contains("a", "b", "d")
	log.Println("Contains a,b,d:", isContained)

	//交换位置
	list.Swap(0, 1)
	log.Println("swap:", list)

	//Remove
	list.Remove(0)
	log.Println("remove:", list)

	//Empty
	isEmpty := list.Empty()
	log.Println("empty:", isEmpty)

	//Size
	size := list.Size()
	log.Println("size:", size)

	//插入insert
	list.Insert(0, "z")
	log.Println("insert:", list)

	//clear
	list.Clear()
	log.Println("clear:", list)
}
