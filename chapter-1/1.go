//桶排序
//这是个简单版的桶排序，一个数组，k是真实的带排序数字，v是这个数字有多少个。
//例如考试满分100，有2个90，1个100，1个60，1个80，那么这个数组就是a[90]=2 a[100]=1等。

package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Print("请输入同学个数：")
	var count int
	fmt.Scanf("%d", &count)
	if count <= 0{
		return
	}
	//初始化桶
	sourceList := make([]int, 101)
	var source int
	for i:=0; i<count; i++{
		fmt.Print("请输入第" + strconv.Itoa(i) +  "位同学的正整数分数：")
		fmt.Scanf("%d", &source)
		sourceList[source]++
	}
	//从大到小输出
	fmt.Print("分数从小到大排序为：")
	for source, num := range sourceList{
		//分数是0的就不打印了
		if source == 0{
			continue
		}
		//该分数段有几个就打印几次
		for i:=0; i<num; i++{
			fmt.Printf("%d, ", source)
		}
	}
}