//快速排序
//逻辑就不多说了，案例还是考试分数的排序，满分100，分数只能是正整数
//下面代码的时间复杂度是O(Nlogn) ~ O(N²)

package main

import (
	"fmt"
	"strconv"
)

var sourceList []int

func main(){
	fmt.Print("请输入同学个数：")
	var count int
	fmt.Scanf("%d", &count)
	if count <= 0{
		return
	}
	var source int
	for i:=0; i<count; i++{
		fmt.Print("请输入第" + strconv.Itoa(i) +  "位同学的正整数分数：")
		fmt.Scanf("%d", &source)
		sourceList = append(sourceList, source)
	}
	//快排
	quickSort(0, count-1)
	//从大到小输出
	fmt.Print("分数从小到大排序为：")
	for _, source := range sourceList{
		fmt.Print(source)
		fmt.Print("。")
	}
}

func quickSort(left, right int){
	if left >= right{
		return
	}
	i := left
	j := right
	//基准书
	standard := sourceList[left]
	for i != j{
		//右侧开始向左移动
		for sourceList[j] >= standard && i < j{
			j--
		}
		//左侧开始向右移动
		for sourceList[i] <= standard && i < j{
			i++
		}
		if i < j{
			tmp := sourceList[j]
			sourceList[j] = sourceList[i]
			sourceList[i] = tmp
		}
	}
	sourceList[left] = sourceList[i]
	sourceList[i] = standard
	//递归，继续排左面
	quickSort(left, i-1)
	//递归，继续排右面
	quickSort(i+1, right)
}