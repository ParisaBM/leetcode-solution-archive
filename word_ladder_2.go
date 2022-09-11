package main

import "fmt"

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

func generate_paths(word_list []string, predecessors [][]int, end_word_index int) [][]string {
	if len(predecessors[end_word_index]) == 0 {
		return [][]string{{word_list[end_word_index]}}
	}
	paths := make([][]string, 0)
	for i := range predecessors[end_word_index] {
		paths = append(paths, generate_paths(word_list, predecessors, predecessors[end_word_index][i])...)
	}
	for i := range paths {
		paths[i] = append(paths[i], word_list[end_word_index])
	}
	return paths
}

func find_ladders(begin_word string, end_word string, word_list []string) [][]string {
	word_list = append(word_list, begin_word)
	distance := make([]int, len(word_list)) // distance from begin_word
	distance[len(word_list)-1] = 1
	visited_queue := []int{len(word_list) - 1} // nodes to take outgoing paths from
	predecessors := make([][]int, len(word_list))
	for i := range predecessors {
		predecessors[i] = make([]int, 0)
	}
	for i := 0; i < len(visited_queue); i += 1 {
		for j := range word_list {
			if (distance[j] == 0 || distance[j] > distance[visited_queue[i]]) && are_adjacent(word_list[visited_queue[i]], word_list[j]) {
				if distance[j] == 0 {
					visited_queue = append(visited_queue, j)
				}
				distance[j] = distance[visited_queue[i]] + 1
				predecessors[j] = append(predecessors[j], visited_queue[i])
			}
		}
		if word_list[visited_queue[i]] == end_word {
			return generate_paths(word_list, predecessors, visited_queue[i])
		}
	}
	return [][]string{} // means no path was found
}

func main() {
	fmt.Printf("%v\n", all_ladders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
