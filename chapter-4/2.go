//深度搜索-最短路径找人
package main

import "fmt"

//起点坐标
var startX = 1
var startY = 1
//终点坐标
var endX = 4
var endY = 3
//已经走过的路
var b = make(map[int]int)
var book = make(map[int]map[int]int)

//地图，#是障碍物，.是空地
var maxX = 5
var maxY = 4
var gameMap = []string{
	"..#.",
	"....",
	"..#.",
	".#..",
	"...#",
}

var min = 999999

var nextList = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func main(){
	b[startY] = 1
	book[startX] = b
	dfs(startX, startY, 0)
	fmt.Println(fmt.Sprintf("最短路径是%d", min))
}

func dfs(x, y, step int){
	if y ==3 {
		fmt.Println(x, "---", y)
	}
	if x == endX && y == endY{
		if step < min{
			min = step
		}
		return
	}
	for _, next := range nextList{
		nowX := x + next[0]
		nowY := y + next[1]
		//超出边界计算
		if nowX < 1 || nowY < 1 || nowX > maxX || nowY > maxY{
			continue
		}
		//已经在路径中了
		if book[nowX][nowY] != 0{
			continue
		}
		//障碍物
		//x,y是从1开始的，gameMap是从0开始的
		if gameMap[nowX-1][nowY-1] != '.'{
			continue
		}
		b[nowY] = 1
		book[nowX] = b
		newStep := step + 1
		dfs(nowX, nowY, newStep)
		book[nowX][nowY] = 0
	}
	return
}