//生成最小树 - Prim算法
package main

import (
	"math"
	"fmt"
)

var pointCount = 6
var lineCount = 9

var userInputMap = [9][3]int{
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

var book [7]int
var dis [7]int
//h是堆， pos是每个定点在堆中的位置
var h, pos [7]int

func swap(x, y int){
	tmp := h[x]
	h[x] = h[y]
	h[y] = tmp

	tmp = pos[h[x]]
	pos[h[x]] = pos[h[y]]
	pos[h[y]] = tmp
}

func siftdown(i int){
	var flag bool
	var t int
	for i * 2 <= pointCount && flag == false{
		if dis[h[i]] > dis[h[i*2]]{
			t = i * 2
		}else{
			t = i
		}
		if i * 2 + 1 <= pointCount{
			if dis[h[t]] > dis[h[i*2+1]]{
				t = i * 2 + 1
			}
		}
		if t != i{
			swap(t, i)
			i = t
		}else{
			flag = true
		}
	}
}

func siftup(i int){
	if i == 1{
		return
	}
	var flag bool
	for i != 1 && flag == false{
		if dis[h[i]] < dis[h[i/2]]{
			swap(i, i/2)
		}else{
			flag = true
		}
		i = i / 2
	}
}

func pop()int{
	t := h[1]
	pos[t] = 0
	h[1] = h[pointCount]
	pointCount--
	siftdown(1)
	return t
}

func main(){
	var u, v, w [19]int
	var first [7]int
	var next [19]int
	inf := math.MaxInt64
	var count, sum int
	for k, userInput := range userInputMap{
		key := k + 1
		u[key] = userInput[0]
		v[key] = userInput[1]
		w[key] = userInput[2]
		u[key+9] = userInput[0]
		v[key+9] = userInput[1]
		w[key+9] = userInput[2]
	}
	//邻接表
	for i, _ := range first{
		first[i] = -1
	}
	for i:=1; i<=2*lineCount; i++{
		next[i] = first[u[i]]
		first[u[i]] = i
	}
	//Prim算法
	book[1] = 1
	count++
	for i, _ := range dis{
		dis[i] = inf
	}
	dis[1] = 0
	k := first[1]
	sum++
	for k!=-1{
		dis[v[k]] = w[k]
		k = next[k]
	}
	for i, _ := range h{
		h[i] = i
		pos[i] = i
	}
	for i:=pointCount/2; i>=1; i--{
		siftdown(i)
	}
	pop()
	for count < pointCount{
		j := pop()
		book[j] = 1
		count++
		sum += dis[j]
		k := first[j]
		for k != -1{
			if book[v[k]] == 0 && dis[v[k]] > w[k]{
				dis[v[k]] = w[k]
				siftup(pos[v[k]])
			}
			k = next[k]
		}
	}
	fmt.Println(sum)
}