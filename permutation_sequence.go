package main
import (
    "fmt"
    "strconv"
)

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n*factorial(n-1)
}

func get_permutation(n int, k int) string {
    k--
    used := make([]bool, n)
    result := ""
    for i := n-1; i >= 0; i-- {
        rank := k / factorial(i)
        k %= factorial(i)
        digit := 0
        for used[digit] {
            digit++
        }
        for j := 0; j < rank; j++ {
            digit++
            for used[digit] {
                digit++
            }
        }
        used[digit] = true
        result += strconv.Itoa(digit + 1)
    }
    return result
}

func main() {
    for i := 1; i <= 24; i++ {
        fmt.Println(get_permutation(4, i))
    }
}
