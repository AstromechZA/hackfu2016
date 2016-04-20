package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "os"
)

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
    shorts := make([]int, len(content)/2)
    for i := 0; i < len(content); i+=2 {
        shorts[i/2] = (int(content[i]) << 8) | int(content[i+1])
    }





}
