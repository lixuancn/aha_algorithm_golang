//栈
//验证回文字符串，回复字符串就是正读反读都一样

package main

import (
	"fmt"
)

func main(){
	str := "xyxyx"
	strlen := len(str)
	//mid是字符串的中间点，默认字符串长度是偶数
	mid := strlen / 2
	top := mid - 1
	//奇数长度
	if strlen%2 != 0{
		mid = strlen / 2 + 1
		top = mid - 2
	}
	stackList := []string{}
	//mid之前（奇数含mid）的入栈
	for i:=0; i<mid; i++{
		stackList = append(stackList, string(str[i]))
	}
	for i:=mid; i<strlen; i++{
		if string(str[i]) != stackList[top]{
			fmt.Println(str + "不是回文字符串")
			return
		}
		top --
	}
	fmt.Println(str + "是回文字符串")
}