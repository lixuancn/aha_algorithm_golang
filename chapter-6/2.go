//最短路径 - Dijkstra - 单源最短路径 - 一个点到其他个点的最短距离
package main

import "fmt"

var cityCount = 6

var userInputGameMap = [9][3]int{
	{1,2,1},
	{1,3,12},
	{2,3,9},
	{2,4,3},
	{3,5,5},
	{4,3,4},
	{4,5,13},
	{4,6,15},
	{5,6,4},
}

var inf = 2<<31 - 1

var min int

func main(){
	var gameMap [6][6]int
	for i:=0; i<cityCount; i++{
		for j:=0; j<cityCount; j++{
			if i == j{
				gameMap[i][j] = 0
			}else{
				gameMap[i][j] = min
			}
		}
	}
	for _, v := range userInputGameMap{
		gameMap[v[0]-1][v[1]-1] = v[2]
	}
	var dis [6]int
	for k, v := range gameMap[0]{
		dis[k] = v
	}
	var book [6]int
	book[0] = 1
	//算法的核心
	var u int
	for i:=0; i<cityCount-1; i++{
		//找到离1号定点最近的点
		min = inf
		for j:=0; j<cityCount; j++{
			if book[j] == 0 && dis[j] < min{
				min = dis[j]
				u = j
			}
		}
		book[u] = 1
		for v:=0; v<cityCount; v++{
			if gameMap[u][v] != inf && gameMap[u][v] != 0{
				if dis[v] > dis[u] + gameMap[u][v]{
					dis[v] = dis[u] + gameMap[u][v]
				}
			}
		}
	}
	fmt.Println(dis)
}