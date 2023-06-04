/*
课后练习1.1
编写一个小程序：
给定一个字符串数组
["I","am","stupid","and","weak"]
用for循环遍历数组，并修改为
["I","am","smart","and","strong"]
*/
package main

import "fmt"

func main() {
	arr := [5]string{"I", "am", "stupid", "and", "weak"}
	for k, v := range arr {
		if v == "stupid" {
			arr[k] = "smart"
		} else if v == "weak" {
			arr[k] = "strong"
		} else {

		}
	}
	fmt.Println(arr)
}
