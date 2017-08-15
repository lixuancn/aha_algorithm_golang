//炸弹人
//地图是13*13，G是敌人，#是墙，.是空地
//x,y对应坐标轴

package main

import (
	"fmt"
)

func main(){
	x := 13
	y := 13
	gameMap := [13]string{
		"#############",
		"#GG.GGG#GGG.#",
		"###.#G#G#G#G#",
		"#.......#..G#",
		"#G#.###.#G#G#",
		"#GG.GGG.#.GG#",
		"#G#.#G#.#.###",
		"##G...G.....#",
		"#G#.#G###.#G#",
		"#...G#GGG.GG#",
		"#G#.#G#G#.#G#",
		"#GG.GGG#G.GG#",
		"#############",
	}
	bestX := 0
	bestY := 0
	bestNum := 0
	var m, n int
	for i:=0; i<x; i++{
		for j:=0; j<y; j++{
			//在空地上放炸弹
			if gameMap[i][j] != '.'{
				continue
			}
			//消灭敌人的总数
			sum := 0
			//计算左边可以消灭的敌人数量
			m = i
			n = j
			for gameMap[m][n] != '#'{
				if gameMap[m][n] == 'G'{
					sum++
				}
				m--
			}
			//计算右边可以消灭的敌人数量
			m = i
			n = j
			for gameMap[m][n] != '#'{
				if gameMap[m][n] == 'G'{
					sum++
				}
				m++
			}
			//计算下边可以消灭的敌人数量
			m = i
			n = j
			for gameMap[m][n] != '#'{
				if gameMap[m][n] == 'G'{
					sum++
				}
				n--
			}
			//计算上边可以消灭的敌人数量
			m = i
			n = j
			for gameMap[m][n] != '#'{
				if gameMap[m][n] == 'G'{
					sum++
				}
				n++
			}
			fmt.Println(sum)
			if sum > bestNum{
				bestX = i
				bestY = j
				bestNum = sum
			}
		}
	}
	msg := fmt.Sprintf("放置在%d, %d可以消灭的敌人最多， 是%d", bestX, bestY, bestNum)
	fmt.Println(msg)
}