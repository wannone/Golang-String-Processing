package function

import (
	"fmt"
	"sort"
	"strings"
)

func Sort(words []string) {
	sortedWords := make([]string, len(words))
	copy(sortedWords, words)
	sort.Strings(sortedWords)

	fmt.Println("Sorted article: ")
	fmt.Println(strings.Join(sortedWords, "\n"))
}