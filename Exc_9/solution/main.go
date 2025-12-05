package main

import (
	"bufio"
	"exc9/mapred" //was there
	"fmt"
	"log"
	"os"
	"sort"
)

// Main function
func main() {
	// todo read file
	file, err := os.Open("res/meditations.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close() //close the file at the end of the function

	//scan file:
	var text []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner error: %v", err)
	}

	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(text)

	// todo print your result to stdout

	// Print number of unique words
	fmt.Printf("Unique words: %d\n\n", len(results))

	// Print top 20 most frequent words
	type kv struct {
		Key   string
		Value int
	}
	kvs := make([]kv, 0, len(results))
	for k, v := range results {
		kvs = append(kvs, kv{k, v})
	}

	// Sort by frequency descending
	sort.Slice(kvs, func(i, j int) bool {
		return kvs[i].Value > kvs[j].Value
	})

	fmt.Println("Top 20 most frequent words:")
	for i := 0; i < 20 && i < len(kvs); i++ {
		fmt.Printf("%s: %d\n", kvs[i].Key, kvs[i].Value)
	}
}
