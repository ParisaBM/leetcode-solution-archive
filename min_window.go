package main
import "fmt"

func min_window(s string, t string) string {
    result := ""
    t_tally := make(map[byte]int)
    for i := range t {
        t_tally[t[i]]++
    }
    w_tally := make(map[byte]int)
    deficiencies := len(t_tally)
    i := 0
    j := 0
    for {
        if deficiencies > 0 {
            if j == len(s) {
                return result
            }
            w_tally[s[j]]++
            if w_tally[s[j]] == t_tally[s[j]] {
                deficiencies--
            }
            j++
        } else {
            if w_tally[s[i]] == t_tally[s[i]] {
                deficiencies++
            }
            w_tally[s[i]]--
            i++
        }
        if deficiencies == 0 && (result == "" || j-i < len(result)) {
            result = s[i:j]
        }
    }
}

func main() {
    fmt.Println(min_window("ADOBECODEBANC", "CBD"))
    fmt.Println(min_window("ADOBECODEBANC", "ABC"))
    fmt.Println(min_window("a", "a"))
    fmt.Println(min_window("a", "aa"))
}
