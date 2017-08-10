//队列
//钓鱼游戏，没牌的是输
//全部用数组， 更好体现链表

package main

import (
	"fmt"
)

//桌面上的牌, 是一个栈
type DesktopList struct{
	data [100]int
	top int
}

//两人手里的牌
type HandList struct{
	name string
	head int
	tail int
	data [100]int
}

func main(){
	//初始化
	aHandList := HandList{name:"A",head:0, tail:5}
	for k, v := range [6]int{2,4,1,2,5,6}{
		aHandList.data[k] = v
	}
	bHandList := HandList{name:"B", head:0, tail:5}
	for k, v := range [6]int{3,1,3,5,6,4}{
		bHandList.data[k] = v
	}
	//初始化 桌面没牌
	desktopList := DesktopList{top:0}
	desktopList.data[desktopList.top] = 99
	desktopList.top++
	desktopList.data[desktopList.top] = 2
	desktopList.top++
	desktopList.data[desktopList.top] = 10
	desktopList.top++
	desktopList.data[desktopList.top] = 11
	desktopList.top++
	desktopList.data[desktopList.top] = 12
	//游戏开始， A先出牌
	startUser := aHandList
	for aHandList.head < aHandList.tail && bHandList.head < bHandList.tail{
		current := startUser.data[startUser.head]
		startUser.head++
		//把牌放到桌上
		desktopList.top++
		desktopList.data[desktopList.top] = current
		//判断是否可以吃牌了
		flag := 0
		for i:=0; i<desktopList.top-1; i++{
			if current == desktopList.data[i]{
				flag = i
				break
			}
		}
		//吃牌
		if flag != 0{
			for i:=desktopList.top; i>=flag; i--{
				//牌入手
				startUser.tail++
				startUser.data[startUser.tail] = desktopList.data[i]
				//清桌面
				desktopList.top--
			}
		}else{
			//换人出牌
			if startUser.name == "A"{
				startUser = bHandList
			}else{
				startUser = aHandList
			}
		}
	}
	fmt.Println("桌面：")
	for i:=desktopList.top; i>=0;i--{
		fmt.Println(desktopList.data[i])
	}
	fmt.Println("A的手牌：")
	for i:=aHandList.head; i<=aHandList.tail;i++{
		fmt.Println(aHandList.data[i])
	}
	fmt.Println("B的手牌：")
	for i:=bHandList.head; i<=bHandList.tail;i++{
		fmt.Println(bHandList.data[i])
	}
}