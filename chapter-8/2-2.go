//生成最小树 - Prim算法
package main

import (
	"math"
	"fmt"
)

var pointCount = 6
var lineCount = 9

var userInputMap = [10][3]int{
	{2,4,11},
	{3,5,13},
	{4,6,3},
	{5,6,4},
	{2,3,6},
	{4,5,7},
	{1,2,1},
	{3,4,9},
	{1,3,2},
}

func main(){
	var book [7]int
	var dis [7]int
	var gameMap [7][7]int
	inf := math.MaxInt64
	var count, sum int
	for i:=0; i<= pointCount; i++{
		for j:=0; j<= pointCount; j++ {
			if i == j{
				gameMap[i][j] = inf
			}else{
				gameMap[i][j] = inf
			}
		}
	}
	for _, v := range userInputMap{
		gameMap[v[0]][v[1]] = v[2]
		gameMap[v[1]][v[0]] = v[2]
	}
	//1号定点入树
	for i:=1; i<=pointCount; i++{
		dis[i] = gameMap[1][i]
	}
	book[1] = 1
	count++
	j := 0
	for count < pointCount{
		min := inf
		for i:=1; i<=pointCount; i++{
			if book[i] == 0 && dis[i] < min{
				min = dis[i]
				j = i
			}
		}
		book[j] = 1
		count++
		sum += dis[j]
		for i, v := range dis{
			if book[i] == 0 && v > gameMap[j][i]{
				dis[i] = gameMap[j][i]
			}
		}
	}
	fmt.Println(sum)
}