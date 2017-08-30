package main

import "fmt"

var peopleCount = 10
var clueCount = 9
var f [11]int

func init(){
	for i:=1; i<=peopleCount; i++{
		f[i] = i
	}
}

func getParent(v int)int{
	if f[v] == v{
		return v
	}
	f[v] = getParent(f[v])
	return f[v]
}

func merge(v, u int){
	t1 := getParent(v)
	t2 := getParent(u)
	if t1 != t2{
		f[t2] = t1
	}
}

func main(){
	groupList := [9][2]int{
		{1,2},
		{3,4},
		{5,2},
		{4,6},
		{2,6},
		{8,7},
		{9,7},
		{1,6},
		{2,4},
	}
	for _, group := range groupList{
		merge(group[0], group[1])
	}
	sum := 0
	for i:=1; i<=peopleCount; i++{
		if i == f[i]{
			fmt.Println(i)
			sum ++
		}
	}
	fmt.Println(sum)
}