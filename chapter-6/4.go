//最短路径 - Bellman-ford - 队列优化法
package main

import "fmt"

var cityCount = 5
var lineCount = 7

var userInputGameMap = [7][3]int{
	{1,2,2},
	{1,5,10},
	{2,3,3},
	{2,5,7},
	{3,4,4},
	{4,5,5},
	{5,3,6},
}

var inf = 2<<31 - 1

var u, v, w [8]int

func main(){
	var book [6]int
	var first [6]int
	for i:=1; i<=cityCount; i++{
		first[i] = -1
	}
	var next [8]int
	for i:=1; i<=lineCount; i++{
		u[i] = userInputGameMap[i-1][0]
		v[i] = userInputGameMap[i-1][1]
		w[i] = userInputGameMap[i-1][2]
		//next[i] = first[u[i]]
		//first[u[i]] = i
	}
	for i:=1; i<=lineCount; i++{
		//u[i] = userInputGameMap[i-1][0]
		//v[i] = userInputGameMap[i-1][1]
		//w[i] = userInputGameMap[i-1][2]
		next[i] = first[u[i]]
		first[u[i]] = i
		fmt.Println(fmt.Sprintf("第%d条边，起始顶点%d, 边为%#v", i, u[i], userInputGameMap[i-1]))
		fmt.Println(first)
		fmt.Println(next)
		fmt.Println("-------")
	}
	fmt.Println(first)
	fmt.Println(next)

	var dis [6]int
	for i:=1; i<=cityCount; i++{
		dis[i] = inf
	}
	dis[1] = 0

	queue := make(map[int]int, 0)
	header := 1
	tail := 1

	//1号点入队
	queue[tail] = 1
	tail++
	book[1] = 1

	for header < tail{
		k := first[queue[header]]
		if k != -1{
			if dis[v[k]] > dis[u[k]] + w[k]{
				dis[v[k]] = dis[u[k]] + w[k]
				if book[v[k]] == 0{
					queue[tail] = v[k]
					tail++
					book[v[k]] = 1
				}
			}
			k = next[k]
		}
		book[queue[header]] = 0
		header++
	}
	fmt.Println(dis)
}