package main

import "fmt"

func maxProfit(prices []int) int {
	var min int = prices[0]
	var max int = 0
	var sum int = 0
	var lengh int = len(prices)
	if lengh < 2 {
		return 0
	}
	for _, j := range prices {

		if max >= j {
			sum = sum + max - min
			min = j
			max = j
		} else {
			max = j
		}
		if min > j {
			min = j
		}
		fmt.Println(min, max, j, sum)

	}
	fmt.Println(min, max, sum, prices[lengh-1], prices[lengh-2])
	if prices[lengh-1] > prices[lengh-2] {
		sum = sum + max - min
	}
	fmt.Println(sum)
	return sum
}
func main() {
	prices := []int{1, 2, 3, 4, 5}
	maxProfit(prices)
}
