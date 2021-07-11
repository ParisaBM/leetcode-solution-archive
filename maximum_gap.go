package main
import "fmt"

func maximum_gap(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	min := nums[0]
	max := nums[0]
	for _, x := range nums[1:] {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	bin_width := (max-min + len(nums)-2) / (len(nums)-1)
	if bin_width == 0 {
		return 0
	}
	mins := make([]int, len(nums))
	maxs := make([]int, len(nums))
	for i := range nums {
		bin := (nums[i] - min) / bin_width
		if nums[i] < mins[bin] || mins[bin] == 0 {
			mins[bin] = nums[i]
		}
		if nums[i] > maxs[bin] {
			maxs[bin] = nums[i]
		}
	}
	found := 0
	previous := min
	for i := range mins {
		if mins[i] != 0 {
			if mins[i] - previous > found {
				found = mins[i] - previous
			}
			previous = maxs[i]
		}
	}
	return found
}

func main() {
	fmt.Println(maximum_gap([]int{3,4,1}))
}
