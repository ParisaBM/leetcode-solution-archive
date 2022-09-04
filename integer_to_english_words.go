package main

var digits = [10]string{"", "One ", "Two ", "Three ", "Four ", "Five ", "Six ", "Seven ", "Eight ", "Nine "}
var teens = [10]string{"Ten ", "Eleven ", "Twelve ", "Thirteen ", "Fourteen ", "Fifteen ", "Sixteen ", "Seventeen ",
	"Eighteen ", "Nineteen "}
var tens = [8]string{"Twenty ", "Thirty ", "Forty ", "Fifty ", "Sixty ", "Seventy ", "Eighty ", "Ninety "}
var thousands = []string{"", "Thousand ", "Million ", "Billion "}

func thousand_grouping(num int) string {
	// represents a number of at most 3 digits as a string
	// 0 is represented as the empty string
	// non-zero numbers will have a trailing space at the end
	representation := ""
	if num >= 100 {
		representation += digits[num/100] + "Hundred "
	}
	if (num%100)/10 == 1 {
		representation += teens[num%10]
	} else {
		if num%100 >= 20 {
			representation += tens[(num%100)/10-2]
		}
		representation += digits[num%10]
	}
	return representation
}

func number_to_words(num int) string {
	// represents an arbitrary number as a string
	// repeatedly breaks the number into 3-digit groupings then add the appropriate suffix
	if num == 0 {
		// The number 0 is a special case, because no other number has the substring "Zero"
		return "Zero"
	}
	representation := ""
	for i := 0; num != 0; i += 1 {
		grouping := thousand_grouping(num % 1000)
		if grouping != "" {
			representation = grouping + thousands[i] + representation
		}
		num /= 1000
	}
	// before we return the result, we have to get rid of the trailing space
	return representation[:len(representation)-1]
}

func main() {
	print(number_to_words(0))
	print(number_to_words(20))
	println(number_to_words(123))
	println(number_to_words(304))
	println(number_to_words(514))
	println(number_to_words(1003))
	println(number_to_words(21000))
	println(number_to_words(12300000401))
}
