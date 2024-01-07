package datastructures

// https://www.geeksforgeeks.org/trie-insert-and-search/

type TrieNode struct {
	Children map[rune]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[rune]*TrieNode),
	}
}

type Trie struct {
	Root *TrieNode
}

func NewTrie(word string) *Trie {
	trie := &Trie{
		Root: NewTrieNode(),
	}
	if word != "" {
		trie.Insert(word)
	}
	return trie
}

func (t *Trie) Insert(word string) {
	current := t.Root
	for _, c := range word {
		if _, ok := current.Children[c]; !ok {
			current.Children[c] = NewTrieNode()
		}
		current = current.Children[c]
	}
	current.Children['*'] = nil
}

func (t *Trie) Search(word string) bool {
	current := t.Root
	for _, c := range word {
		if _, ok := current.Children[c]; !ok {
			return false
		}
		current = current.Children[c]
	}
	return current.Children['*'] != nil
}
