package main

func rotate(nums []int, k int) {

	var lengh = len(nums)
	if lengh < 2 {
		return
	}
	k = k % (lengh)
	c := make([]int, lengh)
	copy(c, nums)
	for p, _ := range nums {
		nums[p] = c[(lengh-k+p)%lengh]

	}

}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 2
	rotate(nums, k)
}
