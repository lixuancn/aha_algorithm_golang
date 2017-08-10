//队列
//一组数组：6 3 1 7 5 8 9 2 4，删除第一个，第二个放队位，直至剩余最后一个数字，求被删除的数字列表
//为体现链表，这组数用数组而不是切片

package main

import (
	"fmt"
)

func main(){
	delList := make([]int, 0)
	numList := [100]int{6,3,1,7,5,8,9,2,4}
	//当前队列中第一位
	header := 0
	//结尾的位置+1
	tail := 9
	for header < tail{
		//删除第一位
		delList = append(delList, numList[header])
		header++
		//第二位放到最后
		numList[tail] = numList[header]
		header++
		tail++
	}
	fmt.Print("被删除的列表为：")
	fmt.Println(delList)
}