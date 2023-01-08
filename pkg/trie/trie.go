package trie

type TrieNode struct {
    value Word
    children []*TrieNode
}

func NewTrie() *TrieNode {
    return new(TrieNode)
}

// put a word into the trie
func (root *TrieNode) Put(w Word) error {
    return nil
}
