package emulation

import (
	"bufio"
	"log"
	"os"
)

func load_word_list(file_path string) []string {
	words := []string{}

	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Successfully read words from file. Number of words read: %d", len(words))

	return words
}
