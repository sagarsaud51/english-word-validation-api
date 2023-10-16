package main

import (
	"bufio"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
			}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	path := "./words.txt"
	router := gin.Default()
	wordsTrie := NewTrie()

	wordList, err := readFile(path)

	if err != nil {
		log.Fatalf("Could not read file %s with error %s", path, err)
		os.Exit(1)
	}

	for _, word := range wordList {
		wordsTrie.Insert(word)
	}

	router.GET("/v1/word/valid", func(c *gin.Context) {
		word := c.Query("word")
		valid := wordsTrie.Search(word)

		c.JSON(200, gin.H{
			"valid": valid,
		})
	})

	router.Run(":8080")
}
