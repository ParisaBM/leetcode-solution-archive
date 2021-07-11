package main
import (
	"fmt"
	"math/rand"
	"time"
)

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func trap_rain_water(heightMap [][]int) int {
 	water_level := make([][]int, len(heightMap))
 	for i := range water_level {
 		water_level[i] = make([]int, len(heightMap[0]))
 	}
 	for i := range water_level {
 		for j := range water_level[0] {
 			if i==0 || i==len(water_level)-1 || j==0 || j==len(water_level[0])-1 {
 					water_level[i][j] = heightMap[i][j]
 			} else {
 				water_level[i][j] = 20000
 			}
 		}
 	}
 	made_change := true
 	for made_change {
 		made_change = false
 		for i := 1; i < len(water_level)-1; i++ {
 			for j := 1; j < len(water_level[0])-1; j++ {
 				offsets := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
 				for _, offset := range offsets {
					if water_level[i+offset[0]][j+offset[1]] < water_level[i][j] && heightMap[i][j] < water_level[i][j] {
						water_level[i][j] = max(water_level[i+offset[0]][j+offset[1]], heightMap[i][j])
						made_change = true
					}
				}
			}
		}
	}
	volume := 0
	for i := range water_level {
		for j := range water_level[0] {
			volume += water_level[i][j] - heightMap[i][j]
		}
	}
	return volume
}

func main() {
	fmt.Println(trap_rain_water([][]int{{3,3,3,3,3},{3,2,2,2,3},{3,2,1,2,3},{3,2,2,2,3},{3,3,3,3,3}}))
	rand.Seed(time.Now().UnixNano())
	problem := make([][]int, 200)
	 for i := range problem {
	 	problem[i] = make([]int, 200)
	 	for j := range problem[i] {
	 		problem[i][j] = rand.Intn(100)
	 	}
	}
	fmt.Println(trap_rain_water(problem))
}
