//广度搜索-再解炸弹人
package main

import "fmt"

//地图是13*13，G是敌人，#是墙，.是空地
//x,y对应坐标轴
var maxX = 13
var maxY = 13
var gameMap = [13]string{
	"#############",
	"#GG.GGG#GGG.#",
	"###.#G#G#G#G#",
	"#.......#..G#",
	"#G#.###.#G#G#",
	"#GG.GGG.#.GG#",
	"#G#.#G#.#.#.#",
	"##G...G.....#",
	"#G#.#G###.#G#",
	"#...G#GGG.GG#",
	"#G#.#G#G#.#G#",
	"#GG.GGG#G.GG#",
	"#############",
}

//链表结构体
type queue struct{
	x int
	y int
}

var startX = 4
var startY = 4

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
	}
	//队列后移
	tail++
	//标记走过的路
	b := make(map[int]int)
	b[startY] = 1
	book[startX] = b
	//获取可以炸到多少个敌人
	bestNum := getNum(startX-1, startY-1)
	fmt.Println(bestNum)
	return
	bestX := startX
	bestY := startY
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
			count := getNum(nextX-1, nextY-1)
			if bestNum < count{
				bestNum = count
				bestX = nextX
				bestY = nextY
			}
		}
		header++
		fmt.Println("-------一个点扩展结束--------")
	}
	fmt.Println(fmt.Sprintf("在%d, %d可炸%d个敌人", bestX, bestY, bestNum))
}

func log(msg string, x, y int){
	fmt.Println(fmt.Sprintf("%s: %d, %d", msg, x, y))
}

func getNum(x, y int)int{
	//消灭敌人的总数
	sum := 0
	//计算左边可以消灭的敌人数量
	m := x
	n := y
	fmt.Println(string(gameMap[m][n]))
	for gameMap[m][n] != '#'{
		if gameMap[m][n] == 'G'{
			sum++
		}
		m--
	}
	//计算右边可以消灭的敌人数量
	m = x
	n = y
	for gameMap[m][n] != '#'{
		if gameMap[m][n] == 'G'{
			sum++
		}
		m++
	}
	//计算下边可以消灭的敌人数量
	m = x
	n = y
	for gameMap[m][n] != '#'{
		if gameMap[m][n] == 'G'{
			sum++
		}
		n--
	}
	//计算上边可以消灭的敌人数量
	m = x
	n = y
	for gameMap[m][n] != '#'{
		if gameMap[m][n] == 'G'{
			sum++
		}
		n++
	}
	return sum
}