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
    inputBytes, _ := ioutil.ReadFile(filepath.Join(thisFile, "..", "input"))
    inputBytes = inputBytes[:len(inputBytes) - 1]

    keyBytes := os.Args[1]

    for i, b := range inputBytes {
        k := keyBytes[i % 16]
        v := int32(b ^ k)
        v = v << 0x18
        v = v >> 0x18
        fmt.Print(string(v))
    }
    fmt.Print("\n\n")
}
