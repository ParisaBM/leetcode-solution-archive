package main
import (
	"fmt"
	"strconv"
	"unicode"
)

func calculate(s string) int {
	total := 0
	parity := 1 //1 or -1
	int_buffer := ""
	bracket_stack := make([]int, 0)
	last := 1
	for _, c := range s {
		if unicode.IsDigit(c) {
			int_buffer += string(c)
		} else {
			if len(int_buffer) > 0 {
				buffer_value, _ := strconv.Atoi(int_buffer)
				total += last * parity * buffer_value
				int_buffer = ""
			}
			switch c {
			case '+':
				last = 1
			case '-':
				last = -1
			case '(':
				bracket_stack = append(bracket_stack, last)
				parity *= last
				last = 1
			case ')':
				parity *= bracket_stack[len(bracket_stack) - 1]
				bracket_stack = bracket_stack[:len(bracket_stack) - 1]
			}
		}
	}
	if len(int_buffer) > 0 {
		buffer_value, _ := strconv.Atoi(int_buffer)
		total += last * parity * buffer_value
	}
	return total
}

func main() {
	fmt.Println(calculate("1 + 1"))
	fmt.Println(calculate("2-1 + 2"))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
	fmt.Println(calculate("+48 + -48"))
	fmt.Println(calculate("1-((1+1)+2)"))
}
