package main
import "fmt"

func find_min(nums []int) int {
	var lower int
	upper := len(nums) - 1
	if nums[0] < nums[len(nums) - 1] {
		return nums[0]
	}
	for i := range nums {
		if i == len(nums) - 1 {
			return nums[0]
		}
		if nums[i] != nums[i+1] {
			if nums[i+1] < nums[i] {
				return nums[i+1]
			}
			if i == len(nums) - 2 {
				return nums[i]
			}
			lower = i + 1
		}
	}
	for upper - lower > 1 {
		mid := (lower + upper) / 2
		if nums[mid] <= nums[0] {
			upper = mid
		} else {
			lower = mid
		}
	}
	return nums[lower]
}

func main() {
	fmt.Println(find_min([]int{1, 3, 5}))
	fmt.Println(find_min([]int{2,2,2,0,1}))
	fmt.Println(find_min([]int{4,5,6,7,0,1,2,3}))
	fmt.Println(find_min([]int{4,4,4,4,5,6,7,0,1,2,3,4,4,4,4,4,4,4,4,4,4,4,4}))
}
