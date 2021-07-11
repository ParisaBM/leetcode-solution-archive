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

func calculateMinimumHP(dungeon [][]int) int {
	n := len(dungeon)
	m := len(dungeon[0])
	row := make([]int, m)
	row[m-1] = max(0, -dungeon[n-1][m-1])
	for i := m-2; i >= 0; i-- {
		row[i] = max(0, row[i+1]-dungeon[n-1][i])
	}
	for i := n-2; i >= 0; i-- {
		row[m-1] = max(0, row[m-1]-dungeon[i][m-1])
		for j := m-2; j >= 0; j-- {
			row[j] = max(0, min(row[j], row[j+1])-dungeon[i][j])
		}
	}
	return row[0]+1
}

func main() {
	fmt.Println(calculateMinimumHP([][]int{{-2,-3,3},{-5,-10,1},{10,30,-5}}))
	fmt.Println(calculateMinimumHP([][]int{{0}}))
}
