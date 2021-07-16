package main
import "fmt"

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func min_distance(word1 string, word2 string) int {
    table := make([][]int, len(word1)+1)
    for i := range table {
        table[i] = make([]int, len(word2)+1)
        for j := range table[i] {
            if i == 0 {
                table[i][j] = j
            } else if j == 0 {
                table[i][j] = i
            } else if word1[i-1] == word2[j-1] {
                table[i][j] = table[i-1][j-1]
            } else {
                table[i][j] = min(min(table[i-1][j], table[i][j-1]), table[i-1][j-1]) + 1
            }
        }
    }
    return table[len(word1)][len(word2)]
}

func main() {
    fmt.Println(min_distance("horse", "ros"))
    fmt.Println(min_distance("intention", "execution"))
}
