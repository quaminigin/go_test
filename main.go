package main

import (
	"fmt"

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
	// class_02.TaskScheduling([]func(){
	// 	func() {
	// 		time.Sleep(1 * time.Second)
	// 	},
	// 	func() {
	// 		time.Sleep(4 * time.Second)
	// 	},
	// 	func() {
	// 		time.Sleep(10 * time.Second)
	// 	},
	// })

	// 5. 面向对象 - 圆和长方形
	// var s1 class_02.Shape = &class_02.Circle{Radius: 2.0}
	// var s2 class_02.Shape = &class_02.Rectangle{Width: 3.0, Height: 4.0}
	// fmt.Println(s1.Area())
	// fmt.Println(s1.Perimeter())
	// fmt.Println(s2.Area())
	// fmt.Println(s2.Perimeter())

	// 6. 面向对象 - 结构体嵌套
	// e := class_02.Employee{Person: class_02.Person{Name: "QN", Age: 33}, EmployeeID: "z001"}
	// e.PrintInfo()

	// 7. Channel - 传int值
	// class_02.ChannelOfInt()

	// 8. Channel - 带有缓冲的通道
	// class_02.BufferedChannel()

	// 9. 锁机制 - 计数器
	// class_02.SyncTest()

	// 10. 锁机制 - 原子操作
	class_02.AtomicTest()
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
