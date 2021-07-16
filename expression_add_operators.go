package main
import "fmt"

func add_operators(num string, target int) []string {
    all_solutions := []string{}
    Outer: for ops := 0; ops < 1 << ((len(num)-1)*2); ops++ {
        val_stack := make([]int, 1)
        op_stack := make([]int, 0)
        val_stack[0] = int(num[0]-'0')
        for i := 1; i < len(num); i++ {
            switch op := (ops >> ((i-1)*2)) & 0x3; op {
            case 0: //concat
                if val_stack[len(val_stack)-1] == 0 {
                    continue Outer
                }
                val_stack[len(val_stack)-1] *= 10
                val_stack[len(val_stack)-1] += int(num[i]-'0')
            case 1: //multiply
                op_stack = append(op_stack, 1)
                val_stack = append(val_stack, int(num[i]-'0'))
            default:
                for len(op_stack) > 0 {
                    switch op_stack[len(op_stack)-1] {
                    case 1:
                        val_stack[len(val_stack)-2] *= val_stack[len(val_stack)-1]
                    case 2:
                        val_stack[len(val_stack)-2] += val_stack[len(val_stack)-1]
                    case 3:
                        val_stack[len(val_stack)-2] -= val_stack[len(val_stack)-1]
                    }
                    op_stack = op_stack[:len(op_stack)-1]
                    val_stack = val_stack[:len(val_stack)-1]
                }
                op_stack = append(op_stack, op)
                val_stack = append(val_stack, int(num[i]-'0'))
            }
        }
        for len(op_stack) > 0 {
            switch op_stack[len(op_stack)-1] {
            case 1:
                val_stack[len(val_stack)-2] *= val_stack[len(val_stack)-1]
            case 2:
                val_stack[len(val_stack)-2] += val_stack[len(val_stack)-1]
            case 3:
                val_stack[len(val_stack)-2] -= val_stack[len(val_stack)-1]
            }
            op_stack = op_stack[:len(op_stack)-1]
            val_stack = val_stack[:len(val_stack)-1]
        }
        if val_stack[0] == target {
            solution := ""
            for i := range num {
                solution += string(num[i])
                switch op := (ops >> (i*2)) & 0x3; op {
                case 1:
                    solution += "*"
                case 2:
                    solution += "+"
                case 3:
                    solution += "-"
                }
            }
            all_solutions = append(all_solutions, solution)
        }
    }
    return all_solutions
}

func main() {
    fmt.Println(add_operators("123", 6))
    fmt.Println(add_operators("232", 8))
    fmt.Println(add_operators("105", 5))
    fmt.Println(add_operators("00", 0))
    fmt.Println(add_operators("3456237490", 9191))
}
