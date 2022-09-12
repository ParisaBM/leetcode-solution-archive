package main

import "fmt"

func is_power_of_two(n int) bool {
	return n&(n-1) == 0
}

func are_connected(s0 int, s1 int) bool {
	// returns true iff strings s0 and s1 represented as ints are connected
	return is_power_of_two(s0&^s1) && is_power_of_two(s1&^s0)
}

func group_strings(words []string) []int {
	words_as_ints := make([]int, len(words))
	for i := range words {
		for _, c := range words[i] {
			words_as_ints[i] = words_as_ints[i] | 1<<(c-'a')
		}
	}
	largest_group := 1
	group_size := 1
	group_end := 1
	num_groups := 1
	for i := 0; i < len(words_as_ints); i += 1 {
		if i == group_end {
			group_end += 1
			if group_size > largest_group {
				largest_group = group_size
			}
			group_size = 1
			num_groups += 1
		}
		for j := group_end; j < len(words_as_ints); j += 1 {
			if are_connected(words_as_ints[i], words_as_ints[j]) {
				words_as_ints[j], words_as_ints[group_end] = words_as_ints[group_end], words_as_ints[j]
				group_size += 1
				group_end += 1
			}
		}
	}
	if group_size > largest_group {
		largest_group = group_size
	}
	return []int{num_groups, largest_group}
}

func main() {
	fmt.Printf("%v\n", group_strings([]string{"b", "dbi", "a", "dbh", "c"}))
	fmt.Printf("%v\n", group_strings([]string{"a", "b", "ab", "cde"}))
}
