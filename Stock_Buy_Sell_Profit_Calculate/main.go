package main

import "fmt"

var newprofit int
var profit int

func main() {
	buyprice := []int{7, 1, 5, 3, 6, 4} //Array With Profit 5
	//buyprice := []int{2, 4, 8, 10, 45, 60, 1, 2} //Array With Profit 1
	//buyprice := []int{7, 6, 4, 3, 1} //Array With No profit 0
	maxprofit := Calprofit(buyprice)
	fmt.Println(maxprofit)
}
func Calprofit(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	for i := 0; i <= len(arr); i++ { //buy
		for j := i + 1; j < len(arr); j++ { //sell
			if arr[j] > arr[i] {
				newprofit = arr[j] - arr[i] //sell - buy
			}
			if newprofit > profit { // Use to check, is previous is greater than new!
				profit = newprofit //If yes than store new profit value to temp
			}

		}

	}
	return profit
}
