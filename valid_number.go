package main
import (
    "fmt"
    "regexp"
)

func is_number(s string) bool {
    re := regexp.MustCompile("^[+-]?((\\d+(\\.\\d*)?)|(\\.\\d+))([eE][+-]?\\d+)?$")
    match := re.MatchString(s)
    return match
}

func main() {
    fmt.Println(is_number("2"))
    fmt.Println(is_number("0089"))
    fmt.Println(is_number("-0.1"))
    fmt.Println(is_number("+3.14"))
    fmt.Println(is_number("4."))
    fmt.Println(is_number("-.9"))
    fmt.Println(is_number("2e10"))
    fmt.Println(is_number("-90E3"))
    fmt.Println(is_number("3e+7"))
    fmt.Println(is_number("+6e-1"))
    fmt.Println(is_number("53.5e93"))
    fmt.Println(is_number("-123.456e789"))
    fmt.Println(is_number("abc"))
    fmt.Println(is_number("1a"))
    fmt.Println(is_number("1e"))
    fmt.Println(is_number("e3"))
    fmt.Println(is_number("99e2.5"))
    fmt.Println(is_number("--6"))
    fmt.Println(is_number("-+3"))
    fmt.Println(is_number("95a54e53"))
}
