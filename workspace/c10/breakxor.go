package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "path/filepath"
)

func isAllowed(b int32) bool {
    return b > 96 && b < 123
}

func calcPossibleKeys(group []byte) []byte {
    output := make([]byte, 0)
    var k byte
    var allowed bool
    for k = 0; k < 255; k++ {
        allowed = true
        for _, g := range group {
            v := int32(k ^ g)
            v = v << 0x18
            v = v >> 0x18
            if !isAllowed(v) {
                allowed = false
                break
            }
        }
        if allowed {
            output = append(output, k)
            break
        }
    }
    return output
}




func main() {

    // first identify this file path
    thisFile, _ := filepath.Abs(os.Args[0])
    inputBytes, _ := ioutil.ReadFile(filepath.Join(thisFile, "..", "input"))
    inputBytes = inputBytes[:len(inputBytes) - 1]
    fmt.Println("bytes", inputBytes)

    ciLength := len(inputBytes)
    groups := make([][]byte, 16)
    for i := 0; i < 16; i++ {
        group := make([]byte, 0)
        for j := i; j < ciLength; j+=16 {
            group = append(group, inputBytes[j])
        }
        groups[i] = group
    }
    fmt.Println("all groups", groups)

    permutations := 1
    for i, g := range groups {
        possibleKeys := calcPossibleKeys(g)
        fmt.Println(i, possibleKeys, len(possibleKeys))
        permutations *= len(possibleKeys)
    }
    fmt.Println("perms", permutations)

    //currentKey := make([]byte, 16)

    //perm(&currentKey, 0, &groups)



}
