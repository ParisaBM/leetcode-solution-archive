package main
import "fmt"

func shortest_palindrome(s string) string {
    prefix := ""
    outer: for i := len(s)-1; i >= 0; i-- {
        for j := 0; j < (i+1)/2; j++ {
            if s[j] != s[i-j] {
                prefix += string(s[i])
                continue outer
            }
        }
        break
    }
    return prefix+s
}

func main() {
    fmt.Println(shortest_palindrome("aacecaaa"))
    fmt.Println(shortest_palindrome("abcd"))
    fmt.Println(shortest_palindrome("abab"))
    fmt.Println(shortest_palindrome(""))
    fmt.Println(shortest_palindrome("1"))
    fmt.Println(shortest_palindrome("11"))
}
