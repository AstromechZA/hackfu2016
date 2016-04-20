package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "os"
)

// PrintableMetric determines a score of how printable the given letter
// frequency is. From best to worst: lowercase, spaces, uppercase,
// punctuation, everything else.
func PrintableMetric(input *map[int16]int) float32 {
    total := 0
    totalCounts := 0
    for k, v := range *input {
        totalCounts += v
        switch {
        // lowercase
        case 97 <= k && k <= 122:
            total += v * 1
        // space
        case k == 32 || k == 10:
            total += v * 3
        // uppercase
        case 65 <= k && k <= 90:
            total += v * 5
        case k == 33 || k == 44 || k == 46 || k == 34 || k == 39 || k == 63:
            total += v * 5
        case 32 <= k && k <= 126:
            total += v * 10
        default:
            total += v * 1000
        }
    }
    return float32(total) / float32(totalCounts)
}

// LetterFrequency calculates the frequency of characters in the given array
func LetterFrequency(letters *[]int16) *map[int16]int {
    dest := make(map[int16]int)
    for _, b := range *letters {
        i, _ := dest[b]
        dest[b] = i + 1
    }
    return &dest
}

// Map maps a function across all input bytes in an array and returns a new
// array
func Map(input *[]int16, f func(int, int16) int16) *[]int16 {
    output := make([]int16, len(*input))
    for i, v := range *input {
        output[i] = f(i, v)
    }
    return &output
}

func CrackSingleByte(cipherText *[]int16, f func(a int16, b int16) int16) (*[]int16, int16) {

    bestI := int16(0)
    bestScore := float32(1000)
    var i int16
    for i = 0; i < 0x7fff; i++ {

        tfunc := func(p int, b int16) int16 {
            return f(i, b)
        }

        transformedBytes := Map(cipherText, tfunc)
        freq := LetterFrequency(transformedBytes)
        score := PrintableMetric(freq)
        if score < bestScore {
            bestI = i
            bestScore = score
        }
    }

    tfunc := func(p int, b int16) int16 {
        return f(bestI, b)
    }

    return Map(cipherText, tfunc), bestI
}

func Xor(a int16, b int16) int16 {
    return a ^ b
}

func NotXor(a int16, b int16) int16 {
    return ^Xor(a, b)
}

func main() {

    // find file
    thisFile, _ := filepath.Abs(os.Args[0])
    bytesFile := filepath.Join(thisFile, "..", "bytes")
    fmt.Println("file:", bytesFile)

    // read bytes
    content, err := ioutil.ReadFile(bytesFile)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    fmt.Println("number of bytes:", len(content))

    // assume we must convert to int16
    shorts := make([]int16, len(content)/2)
    for i := 0; i < len(content); i+=2 {
        shorts[i/2] = (int16(content[i]) << 8) | int16(content[i+1])
    }

    fmt.Println("number of shorts:", len(shorts))

    bestResult, _ := CrackSingleByte(&shorts, NotXor)

    for _, r := range *bestResult {
        if r < 128 {
            fmt.Print(string(r))
        }
    }

}
