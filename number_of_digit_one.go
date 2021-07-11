package main
import "fmt"

func count_digit_one(n int) int {
	count := 0
	for pow := 1; pow <= n; pow *= 10 {
		count += (n / (pow * 10)) * pow
		leftover := n % (pow * 10)
		if leftover >= 2 * pow {
			count += pow
		} else if leftover >= pow {
			count += leftover - pow + 1
		}
	}
	return count
}

func main() {
	fmt.Println(count_digit_one(0))
	fmt.Println(count_digit_one(1))
	fmt.Println(count_digit_one(9))
	fmt.Println(count_digit_one(10))
	fmt.Println(count_digit_one(11))
	fmt.Println(count_digit_one(20))
	fmt.Println(count_digit_one(100))
	fmt.Println(count_digit_one(101))
}
