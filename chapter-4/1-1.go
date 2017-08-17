//深度优先搜索 - 纸盒子
//有3个纸盒子，有1、2、3三个数字，那么组合有123、132、213、231、312、321

package main

import (
	"fmt"
)

//总数
var count = 3
//从1开始，所有个数是count+1
var book [4]int
var a [4]int
func main(){
	dfs(1)
}

func dfs(step int){
	if step > count{
		fmt.Println(a[1:])
		return
	}
	for i:=1; i<=count; i++{
		if book[i] == 1{
			continue
		}
		a[step] = i
		book[i] = 1
		nextStep := step + 1
		dfs(nextStep)
		book[i] = 0
	}
	return
}