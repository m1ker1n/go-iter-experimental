package main

import (
	"fmt"
	"iter"
	"math/rand"
)

// RandomOrderIterator iterates through s in random order returning index and value of each element.
func RandomOrderIterator[S ~[]E, E any](s S) iter.Seq2[int, E] {
	order := rand.Perm(len(s))
	return func(yield func(int, E) bool) {
		for _, v := range order {
			if !yield(v, s[v]) {
				return
			}
		}
	}
}

func main() {
	s1 := []int{5, 10, 4, 7, 3}
	for i, v := range RandomOrderIterator(s1) {
		fmt.Printf("s1[%d]=%v\n", i, v)
	}
	fmt.Println()

	s2 := []string{"five", "ten", "four", "seven", "three"}
	for i, v := range RandomOrderIterator(s2) {
		fmt.Printf("s2[%d]=%v\n", i, v)
	}
	fmt.Println()

	type someStruct struct {
		i   int
		str string
	}

	s3 := []someStruct{
		{5, "five"},
		{10, "ten"},
		{4, "four"},
		{7, "seven"},
		{3, "three"},
	}
	for i, v := range RandomOrderIterator(s3) {
		fmt.Printf("s3[%d]=%v\n", i, v)
	}
}
