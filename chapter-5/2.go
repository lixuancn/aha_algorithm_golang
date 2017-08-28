//图，深度优先搜索，找寻2个点之间的最短路径，有向图

package main

import "fmt"

var pointCount = 5
var lineCount = 8
var start = 1
var end = 5
var min = -1

var data = [8][3]int{
	{1,2,2},
	{1,5,10},
	{2,3,3},
	{2,5,7},
	{3,1,4},
	{3,4,4},
	{4,5,5},
	{5,3,3},
}
var book = make(map[int]int)
var path = make([]int, 0)
var chart = map[int]map[int]int{}
var minPath = make([]int, 0)

func main(){
	//初始化图
	for i:=1; i<=pointCount; i++{
		chart[i] = map[int]int{}
		for j:=1; j<=pointCount; j++{
			if i == j{
				chart[i][j] = 0
			}else{
				chart[i][j] = -1
			}
		}
	}
	for _, d := range data{
		chart[d[0]][d[1]] = d[2]
	}
	//标记起点
	book[start] = 1
	//加入轨迹
	path = append(path, start)
	dfs(start, 0)
	fmt.Println(fmt.Sprintf("最短路程：%d", min))
	fmt.Println(minPath)
}

func dfs(point, nowDis int){
	if min != -1 && nowDis > min{
		return
	}
	//是否到达目标
	if point == end{
		if min == -1 || nowDis < min{
			min = nowDis
			minPath = path
		}
		return
	}
	for nextPoint:=1; nextPoint<=pointCount;nextPoint++{
		if book[nextPoint] == 0 && chart[point][nextPoint] > 0{
			book[nextPoint] = 1
			path = append(path, nextPoint)
			dfs(nextPoint, nowDis + chart[point][nextPoint])
			//取消标记
			book[nextPoint] = 0
			//路径删除
			path = path[ : len(path)-1]
		}
	}
	return
}