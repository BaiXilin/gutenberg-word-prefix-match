package main

import (
    "fmt"

    "github.com/BaiXilin/gutenberg-word-prefix-match/pkg/fileio"
)

func main() {
    done := make(chan interface{})
    wordStream := fileio.ReadWords(done, "data/unsorted_words.txt")
    var w fileio.Word
    counter := 0

    for {
        counter++
        if counter == 10 {
            close(done)
            break
        }
        w = <-wordStream
        fmt.Printf("%+v\n", w)
    }

    fmt.Println("done")
}
