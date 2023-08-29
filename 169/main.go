package main

import "fmt"

func majorityElement(nums []int) int {
	var num int
	var lengh int = len(nums) / 2
	numj := make(map[int]int)

	for _, j := range nums {
		numj[j]++
	}
	for i, j := range numj {
		fmt.Println(i, j)
		if j > lengh {
			return i
		}
	}

	return num
}

func main() {
	nums := []int{3, 2, 3}
	majorityElement(nums)
}
