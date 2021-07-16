package main
import "fmt"

type Rectangle struct {
    Start, Height int
}

func largest_rectangle_area(heights []int) int {
    heights = append(heights, 0)
    max := 0
    rectangles := make([]Rectangle, 0, len(heights))
    for i, height := range heights {
        new_rectangle := Rectangle{i, height}
        for len(rectangles) > 0 && rectangles[len(rectangles)-1].Height >= height {
            new_rectangle.Start = rectangles[len(rectangles)-1].Start
            if rectangles[len(rectangles)-1].Height * (i - rectangles[len(rectangles)-1].Start) > max {
                max = rectangles[len(rectangles)-1].Height * (i - rectangles[len(rectangles)-1].Start)
            }
            rectangles = rectangles[:len(rectangles)-1]
        }
        rectangles = append(rectangles, new_rectangle)
    }
    return max
}

func main() {
    fmt.Println(largest_rectangle_area([]int{2, 1, 5, 6, 2, 3}))
    fmt.Println(largest_rectangle_area([]int{2, 4}))
    fmt.Println(largest_rectangle_area([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
}
