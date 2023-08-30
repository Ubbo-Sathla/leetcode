package main

import "fmt"

func canJump(nums []int) bool {

	var lengh int = len(nums) - 1
	var max int = 0
	for i, j := range nums {
		//fmt.Println("i,j:", i, j)

		if i+j < lengh {
			if nums[i+j] == 0 {

				if i+j > i {

				} else {
					for p, q := range nums {
						if p < i {
							//fmt.Println("p,q:", p, q)
							if max < p+q {
								max = p + q
							}
						}

					}
					//fmt.Println(max)
					if max <= i {
						return false
					}
					max = 0
				}

			}
		} else {
			return true
		}
	}

	return false
}
func main() {
	//nums := []int{2, 5, 0, 0}
	nums := []int{1, 1, 2, 2, 0, 1, 1}
	//nums := []int{0, 1}
	//nums := []int{0}
	fmt.Println(canJump(nums))
}
