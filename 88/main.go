package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	q := m - 1

	l := n - 1
	if n == 0 {
		return
	}
	for {
		fmt.Println(p, q, l)

		if q < 0 && l < 0 {
			break
		}
		if q == -1 {
			nums1[0] = nums2[l]
		}
		if l == -1 {
			break
		}
		if l > -1 && q > -1 && nums1[q] > nums2[l] {
			nums1[p] = nums1[q]
			q -= 1
		} else {
			nums1[p] = nums2[l]
			l -= 1
		}
		p -= 1
		fmt.Println(nums1)

	}
	fmt.Println(nums1)
}

func main() {
	nums1 := []int{0, 0, 0, 0, 0}

	m := 0

	nums2 := []int{1, 2, 3, 4, 5}

	n := 5
	merge(nums1, m, nums2, n)

}
