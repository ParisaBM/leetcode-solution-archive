package main
import (
    "fmt"
    "strings"
)

func fullJustify(words []string, maxWidth int) []string {
    var lines []string
    line_start := 0
    for line_start < len(words) {
        line_end := line_start
        line_length := len(words[line_start])
        for line_end+1 < len(words) && line_length+len(words[line_end+1])+1 <= maxWidth {
            line_end++
            line_length += len(words[line_end])+1
        }
        line := ""
        for i := line_start; i <= line_end; i++ {
            line += words[i]
            if line_end+1 < len(words) && line_start != line_end {
                if i < line_end {
                    line += strings.Repeat(" ", (maxWidth-line_length+line_end-i-1)/(line_end-line_start)+1)
                }
            // left-justified
            } else {
                if i < line_end {
                    line += " "
                } else {
                    line += strings.Repeat(" ", maxWidth-line_length)
                }
            }
        }
        lines = append(lines, line)
        line_start = line_end + 1
    }
    return lines
}

func main() {
    text := fullJustify([]string{"Science","is","what","we","understand","well","enough","to","explain","to","a",
        "computer.","Art","is","everything","else","we","do"}, 20)
    for i := 0; i < len(text); i++ {
        fmt.Println(text[i])
    }
}
