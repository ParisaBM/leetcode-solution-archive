package main

func gcf(x int, y int) int {
	for y != 0 {

	}
	return x
}

func find_line(points [][]int) [3]int {
	// given two points returns a unique line passing through both
	// line given in the form ax + by + c = 0
	a := points[1][1] - points[0][1]
	b := points[0][0] - points[1][0]
	c := -a*points[0][0] - b*points[0][1]
	dividend := gcf(gcf(a, b), c)
	a /= dividend
	b /= dividend
	c /= dividend
	return [3]int{a, b, c}
}

func max_points(points [][]int) int {
	lines := make(map[[3]int]int) // counts the number of points on each line
	max := 1                      // max points found so far
	for i := 1; i < len(points); i += 1 {
		// increments the lines that the point is on, and creates new lines that the point is on
		for coef := range lines {
			if points[i][0]*coef[0]+points[i][1]*coef[1]+coef[2] == 0 {
				lines[coef] += 1
				if lines[coef] > max {
					max = lines[coef]
				}
			}
		}
		for j := 0; j < i; j += 1 {
			new_line := find_line([][]int{points[j], points[i]})
			_, ok := lines[new_line]
			if !ok {
				lines[new_line] = 2
				if 2 > max {
					max = 2
				}
			}
		}
	}
	return max
}

func main() {
	println(max_points([][]int{}))
	println(max_points([][]int{{0, 0}}))
	println(max_points([][]int{{0, 0}, {1, 2}}))
	println(max_points([][]int{{0, 0}, {1, 2}, {2, 4}}))
	println(max_points([][]int{{1, 1}, {2, 2}, {3, 3}}))
	println(max_points([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))
}
