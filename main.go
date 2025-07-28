package main

import (
	"fmt"
	"time"

	"github.com/quaminigin/go_test/class_01"
	"github.com/quaminigin/go_test/class_02"
)

func main() {
	// 第2课习题测试
	Class_02_Test()

	// 第1课习题测试
	// Class_01_Test()

}

/*
第2课习题测试
*/
func Class_02_Test() {
	// 1. 指针 - 加10
	// var num int = 100
	// class_02.Add10(&num)
	// fmt.Println(num)

	// 2. 指针 - 乘以2
	// numlist := []int{1, 3, 5, 2, 8, 10, 0, 1}
	// class_02.TimeSliceBy2(&numlist)
	// fmt.Println(numlist)

	// 3. 协程 - 打印奇偶序列
	// var wg sync.WaitGroup
	// wg.Add(2)
	// class_02.PrintWith2Go(&wg)
	// wg.Wait()
	// fmt.Println("Over")

	// 4. 协程 - 任务调度打印时间
	class_02.TaskScheduling([]func(){
		func() {
			time.Sleep(1 * time.Second)
		},
		func() {
			time.Sleep(4 * time.Second)
		},
		func() {
			time.Sleep(10 * time.Second)
		},
	})
}

/*
第1课习题测试
*/
func Class_01_Test() {
	// 两数之和
	fmt.Println(class_01.SumTwoNumbers())
	// 合并区间
	fmt.Println(class_01.MergeIntervals())
	// 删除有序数组中的重复项
	fmt.Println(class_01.DeleteDumplicate())
	// 加一
	fmt.Println(class_01.Add1())
	// 最长公共前缀
	fmt.Println(class_01.LongestCommonPrefix())
	// 有效的括号
	fmt.Println(class_01.CheckBracket("[(){}][]{}({[()[{}]]})"))
	// 回文数
	fmt.Println(class_01.CheckPalindrome(109088880901))
	// 只出现一次的数字
	fmt.Println(class_01.SearchSingle())
}
