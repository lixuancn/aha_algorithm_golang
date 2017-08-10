//冒泡
//逻辑就不多说了，案例还是考试分数的排序，满分100，分数只能是正整数
//下面代码的时间复杂度是O(N*(N+(N-1)+(N-2)+(N-3)...))=O(N*N)=O(N²)

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
	sourceList := make([]int, count)
	var source int
	for i:=0; i<count; i++{
		fmt.Print("请输入第" + strconv.Itoa(i) +  "位同学的正整数分数：")
		fmt.Scanf("%d", &source)
		sourceList[i] = source
	}
	fmt.Println(len(sourceList))
	//冒泡
	for i:=0; i<count-1; i++{
		for j:=0; j<count-1-i;j++{
			if sourceList[j] > sourceList[j+1]{
				tmp := sourceList[j+1]
				sourceList[j+1] = sourceList[j]
				sourceList[j] = tmp
			}
		}
	}
	fmt.Println(len(sourceList))
	//从大到小输出
	fmt.Print("分数从小到大排序为：")
	for _, source := range sourceList{
		fmt.Print(source)
		fmt.Print("。")
	}
}