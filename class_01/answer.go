package class_01

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strconv"
)

/*
### 基础
- **两数之和 **

考察：数组遍历、map使用

题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数

链接：https://leetcode-cn.com/problems/two-sum/
*/
func SumTwoNumbers() (int, int, error) {
	var numlist = [...]int{2, 7, 11, 15, 44, 23, 21, 8, 46}
	var sum = 91

	for i := 0; i < len(numlist); i++ {
		remain := sum - numlist[i]
		for j := i + 1; j < len(numlist); j++ {
			if numlist[j] == remain {
				return numlist[i], remain, nil
			}
		}
	}
	return 0, 0, errors.New("not found")
}

/*
合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
*/
func MergeIntervals() [][2]int {
	var intervals = [...][2]int{
		{8, 10},
		{1, 3},
		{2, 6},
		{17, 32},
		{15, 28},
		{16, 20},
		{0, 2},
	}
	if len(intervals) == 0 {
		panic("初始数组长度为0")
	}
	sort.Slice(intervals[:], func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
	var merged = make([][2]int, 0, len(intervals))
	merged = append(merged, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lastIndex := len(merged) - 1
		if intervals[i][0] <= merged[lastIndex][1] {
			merged[lastIndex][1] = max(merged[lastIndex][1], intervals[i][1])
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}

/*
### 引用类型：切片
- **[26. 删除有序数组中的重复项](https://leetcode.cn/problems/remove-duplicates-from-sorted-array/)**：
// 给你一个有序数组 `nums` ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
// 可以使用双指针法，一个慢指针 `i` 用于记录不重复元素的位置，一个快指针 `j` 用于遍历数组，
// 当 `nums[i]` 与 `nums[j]` 不相等时，将 `nums[j]` 赋值给 `nums[i + 1]`，并将 `i` 后移一位。
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
*/
func DeleteDumplicate() int {
	var numlist = [...]int{1, 1, 2, 3, 3, 3, 4, 5, 6, 6, 6, 6, 7, 8, 8, 9, 10, 10, 10, 10}
	i := 0
	for j := 1; i < len(numlist) && j < len(numlist); {
		if numlist[i] == numlist[j] {
			j++
		} else {
			i++
			numlist[i] = numlist[j]
			j++
		}
	}
	fmt.Println(i+1, numlist)
	return i + 1
}

/*
加一

难度：简单

考察：数组操作、进位处理

题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一

链接：https://leetcode-cn.com/problems/plus-one/
*/
func Add1() []int {
	numArray := [...]int{9, 9, 9}
	fmt.Println(numArray)
	result := [len(numArray) + 1]int{0}

	for i := 0; i < len(numArray); i++ {
		result[i+1] = numArray[i]
	}

	needAdd1 := true
	for i := len(result) - 1; i >= 0; i-- {
		if needAdd1 {
			if result[i] == 9 {
				result[i] = 0
				needAdd1 = true
			} else {
				result[i]++
				needAdd1 = false
				break
			}
		} else {
			break
		}
	}
	if result[0] == 0 {
		return result[1:]
	} else {
		return result[:]
	}
}

/*
最长公共前缀

考察：字符串处理、循环嵌套

题目：查找字符串数组中的最长公共前缀

链接：https://leetcode-cn.com/problems/longest-common-prefix/
*/
func LongestCommonPrefix() string {
	var strArray = [...]string{"12e123123", "12e", "12easdas", "12e43f", "12easdasdww", "12e1"}
	control := strArray[0]
	switch len(strArray) {
	case 0:
		panic("数组初始长度为0")
	case 1:
		return control
	}
	leastCommonIndex := math.MaxInt
	for i := 1; i < len(strArray); i++ {
		current := strArray[i]
		currentCommonIndex := 0
		for strIndex := 0; strIndex < len(control) && strIndex < len(current); strIndex++ {
			if control[strIndex] == current[strIndex] {
				currentCommonIndex++
			}
		}
		if currentCommonIndex < leastCommonIndex {
			leastCommonIndex = currentCommonIndex
		}
	}
	return control[0:leastCommonIndex]
}

/*
有效的括号

考察：字符串处理、栈的使用

题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效

链接：https://leetcode-cn.com/problems/valid-parentheses/
*/
func CheckBracket(str string) bool {
	charStack := make([]rune, 0, len(str))
	for _, char := range str {
		switch char {
		case '[', '{', '(':
			charStack = append(charStack, char)
		case ']':
			index := len(charStack) - 1
			if index >= 0 && charStack[index] == '[' {
				charStack = charStack[:index]
			} else {
				return false
			}
		case '}':
			index := len(charStack) - 1
			if index >= 0 && charStack[index] == '{' {
				charStack = charStack[:index]
			} else {
				return false
			}
		case ')':
			index := len(charStack) - 1
			if index >= 0 && charStack[index] == '(' {
				charStack = charStack[:index]
			} else {
				return false
			}
		default:
			panic("illegal string")
		}
	}
	if len(charStack) == 0 {
		return true
	} else {
		return false
	}
}

/*
回文数

考察：数字操作、条件判断
题目：判断一个整数是否是回文数
*/
func CheckPalindrome(num int) bool {
	if num < 0 {
		return false
	} else {
		if num > 9 {
			str := strconv.Itoa(num)
			for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
				if str[i] != str[j] {
					return false
				}
			}
			return true
		} else {
			return true
		}
	}
}

/*
只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func SearchSingle() int {
	numlist := [...]int{-5, 9, 8, 3, 0, 31, -99, 31, 0, 9, 3, -5, -99}
	control := numlist[0]
	for i := 1; i < len(numlist); i++ {
		control = control ^ numlist[i]
	}
	return control
}
