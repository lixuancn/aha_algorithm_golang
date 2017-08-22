//广度搜索-最短路径找人
//深度优先是先把一种可能走到头，然后再回头一步步推， 广度优先是先把这个点的所有可能都尝试一遍，再走下一步
package main

import "fmt"

//起点坐标
var startX = 1
var startY = 1
//终点坐标
var endX = 4
var endY = 3

//链表结构体
type queue struct{
	x int
	y int
	parent int
	step int
}

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

func main(){
	//已经走过的路
	var book = make(map[int]map[int]int)
	//下一步的方向
	nextList := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0},}
	//入口入队
	header := 1
	tail := 1
	queueList := make(map[int]queue)
	queueList[tail] = queue{
		x: startX,
		y: startY,
		parent: 0,
		step: 0,
	}
	//队列后移
	tail++
	//标记走过的路
	b := make(map[int]int)
	b[startY] = 1
	book[startX] = b
	//标记是否到达终点
	flag := 0
	//循环非空队列
	for header < tail{
		for _, next := range nextList{
			//下一个点
			nextX := queueList[header].x + next[0]
			nextY := queueList[header].y + next[1]
			log("当前位置", nextX, nextY)
			//是否越界
			if nextX <= 0 || nextY <= 0 || nextX > maxX || nextY > maxY{
				log("超出边界", nextX, nextY)
				continue
			}
			//已经走过
			if _, ok := book[nextX][nextY]; ok{
				log("已经走过", nextX, nextY)
				continue
			}
			//障碍物
			if gameMap[nextX-1][nextY-1] != '.'{
				log("障碍物", nextX, nextY)
				continue
			}
			//入队
			queueList[tail] = queue{
				x: nextX,
				y: nextY,
				parent: header,
				step: queueList[header].step + 1,
			}
			tail++
			//标记已经走过
			if b, ok := book[nextX]; ok{
				b[nextY] = 1
				book[nextX] = b
			}else{
				b = make(map[int]int)
				b[nextY] = 1
				book[nextX] = b
			}
			//到了地方
			if nextX == endX && nextY == endY{
				flag = 1
				break
			}
		}
		if flag == 1{
			break
		}
		header++
		fmt.Println("-------一个点扩展结束--------")
	}
	fmt.Println(fmt.Sprintf("最短路径是%d", queueList[tail-1].step))
}

func log(msg string, x, y int){
	fmt.Println(fmt.Sprintf("%s: %d, %d", msg, x, y))
}