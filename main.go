package main

import (
	"fmt"
	"iter"
	"math/rand"
)

//func ReadTextFileByLines(r io.Reader) iter.Seq2[string, error] {
//	scanner := bufio.NewScanner(r)
//	return func(yield func(string, error) bool) {
//		for scanner.Scan() {
//			if !yield(scanner.Text(), nil) {
//				return
//			}
//		}
//		if err := scanner.Err(); err != nil {
//			yield("", err)
//		}
//	}
//}
//
//func main1() {
//	file, err := os.Open("README.md")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for text, err := range ReadTextFileByLines(file) {
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(text)
//	}
//}

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
	s := []int{5, 10, 4, 7, 3}
	for i, v := range RandomOrderIterator(s) {
		fmt.Printf("s[%d]=%d\n", i, v)
	}
}
