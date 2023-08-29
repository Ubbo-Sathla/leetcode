package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	var num int = 2

	for i, j := range nums {
		if i > 1 {
			if nums[num-2] != j {
				nums[num] = j
				num++
			}
		}
	}
	fmt.Println(nums)
	return num

}
func main() {
	nums := []int{1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 4, 5, 5}
	removeDuplicates(nums)
}
