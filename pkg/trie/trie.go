package trie

import (
    "fmt"

    "github.com/BaiXilin/gutenberg-word-prefix-match/pkg/fileio"
)

type TrieNode struct {
    value fileio.Word
    children map[rune]*TrieNode
    isCompleteWord bool
}

func NewTrie() *TrieNode {
    node := new(TrieNode)
    node.children = make(map[rune]*TrieNode)
    node.isCompleteWord = false

    return node
}

func (root *TrieNode) BuildTrie(done <-chan interface{}, wordStream <-chan fileio.Word) *TrieNode {
    // children nodes
    // organized in a map. The keys are words' initial letter
    for i := 'a'; i <= 'z'; i++ {
        root.children[i] = new(TrieNode)
    }
    
    // launch 26 goroutines. Each will handle words start with one letter
    trieBuilders := make(map[rune]chan fileio.Word, 26)
    for i := 'a'; i <= 'z'; i++ {
        trieBuilders[i] = make(chan fileio.Word)
        defer close(trieBuilders[i])
        go root.children[i].trieBuilder(done, trieBuilders[i])
    }

    for w := range wordStream {
            // send the word to its corresponding worker
            fmt.Println([]rune(w.Val)[0])
            trieBuilders[[]rune(w.Val)[0]] <- w
    }
    
    return root
}

func (root *TrieNode) trieBuilder(done <-chan interface{}, wordStream <-chan fileio.Word) {
    for w := range wordStream {
        //select {
        //case <-done:
        //    return
        //}
        fmt.Printf("TrieBuilder %v working on %v\n", w.Val[0], w.Val)
    }
}

// put a word into the trie
func (root *TrieNode) Put(w fileio.Word) error {
    return nil
}
