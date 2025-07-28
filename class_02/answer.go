package class_02

import (
	"fmt"
	"sync"
	"time"
)

/*
指针
1. 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/
func Add10(p *int) {
	*p += 10
}

/*
指针
2. 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/
func TimeSliceBy2(s *[]int) {
	for index := range *s {
		(*s)[index] *= 2
	}
}

/*
Goroutine
1. 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/
func PrintWith2Go(wg *sync.WaitGroup) {
	go func() {
		for i := 1; i < 10; i += 2 {
			fmt.Println(i)
		}
		defer wg.Done()
	}()

	go func() {
		for i := 2; i < 11; i += 2 {
			fmt.Println(i)
		}
		defer wg.Done()
	}()

}

/*
Goroutine
2. 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func TaskScheduling(tasks []func()) {
	timeStatistics := make(map[int]int64)
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for index, task := range tasks {
		go func() {
			startTime := time.Now().UnixNano()
			task()
			interval := time.Now().UnixNano() - startTime
			timeStatistics[index] = interval
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(timeStatistics)
}
