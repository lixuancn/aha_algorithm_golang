//最短路径 - Floyd-Warshall

package main

import "fmt"

var cityCount = 4

var userInputGameMap = [8][3]int{
	{1,2,2},
	{1,3,6},
	{1,4,4},
	{2,3,3},
	{3,1,7},
	{3,4,1},
	{4,1,5},
	{4,3,12},
}

func main(){
	var gameMap [4][4]int
	for i:=0; i<cityCount; i++{
		for j:=0; j<cityCount; j++{
			gameMap[i][j] = -1
		}
	}
	for _, line := range userInputGameMap{
		gameMap[line[0]-1][line[1]-1] = line[2]
	}
	for k:=0; k<cityCount; k++{
		for i:=0; i<cityCount; i++{
			for j:=0; j<cityCount; j++{
				if i == j {
					continue
				}
				if gameMap[i][k] == -1 || gameMap[k][j] == -1{
					continue
				}
				if gameMap[i][j] == -1 || gameMap[i][j] > gameMap[i][k] + gameMap[k][j]{
					gameMap[i][j] = gameMap[i][k] + gameMap[k][j]
				}
			}
		}
	}
	fmt.Println(gameMap)
}