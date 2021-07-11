package main
import "fmt"

func num_distinct(s string, t string) int {
	num_prefix := make([]int, len(t)+1)
	num_prefix[0] = 1
	for i := range s {
		for j := len(t); j >= 1; j-- {
			if t[j - 1] == s[i] {
				num_prefix[j] += num_prefix[j-1]
			}
		}
	}
	return num_prefix[len(t)]
}

func main() {
	fmt.Println(num_distinct("rabbbit", "rabbit"))
	fmt.Println(num_distinct("babgbag", "bag"))
}
