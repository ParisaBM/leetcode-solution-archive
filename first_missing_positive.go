package main
import "fmt"

func first_missing_positive(nums []int) int {
    for i := range nums {
        for 0 <= nums[i]-1 && nums[i]-1 < len(nums) && nums[nums[i]-1] != nums[i] {
            temp := nums[nums[i]-1]
            nums[nums[i]-1] = nums[i]
            nums[i] = temp
        }
    }
    for i := range nums {
        if nums[i] != i+1 {
            return i+1
        }
    }
    return len(nums)+1
}

func main() {
    fmt.Println(first_missing_positive([]int{1, 2, 0}))
    fmt.Println(first_missing_positive([]int{3, 4, -1, 1}))
    fmt.Println(first_missing_positive([]int{7, 8, 9, 11, 12}))
    fmt.Println(first_missing_positive([]int{-1}))
}
