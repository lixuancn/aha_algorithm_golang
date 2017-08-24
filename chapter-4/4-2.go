//深度搜索-再解炸弹人
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

//已经走过的路
var book = make(map[int]map[int]int)

var bestX = 0
var bestY = 0
var bestSum = 0

func main(){
	//入口
	bestX = startX
	bestY = startY
	bestSum = getNum(startX-1, startY-1	)
	book[startX] = map[int]int{startY: 1}
	dfs(startX, startY)
	fmt.Println(fmt.Sprintf("在%d, %d可炸%d个敌人", bestX, bestY, bestSum))
}

func dfs(x, y int){
	//下一步的方向
	nextList := [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0},}
	for _, next := range nextList{
		//下一个点
		nextX := x + next[0]
		nextY := y + next[1]
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
		//记录已走过
		if b, ok := book[nextX]; ok{
			b[nextY] = 1
			book[nextX] = b
		}else{
			book[nextX] = map[int]int{nextY: 1}
		}
		//计算
		count := getNum(nextX-1, nextY-1)
		if bestSum < count{
			bestSum = count
			bestX = nextX
			bestY = nextY
		}
		dfs(nextX, nextY)
		book[nextX][nextY] = 0
	}
}

func getNum(x, y int)int{
	//消灭敌人的总数
	sum := 0
	//计算左边可以消灭的敌人数量
	m := x
	n := y
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

func log(msg string, x, y int){
	fmt.Println(fmt.Sprintf("%s: %d, %d", msg, x, y))
}