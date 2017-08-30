//堆建立和排序 - 二叉树
//建立二叉树，最小堆，并从小到大输出

package main

import "fmt"

//存放堆的数组
var h [15]int
//堆的元素个数
var n int

func swap(x, y int){
	tmp := h[x]
	h[x] = h[y]
	h[y] = tmp
}

//从上向下调整
func siftDown(i int){
	var minKey, flag int
	for i*2 <= n && flag == 0{
		//和左子节点比较
		if h[i] > h[i*2]{
			minKey = i*2
		}else{
			minKey = i
		}
		//和右子节点比较
		if i*2+1 <= n{
			if h[minKey] > h[i*2+1]{
				minKey = i*2+1
			}
		}
		if minKey != i{
			swap(minKey, i)
			i = minKey
		}else{
			flag = 1
		}
	}
}

func create(){
	for i:=n/2; i>=1; i--{
		siftDown(i)
	}
}

//删除最小值，根节点永远最小
func deleteMin()int{
	t := h[1]
	h[1] = h[n]
	n--
	siftDown(1)
	return t
}

func main(){
	n = 14
	h = [15]int{0, 99, 5, 36, 7, 22, 17, 46, 12, 2, 19, 25, 28, 1, 92}

	create()
	fmt.Println(h)
	for i:=0; i<n; i++{
		fmt.Println(deleteMin())
	}
}