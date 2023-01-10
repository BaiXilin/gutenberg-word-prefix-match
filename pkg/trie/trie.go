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
        //root.children[i] = new(TrieNode)
        root.children[i] = NewTrie()
    }
    
    // launch 26 goroutines. Each will handle words start with one letter
    trieBuilders := make(map[rune]chan fileio.Word, 26)
    for i := 'a'; i <= 'z'; i++ {
        trieBuilders[i] = make(chan fileio.Word)
        defer close(trieBuilders[i])
        go root.children[i].trieBuilder(trieBuilders[i])
    }

    for {
        select {
            case <-done:
                return root
            case w, ok := <-wordStream:
                if !ok {
                    return root
                }
                // send the word to its corresponding worker
                trieBuilders[[]rune(w.Val)[0]] <- w
        }
    }
    
    return root
}

func (root *TrieNode) trieBuilder(wordStream <-chan fileio.Word) {
    for w := range wordStream {
        fmt.Printf("TrieBuilder %v working on %v\n", w.Val[0], w.Val)
        // first character determines which worker to handle the word
        // get rid of it now
        runes := []rune(w.Val)[1:]
        err := root.Put(runes, w)
        if err != nil {
            fmt.Printf("trieBuilder err:%v\n", err)
        }
    }
}

// put a word into the trie
func (root *TrieNode) Put(runes []rune, w fileio.Word) error {
    // recursion base case
    if len(runes) == 0 && !root.isCompleteWord{
        root.isCompleteWord = true
        root.value = w
        return nil
    } else if len(runes) == 0 && root.isCompleteWord {
        return fmt.Errorf("Duplicate words")
    }
    // the runes is not empty, need to go down the recursion
    // check if the next child already exists
    _, exist := root.children[runes[0]]
    if !exist {
        // if it doesn't exist, initialize first
        root.children[runes[0]] = NewTrie()

    }
    // recursively go to the child
    first, runes := runes[0], runes[1:]
    root.children[first].Put(runes, w)
    
    return nil
}

// identify if a word is in the trie, return its associated freq if found
func (root *TrieNode) WordIsIn(target string) (float32, bool) {
    // recursion base case
    if len(target) == 0 && root.isCompleteWord {
        return root.value.Freq, true
    } else if len(target) == 0 && !root.isCompleteWord {
        return float32(0), false
    }

    // if the target word is not empty, try to move down the recursion level
    first := []rune(target)[0]
    _, exist := root.children[first]
    if !exist {
        // no more level to go down
        return float32(0), false
    }
    // if children exists, go down to that children
    target = string([]rune(target)[1:])
    return root.children[first].WordIsIn(target)
}
