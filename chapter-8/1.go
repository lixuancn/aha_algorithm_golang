//生成最小树，N个点，只需要N-1条边
//Kruskal算法

package main

import "fmt"

var cityCount = 6
var lineCount = 9
var userInputMap = [9][3]int{
	{2,4,11},
	{3,5,13},
	{4,6,3},
	{5,6,4},
	{2,3,6},
	{4,5,7},
	{1,2,1},
	{3,4,9},
	{1,3,2},
}

type edge struct{
	u int
	v int
	w int
}

var edgeList [10]edge
var f [7]int
var sum, count int

//从小到大
func quickSort(left, right int){
	if left > right{
		return
	}
	i := left
	j := right
	for i != j{
		//先从右面开始
		for edgeList[j].w >= edgeList[i].w && i < j{
			j--
		}
		for edgeList[i].w <= edgeList[j].w && i < j{
			i++
		}
		if i < j{
			t := edgeList[i]
			edgeList[i] = edgeList[j]
			edgeList[j] = t
		}
	}
	quickSort(left, i-1)
	quickSort(i+1, right)
}

func getParent(v int)int{
	if f[v] == v{
		return v
	}
	f[v] = getParent(f[v])
	return f[v]
}

//合并两个子集
func merge(u, v int)bool{
	t1 := getParent(v)
	t2 := getParent(u)
	if t1 != t2{
		f[t2] = t1
		return true
	}
	return false
}

func main(){
	for k, line := range userInputMap{
		edgeList[k+1].u = line[0]
		edgeList[k+1].v = line[1]
		edgeList[k+1].w = line[2]
	}
	//从小对大对边的权值排序
	quickSort(1, lineCount)
	//初始化查并集
	for k, _ := range f{
		f[k] = k
	}
	//Kruskal算法核心
	for i:=1; i<=lineCount; i++{
		if merge(edgeList[i].u, edgeList[i].v){
			count++
			sum = sum + edgeList[i].w
		}
		if i == cityCount -1{
			break
		}
	}
	fmt.Println(edgeList)
	fmt.Println(sum)
}