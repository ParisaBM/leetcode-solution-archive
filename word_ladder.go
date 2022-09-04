package main

func are_adjacent(s0 string, s1 string) bool {
	// takes as input distinct 2 string of equal length
	// returns true iff they have exactly one different character i.e. are adjacent
	found_difference := false // tracks whether a difference been found yet
	for i := range s0 {
		if s0[i] != s1[i] {
			if !found_difference {
				found_difference = true
			} else {
				return false
			}
		}
	}
	return true
}

func ladder_length(begin_word string, end_word string, word_list []string) int {
	word_list = append(word_list, begin_word)
	distance := make([]int, len(word_list)) // distance from begin_word
	distance[len(word_list)-1] = 1
	visited_queue := []int{len(word_list) - 1} // nodes to take outgoing paths from
	for i := 0; i < len(visited_queue); i += 1 {
		for j := range word_list {
			if distance[j] == 0 && are_adjacent(word_list[visited_queue[i]], word_list[j]) {
				distance[j] = distance[visited_queue[i]] + 1
				if word_list[j] == end_word {
					return distance[j]
				}
				visited_queue = append(visited_queue, j)
			}
		}
	}
	return 0 // means no path was found
}

func main() {
	println(ladder_length("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	println(ladder_length("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}))
	println(ladder_length("hit", "cog", []string{"dot", "dog", "lot", "log", "cog"}))
}
