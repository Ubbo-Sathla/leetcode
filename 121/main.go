package main

func maxProfit(prices []int) int {
	var min int = prices[0]
	var max int = 0
	for _, j := range prices {
		if min > j {
			min = j
		}
		if j-min > max {
			max = j - min
		}
	}
	return max
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	maxProfit(prices)
}
