package function

import "fmt"

func Replace(words *[]string, oldWord, newWord string) {
	count := 0
	indices := []int{}

	for i := 0; i < len(*words); i++ {
		if (*words)[i] == oldWord {
			(*words)[i] = newWord
			count++
			indices = append(indices, i)
		}
	}

	if count > 0 {
		fmt.Printf("Word '%s' replaced %d times at:\n", oldWord, count)
		for i, index := range indices {
			fmt.Printf("%d: %s [word %d]\n", i+1, newWord, index)
		}
	} else {
		fmt.Println("Word not found!")
	}
}