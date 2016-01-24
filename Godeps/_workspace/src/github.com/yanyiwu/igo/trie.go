package igo

import (
	"errors"
	"fmt"
)

type Trie struct {
	root *TrieNode
	size int
}

type TrieNode struct {
	b    byte
	next [256]*TrieNode
	data interface{}
}

type TrieMatchItem struct {
	Offset  int
	Pattern string
	Data    interface{}
}

func (tmi *TrieMatchItem) String() string {
	return fmt.Sprintf("%v %v %v", tmi.Offset, tmi.Pattern, tmi.Data)
}

func (t *Trie) Size() int {
	return t.size
}

func (t *Trie) Insert(word string, data ...interface{}) error {
	node := t.root
	if len(data) >= 2 {
		return errors.New("args illegal")
	}
	for _, b := range []byte(word) {
		if node.next[b] == nil {
			newNode := new(TrieNode)
			newNode.b = b
			newNode.data = nil
			node.next[b] = newNode
		}
		node = node.next[b]
	}
	if len(data) == 1 {
		node.data = data[0]
	} else {
		node.data = true
	}
	t.size++
	return nil
}

func (t *Trie) Find(text string) []*TrieMatchItem {
	results := make([]*TrieMatchItem, 0)
	for i := 0; i < len(text); i++ {
		j := 0
		node := t.root.next[text[i+j]]
		for node != nil {
			if node.data != nil {
				results = append(results, &TrieMatchItem{
					Offset:  i,
					Pattern: text[i : i+j+1],
					Data:    node.data})
			}
			j++
			if i+j >= len(text) {
				break
			}
			node = node.next[text[i+j]]
		}
	}
	return results
}

func NewTrie() *Trie {
	trie := new(Trie)
	trie.root = new(TrieNode)
	return trie
}
