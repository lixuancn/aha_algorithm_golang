//最短路径 - Bellman-ford
package main

import "fmt"

var cityCount = 5
var lineCount = 5

var userInputGameMap = [5][3]int{
	{2,3,2},
	{1,2,-3},
	{1,5,5},
	{4,5,2},
	{3,4,3},
}

var inf = 2<<31 - 1

var u, v, w [5]int

func main(){
	for i:=0; i<lineCount; i++{
		u[i] = userInputGameMap[i][0]
		v[i] = userInputGameMap[i][1]
		w[i] = userInputGameMap[i][2]
	}
	var dis [5]int
	for i:=0; i<cityCount; i++{
		dis[i] = inf
	}
	dis[0] = 0
	for k:=0; k<=cityCount-1; k++{
		for i:=0; i<lineCount; i++{
			if dis[v[i]-1] > dis[u[i]-1] + w[i]{
				dis[v[i]-1] = dis[u[i]-1] + w[i]
			}
		}
	}
	fmt.Println(dis)
}