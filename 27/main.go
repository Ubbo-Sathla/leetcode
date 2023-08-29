package main

import "fmt"

func removeElement(nums []int, val int) int {
	var num int = 0
	var right int = len(nums) - 1
	var cur int = len(nums) - 1
	for _, _ = range nums {
		fmt.Println(nums)

		if nums[right] == val {
			nums[right] = nums[cur]
			nums[cur] = val
			cur -= 1

		} else {
			num++
		}
		right--

	}
	fmt.Println(cur)
	fmt.Println("-------")
	fmt.Println(num, nums)
	return num
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	removeElement(nums, val)
}
