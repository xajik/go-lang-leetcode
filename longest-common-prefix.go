package main

//https://leetcode.com/problems/longest-common-prefix/

import (
	"fmt"
	"strings"
)

//AlphabetSize num of letters
const AlphabetSize = 26
//MaxWordLenght
const MaxWordLenght = 250

//TrieNode struct solution
type TrieNode struct {
	children [AlphabetSize]*TrieNode
	childerCount int
	endOfWords bool
}

func getTrieNode() *TrieNode {
	node := &TrieNode{}
	node.endOfWords = false
	for i := 0; i < AlphabetSize; i++ {
		node.children[i] = nil
		node.childerCount = 0
	}
	return node
}

func insertTrie(root *TrieNode, key string) {
	temp := root
	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if(temp.children[index] == nil) {
			temp.children[index] = getTrieNode()
			temp.childerCount ++
		}
		temp = temp.children[index]
	}
	temp.endOfWords = true
}

func createTrie(strs []string, minLenght *int) *TrieNode {
	root := getTrieNode()
	for i:= 0; i < len(strs); i++ {
		if(len(strs[i]) == 0) {
			return nil
		}
		if(len(strs[i]) < *minLenght) {
			*minLenght = len(strs[i])
		}
		insertTrie(root, strs[i])
	}
	return root
}

func findCommonPrefix(root *TrieNode, prefix *string) {
	if root.childerCount == 1 {
		for i := 0; i < AlphabetSize; i++ {
			if(root.children[i] != nil) {
				*prefix += string('a' + i)
				findCommonPrefix(root.children[i], prefix)
				break
			}
		}
	}
}

func longestCommonPrefix(strs []string) string {
	if (strs == nil || len(strs) == 0) {
		return ""
	} else if(len(strs) == 1) { 
		return strs[0]
	}
	minLenght := MaxWordLenght
	trie := createTrie(strs, &minLenght)
	if(trie == nil) {
		return ""
	}
	result := ""
	findCommonPrefix(trie, &result)
	if(len(result) > minLenght) {
		return result[0:minLenght]
	}
	return result
}

func trieLenght(lenght *int, root *TrieNode) {
	for i:= 0; i < len(root.children); i++ {
		if(root.children[i] != nil) {
			*lenght ++
			trieLenght(lenght, root.children[i])
		}
	}
}

func search(root *TrieNode, key string) bool {
	temp := root
	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if(temp.children[index] != nil) {
			temp = temp.children[index]
		} else {
			return false
		}
	}
	return (temp != nil && temp.endOfWords)
} 

func testLongestCommonPrefix() {
	fmt.Println()
	fmt.Println("Longest Common Prefix Trie")
	
	words := []string{"flower","flow","flight"}
	output := "fl"
	result := longestCommonPrefix(words)
	fmt.Printf("Longes preffix. Input: %v; result: %s; expected: %s\n", words, result, output)

	words = []string{"dog","racecar","car"}
	output = ""
	result = longestCommonPrefix(words)
	fmt.Printf("Longes preffix. Input: %v; result: %s; expected: %s\n", words, result, output)

	words = []string{"","b"}
	output = ""
	result = longestCommonPrefix(words)
	fmt.Printf("Longes preffix. Input: %v; result: %s; expected: %s\n", words, result, output)

	words = []string{"aa","a"}
	output = "a"
	result = longestCommonPrefix(words)
	fmt.Printf("Longes preffix. Input: %v; result: %s; expected: %s\n", words, result, output)
}

// Simple sub string solution

func longestCommonPrefixSimple(strs []string) string {
	if(strs == nil || len(strs) == 0) {
		return ""
	}
	prefix := strs[0]
	for i := 0; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 && len(prefix) > 0 {
			prefix = prefix[:len(prefix) - 1]
			if(len(prefix) == 0) {
				return ""
			}
		}
	}
	return prefix
}

func testLongestCommonPrefixSimple() {
	fmt.Println()
	fmt.Println("Longest Common Prefix Simple")
	
	words := []string{"flower","flow","flight"}
	output := "fl"
	result := longestCommonPrefixSimple(words)
	fmt.Printf("Longes preffix. Input: %v; result: %s; expected: %s\n", words, result, output)

	words = []string{"dog","racecar","car"}
	output = ""
	result = longestCommonPrefixSimple(words)
	fmt.Printf("Longes preffix. Input: %v; result: %s; expected: %s\n", words, result, output)

	words = []string{"","b"}
	output = ""
	result = longestCommonPrefixSimple(words)
	fmt.Printf("Longes preffix. Input: %v; result: %s; expected: %s\n", words, result, output)

	words = []string{"aa","a"}
	output = "a"
	result = longestCommonPrefixSimple(words)
	fmt.Printf("Longes preffix. Input: %v; result: %s; expected: %s\n", words, result, output)

	words = []string{"c","acc","ccc"}
	output = ""
	result = longestCommonPrefixSimple(words)
	fmt.Printf("Longes preffix. Input: %v; result: %s; expected: %s\n", words, result, output)
}