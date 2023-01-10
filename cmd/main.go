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
    


    fmt.Println("done")

    // check if trie is successful
    upburstFreq, upburstExist := root.WordIsIn("upburst")
    fmt.Printf("upburst freq: %v, exist: %v\n", upburstFreq, upburstExist)

    // some word that does not exist
    upbxFreq, upbxExist := root.WordIsIn("upbx")
    fmt.Printf("upbx freq: %v, exist: %v\n", upbxFreq, upbxExist)
}
