//图，广度优先搜索，找寻2个点之间的最短路径，无向图，2点之间距离为1
//这里适合广度优先，因为两点之间的边长度都是1，广度优先就适合这种权重相同的情况

package main

import "fmt"

var pointCount = 5
var lineCount = 7
var start = 1
var end = 5
var min = -1

type node struct{
	x int
	s int
}
var data = [7][3]int{
	{1,2,1},
	{1,3,1},
	{2,3,1},
	{2,4,1},
	{3,4,1},
	{3,5,1},
	{4,5,1},
}
var book = make(map[int]int)
var path = make([]int, 0)
var chart = map[int]map[int]int{}
var minPath = make([]int, 0)
var queueList = make(map[int]node)

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
		//无向图
		chart[d[0]][d[1]] = d[2]
		chart[d[1]][d[0]] = d[2]
	}
	//队列
	header := 0
	tail := 0
	//起点入队
	queueList[tail] = node{x:start, s:0}
	//标记起点
	book[start] = 1
	//加入轨迹
	path = append(path, start)
	tail++
	//是否到达
	isEnd := false
	//开始循环队列
	for header < tail{
		nowPoint := queueList[header]
		for nextPoint:=1; nextPoint<=pointCount;nextPoint++{
			if book[nextPoint] == 0 && chart[nowPoint.x][nextPoint] > 0{
				queueList[tail] = node{x: nextPoint, s: nowPoint.s + chart[nowPoint.x][nextPoint]}
				tail++
				book[nextPoint] = 1
				path = append(path, nextPoint)
			}
			if nowPoint.x == end{
				isEnd = true
				break
			}
		}
		if isEnd{
			break
		}
		header++
	}
	fmt.Println(fmt.Sprintf("最短路程：%d", queueList[tail-1].s))
}