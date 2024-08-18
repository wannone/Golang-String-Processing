package main

import (
	"article/function"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("article.txt")
	if err != nil {
		log.Fatal(err)
	}
	
	article := strings.Fields(string(content))
	exit := false

	for !exit {
		choice := Menu()
		switch choice {
		case 1:
			fmt.Println("Article: ")
			fmt.Println(strings.Join(article, " "))
		case 2:
			fmt.Print("Enter word to find: ")
			var word string
			fmt.Scanln(&word)
			function.Find(article, word)
		case 3:
			fmt.Print("Enter word to replace: ")
			var oldWord string
			fmt.Scanln(&oldWord)
			fmt.Print("Enter new word: ")
			var newWord string
			fmt.Scanln(&newWord)
			function.Replace(&article, oldWord, newWord)
		case 4:
			function.Sort(article)
		case 5:
			exit = true
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

func Menu() int {
	fmt.Println("Menu: ")
	fmt.Println("1. Article")
	fmt.Println("2. Find")
	fmt.Println("3. Replace")
	fmt.Println("4. Sort")
	fmt.Println("5. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	return choice
}
