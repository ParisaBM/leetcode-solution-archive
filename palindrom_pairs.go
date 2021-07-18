package main
import "fmt"

func reverse(s string) string {
    reversed := ""
    for _, c := range s {
        reversed = string(c) + reversed
    }
    return reversed
}

func isPalindrome(s string) bool {
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}

type PrefixTree struct {
    Index int
    Children [26]*PrefixTree
}

func (pt *PrefixTree) AddString(s string, index int) {
    if s == "" {
        pt.Index = index
    } else {
        nextChild := &pt.Children[s[0]-'a']
        if *nextChild == nil {
            *nextChild = &PrefixTree{Index: -1}
        }
        (*nextChild).AddString(s[1:], index)
    }
}

func (pt PrefixTree) FindMatches(s string, index int, excludeReverse bool) []int  {
    current := &pt
    matches := make([]int, 0)
    for i := len(s); i>=0; i-- {
        if current == nil || excludeReverse && i == 0 {
            break
        }
        if current.Index != -1 && current.Index != index && isPalindrome(s[:i]) {
            matches = append(matches, current.Index)
        }
        if i > 0 {
            current = current.Children[s[i-1]-'a']
        }
    }
    return matches
}

func palindromePairs(words []string) [][]int {
    pt := PrefixTree{Index: -1}
    pairs := make([][]int, 0)
    for i, word := range words {
        pt.AddString(word, i)
    }
    for i, word := range words {
        for _, match := range pt.FindMatches(word, i, false) {
            pairs = append(pairs, []int{match, i})
        }
    }
    pt = PrefixTree{Index: -1}
    for i, word := range words {
        pt.AddString(reverse(word), i)
    }
    for i, word := range words {
        for _, match := range pt.FindMatches(reverse(word), i, true) {
            pairs = append(pairs, []int{i, match})
        }
    }
    return pairs
}

func main() {
    fmt.Println(palindromePairs([]string{"abcd","dcba","lls","s","sssll"}))
    fmt.Println(palindromePairs([]string{"bat","tab","cat"}))
    fmt.Println(palindromePairs([]string{"", "a"}))
}
