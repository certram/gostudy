package main

import (
	"fmt"

	"github.com/certram/gostudy/homeworkfirstweek/slice"
)

func main() {
	test1 := []int{1, 4, 2, 6, 3, 7, 998}
	newSlice, err1 := slice.DeleteElementAtV1(test1, 0)
	if err1 != nil {
		fmt.Printf("slice删除元素发生错误")
	} else {
		fmt.Println(newSlice)
	}

	test2 := []int{4, 62, 7, 78, 31, 5, 47, 35, 998}
	res2, err2 := slice.DeleteElementAtV1(test2, 7)
	if err2 != nil {
		fmt.Println(err2.Error())
	} else {
		fmt.Println(res2)
	}
	test3 := []int{4, 62, 7, 78, 31, 5, 47, 35}
	res3, err3 := slice.DeleteElementAtV3[int](test3, 6)
	if err3 != nil {
		fmt.Printf("泛型方法调用错误，%s", err3.Error())
	} else {
		fmt.Println(res3)
	}

}
