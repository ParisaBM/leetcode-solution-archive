package main
import "fmt"

func max_profit(k int, prices []int) int {
	best_buys := make([]int, k)
	for i := range best_buys {
		best_buys[i] = 1001
	}
	max_profits := make([]int, k+1)
	for _, price := range prices {
		for i := range best_buys {
			if price - max_profits[i] < best_buys[i] {
				best_buys[i] = price - max_profits[i]
			}
		}
		for i := range max_profits[1:] {
			if price - best_buys[i] > max_profits[i+1] {
				max_profits[i+1] = price - best_buys[i]
			}
		}
	}
	return max_profits[k]
}

func main() {
	fmt.Println(max_profit(2, []int{2, 4, 1}))
	fmt.Println(max_profit(2, []int{3, 2, 6, 5, 0, 3}))
	p := make([]int, 1000)
	for i := range p {
		p[i] = i % 10
	}
	fmt.Println(max_profit(100, p))
	fmt.Println(max_profit(0, p))
	fmt.Println(max_profit(2, []int{4}))
	fmt.Println(max_profit(2, []int{}))
}
