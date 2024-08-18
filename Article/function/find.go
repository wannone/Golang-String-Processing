package function

import (
	"fmt"
	"strings"
)

func Find(words []string, word string) {
	count := 0
	indices := []int{}
	contexts := []string{}

	for i := 0; i < len(words); i++ {
		if strings.EqualFold(words[i], word) {
			count++
			indices = append(indices, i)

			afterContext := strings.Join(words[i+1:min(len(words), i+3)], " ")
			context := fmt.Sprintf("[%s] %s", word, afterContext)
			contexts = append(contexts, context)
		}
	}

	if count > 0 {
		fmt.Printf("Word '%s' found %d times at:\n", word, count)
		for i, context := range contexts {
			fmt.Printf("%d: %s [word %d]\n", i+1, context, indices[i])
		}
	} else {
		fmt.Println("Word not found!")
	}
}

// Helper functions to ensure the context does not go out of bounds
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}