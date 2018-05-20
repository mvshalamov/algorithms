package main

import (
	"fmt"
)

type ITrie interface {
	Add_word(word string)
	Find(word string) (bool, int)
}

type nodeTrie struct {
	char     string
	children []*nodeTrie
}

type Trie struct {
	root nodeTrie
}

func (t *Trie) Add_word(word string) {
	node := &t.root
	var foundInChildren bool
	for _, ch := range word {
		foundInChildren = false
		for _, child := range node.children {
			if child.char == string(ch) {
				foundInChildren = true
				node = child
				break
			}
		}
		if !foundInChildren {
			trie := nodeTrie{char: string(ch)}
			node.children = append(node.children, &trie)
			node = &trie
		}
	}
}

func (t *Trie) Find(word string) (bool, int) {
	node := t.root
	counter := 0
	var foundInChildren bool
	for _, ch := range word {
		foundInChildren = false
		for _, child := range node.children {
			if child.char == string(ch) {
				counter++
				foundInChildren = true
				node = *child

			}
		}
		if !foundInChildren {
			return false, counter
		}
	}
	return true, counter
}

func testTrie(trie ITrie) (bool, int) {
	trie.Add_word("frog")
	trie.Add_word("grod")
	return trie.Find("frof")
}

func main() {
	fmt.Println(testTrie(&Trie{}))
}
