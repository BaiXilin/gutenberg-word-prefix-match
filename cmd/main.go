package main

import (
    "fmt"

    "github.com/BaiXilin/gutenberg-word-prefix-match/pkg/fileio"
    "github.com/BaiXilin/gutenberg-word-prefix-match/pkg/trie"
)

func main() {
    done := make(chan interface{})
    defer close(done)

    // source of data: https://en.wiktionary.org/wiki/Wiktionary:Frequency_lists#English
    wordStream := fileio.ReadWords(done, "data/unsorted_words.txt")
    
    root := trie.NewTrie()
    root.BuildTrie(done, wordStream)
    //for w := range wordStream {
    //    fmt.Printf("%+v\n", w)
    //}

    fmt.Println("done")
}
