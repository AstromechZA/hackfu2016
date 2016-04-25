package main

import (
    "fmt"
)

const CipherText = "eFjdlwKgeFlscbApnQEsny3tnye0frxnlrQ5vliW3Yx=5Al●S1nT4obQHql●Ozl●KqeG5252"
const AlphabetChars = "abcdefghijklmnopqrstuvwxyz"
const OtherChards = "0123456789=●"

func CalcFitness(input *[]rune) float32 {
    total := float32(0)
    for _, r := range (*input) {
        if r > 96 and r < 123 {
            total += 1
        }
    }
    return total
}

func main() {

}
