package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "path/filepath"
)

func main() {
    // arg thing
    if len(os.Args) != 2 || len(os.Args[1]) != 16 {
        fmt.Println("Expected a single 16 byte key as arg")
        os.Exit(1)
    }

    // first identify this file path
    thisFile, _ := filepath.Abs(os.Args[0])
    inputBytes, _ := ioutil.ReadFile(filepath.Join(thisFile, "..", "messages.enc"))
    inputBytes = inputBytes[:len(inputBytes) - 1]

    keyBytes := os.Args[1]

    output := make([]byte, len(inputBytes))

    for i, b := range inputBytes {
        k := keyBytes[i % 16]
        v := int32(b ^ k)
        v = v << 0x18
        v = v >> 0x18
        output[i] = byte(v)
    }
    fmt.Println("in", inputBytes)
    fmt.Println("out", output)
    fmt.Print(string(output))
    fmt.Print("\n\n")
}
