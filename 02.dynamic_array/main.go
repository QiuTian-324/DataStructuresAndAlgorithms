package main

import "fmt"

// 动态数组

// 数组容量
const mCap int = 8

// 数组当前长度
var mLen int = 0

// 数组元素
var mArray = [mCap]int{}

func main() {

	// 判断index
	addLast(1)
}

// 添加元素到数组末尾
func addLast(element int) {
	mArray[mLen] = element
	fmt.Println(mArray)
	mLen++
}
