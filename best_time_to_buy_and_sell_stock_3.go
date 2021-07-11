package main
import "fmt"

func max_profit(prices []int) int {
	single_transaction_profit := 0
	profit := 0
	best_buy := prices[0]
	best_buy_with_transaction := prices[0]
	for _, price := range prices[1:] {
		if price - best_buy > single_transaction_profit {
			single_transaction_profit = price - best_buy
		}
		if price - best_buy_with_transaction > profit {
			profit = price - best_buy_with_transaction
		}
		if price < best_buy {
			best_buy = price
		}
		if price - single_transaction_profit < best_buy_with_transaction {
			best_buy_with_transaction = price - single_transaction_profit
		}
	}
	return profit
}

func main() {
	fmt.Println(max_profit([]int{3,3,5,0,0,3,1,4}))
	fmt.Println(max_profit([]int{1,2,3,4,5}))
	fmt.Println(max_profit([]int{7,6,4,3,1}))
	fmt.Println(max_profit([]int{1}))
	fmt.Println(max_profit([]int{4,5,2,3,0,1}))
	fmt.Println(max_profit([]int{6,7,4,5,2,3,0,1}))
	fmt.Println(max_profit([]int{3,3,5,0,0,3,1,4,6}))
	fmt.Println(max_profit([]int{5,6,3,4,0,2}))
}
