package main

import (
    "fmt"
    "os"
    "path/filepath"
    "./scorer"
)

func main() {
    thisFile, _ := filepath.Abs(os.Args[0])
    s := scorer.BytesScorer{}
    s.Init(
        filepath.Join(thisFile, "..", "scorer", "bigrams.txt"),
        filepath.Join(thisFile, "..", "scorer", "trigrams.txt"),
        filepath.Join(thisFile, "..", "scorer", "quadgrams.txt"))

    s1 := []byte("youhavesolvedthefourthchallengeproceedtoedentofacethefifthchallenge")
    s2 := []byte("xxxxxxxxxxxxxxxxxxxxxxadwahwdiuaahwdiuhawiduhaiwuhdaiwuhdiauwhdiauw")

    fmt.Println(s.BigramScore(&s1))
    fmt.Println(s.TrigramScore(&s1))
    fmt.Println(s.QuadgramScore(&s1))
    fmt.Println(s.BigramScore(&s2))
    fmt.Println(s.TrigramScore(&s2))
    fmt.Println(s.QuadgramScore(&s2))
    fmt.Println()
    fmt.Println(s.CompositeScore(&s1))
    fmt.Println(s.CompositeScore(&s2))
}
