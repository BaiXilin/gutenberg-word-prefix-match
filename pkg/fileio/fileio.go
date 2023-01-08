package fileio

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Word struct {
    Word string
    Freq float32
}

func ReadWords(done <-chan interface{}, path string) <-chan Word {
    wordStream := make(chan Word)
    go func() {
        defer close(wordStream)
        file, err := os.Open(path)
        if err != nil {
            fmt.Printf("error opening file: %v\n", err)
            return
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            w, f := splitLine(scanner.Text())
            wordStream <- Word{
                Word: w,
                Freq: f,
            }
        }
    }()
    return wordStream
}

func splitLine(l string) (string, float32) {
    res := strings.Split(l, "\t")
    return res[0], strToFloat32(res[1])
}

func strToFloat32(str string) float32 {
    f, _ := strconv.ParseFloat(str, 64)
    return float32(f)
}
