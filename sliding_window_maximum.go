package main
import "fmt"

type Item struct {
	Position, Value int
}

func max_sliding_window(nums []int, k int) []int {
	output := make([]int, 0, len(nums)-k+1)
	buffer := make([]Item, k)
	base := 0
	length := 0
	for i := range nums {
		for length > 0 && buffer[(base+length-1) % k].Value <= nums[i] {
			length--
		}
		if length > 0 && buffer[base].Position == i - k {
			base = (base + 1) % k
			length--
		}
		buffer[(base+length) % k] = Item{i, nums[i]}
		length++
		if i >= k - 1 {
			output = append(output, buffer[base].Value)
		}
	}
	return output
}

func main() {
	fmt.Println(max_sliding_window([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(max_sliding_window([]int{1}, 1))
	fmt.Println(max_sliding_window([]int{1, -1}, 1))
	fmt.Println(max_sliding_window([]int{9, 11}, 2))
	fmt.Println(max_sliding_window([]int{4, -2}, 2))
}
