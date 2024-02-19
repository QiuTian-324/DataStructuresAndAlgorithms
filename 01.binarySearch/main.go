package main

import (
	"fmt"
)

func main() {
	var (
		target = 3
		arr    = []int{1, 2, 3, 3, 4, 5, 6, 7, 8}
	)

	var (
		nums = []int{-1, 0, 3, 5, 9, 12}
		tar  = 9
	)

	fmt.Println(binarySearch(arr, target))
	fmt.Println(leftMost(arr, target))
	fmt.Println(rightMost(arr, target))
	fmt.Println(search(nums, tar))
}

func binarySearch(arr []int, t int) bool {
	var (
		firstIndex = 0
		lastIndex  = len(arr)
	)
	for firstIndex < lastIndex {
		m := (firstIndex + lastIndex) >> 1
		// 判断目标值与中间值的关系
		if t < arr[m] {
			lastIndex = m
		} else if t > arr[m] {
			firstIndex = m
		} else {
			return true
		}
	}
	return false
}

func leftMost(arr []int, t int) int {
	// 定义数组的起始和结束索引
	var (
		firstIndex = 0
		lastIndex  = len(arr) - 1
	)

	// 使用二分查找算法
	// 这里判断为小于等于是考虑到first和last会有相等的情况
	for firstIndex <= lastIndex {
		// 计算中间索引
		m := (firstIndex + lastIndex) / 2

		// 判断目标值与中间值的关系
		if t <= arr[m] {
			lastIndex = m - 1
		} else {
			firstIndex = m + 1
		}
	}

	// 未找到目标值，返回false
	return firstIndex
}

func rightMost(arr []int, t int) int {
	// 定义数组的起始和结束索引
	var (
		firstIndex = 0
		lastIndex  = len(arr) - 1
	)

	// 使用二分查找算法
	// 这里判断为小于等于是考虑到first和last会有相等的情况
	for firstIndex <= lastIndex {
		// 计算中间索引
		m := (firstIndex + lastIndex) / 2

		// 判断目标值与中间值的关系
		if t < arr[m] {
			lastIndex = m - 1
		} else {
			firstIndex = m + 1
		}
	}

	// 未找到目标值，返回false
	return firstIndex - 1
}

func search(nums []int, target int) int {
	i := 0
	j := len(nums) - 1
	index := -1

	for i <= j {
		m := (i + j) >> 1
		if target < nums[m] {
			j = m - 1
		} else if target > nums[m] {
			i = m + 1
		} else {
			return m
		}
	}
	return index
}
