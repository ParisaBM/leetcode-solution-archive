package main
import "fmt"

func maximal_rectangle(matrix [][]byte) int {
    max := 0
    for i := range matrix {
        row := make([]bool, len(matrix[i]))
        for j := range row {
            row[j] = true
        }
        for down := 0; i+down < len(matrix); down++ {
            contiguous := 0
            for across := range matrix[i+down] {
                if matrix[i+down][across] == '0' {
                    row[across] = false
                }
                if row[across] {
                    contiguous++
                    if contiguous*(down+1) > max {
                        max = contiguous*(down+1)
                    }
                } else {
                    contiguous = 0
                }
            }
        }
    }
    return max
}

func main() {
    fmt.Println(maximal_rectangle([][]byte{{'1','0','1','0','0'},{'1','0','1','1','1'},{'1','1','1','1','1'},{'1','0','0','1','0'}}))
    fmt.Println(maximal_rectangle([][]byte{{'0'}}))
    fmt.Println(maximal_rectangle([][]byte{{'1'}}))
    fmt.Println(maximal_rectangle([][]byte{{'0', '0'}}))
    fmt.Println(maximal_rectangle([][]byte{{'0','1','1'},{'1','1','1'},{'1','1','0'}}))
    fmt.Println(maximal_rectangle([][]byte{{'0','1','0'},{'1','1','1'},{'0','1','0'}}))
}
