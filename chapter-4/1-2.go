//深度优先搜索 - 火柴棍拼数字
package main

import "fmt"

//火柴总数
var total = 48

var boxCount = 9
var a [10]int

var numCount = 10
var book [10]int


func main(){
	dfs(1)
}

func dfs(step int){
	if step > boxCount{
		intAndNumMap := [10]int{6, 2, 5,5,4,5,6,3,7,6}
		if a[1] == 0 || a[4] == 0 || a[5] ==0{
			return
		}
		if 100*a[1] + 10*a[2] + a[3] + 100*a[4] + 10*a[5] + a[6] == 100*a[7] + 10*a[8] + a[9]{
			sum := 0
			for k, i := range a{
				if k == 0{
					continue
				}
				sum += intAndNumMap[i]
			}
			//等号和加好个需要2个火柴
			if sum == total - 4{
				fmt.Println(fmt.Sprintf("%d+%d=%d", 100*a[1] + 10*a[2] + a[3], 100*a[4] + 10*a[5] + a[6], 100*a[7] + 10*a[8] + a[9]))
			}
		}
		return
	}
	for i:=0; i<numCount; i++{
		if book[i] == 1{
			continue
		}
		a[step] = i
		book[i] = 1
		newStep := step + 1
		dfs(newStep)
		book[i] = 0
	}
	return
}