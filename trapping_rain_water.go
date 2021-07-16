package main
import "fmt"

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    return -max(-x, -y)
}

func trap(height []int) int {
    if len(height) == 0 {
        return 0
    }
    right_max := make([]int, len(height))
    right_max[len(height)-1] = height[len(height)-1]
    for i := len(height)-2; i >= 0; i-- {
        right_max[i] = max(right_max[i+1], height[i])
    }
    left_max := height[0]
    sum := 0
    for i, h := range height {
        left_max = max(left_max, h)
        sum += min(left_max, right_max[i]) - h
    }
    return sum
}

func main() {
    fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
    fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
    fmt.Println(trap([]int{3, 0, 5, 0, 1, 1, 2, 3, 0}))
}
