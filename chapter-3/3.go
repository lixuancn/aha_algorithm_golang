//火柴棍拼数字
package main

import "fmt"

func main(){
	//总共火柴棍个数
	total := 18
	count := 0
	for i:=0; i<=1111; i++{
		for j:=0; j<1111; j++{
			m := i + j
			if (getNumByInt(i) + getNumByInt(j) + getNumByInt(m)) == (total - 2 -2){
				count ++
				fmt.Println(fmt.Sprintf("组合为%d+%d=%d", i, j, m))
			}
		}
	}
	fmt.Println(fmt.Sprintf("共有%d个组合", count))
}

//输入一个数，获得该数需要几个火柴棍
func getNumByInt(num int)int{
	//0-9每个数字所需要的火柴棍的个数
	intAndNumMap := [10]int{6, 2, 5,5,4,5,6,3,7,6}
	sum := 0
	if num == 0{
		sum = intAndNumMap[num]
	}
	for num > 0{
		n := num % 10
		num /= 10
		sum += intAndNumMap[n]
	}
	return sum
}