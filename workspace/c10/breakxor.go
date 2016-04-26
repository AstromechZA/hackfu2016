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
func isAllowedKey(b byte) bool {
    return (b > 96 && b < 123) || (b > 66 && b < 91) || b == 32
}

func calcPossibleKeys(group []byte) []byte {
    output := make([]byte, 0)
    var k byte
    var allowed bool

    for k = 0; k < 255; k++ {
        // if ! isAllowedKey(k) {
        //     continue
        // }
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
        }
    }
    return output
}

func permuteInner(target *[]byte, targetN int, source *[][]byte, f func(*[]byte)) {
    targetR := (*target)
    lenn := len(targetR)
    if targetN < lenn {
        column := (*source)[targetN]
        for _, v := range column {
            targetR[targetN] = v
            permuteInner(target, targetN + 1, source, f)
        }
    } else {
        f(&targetR)
    }
}

func permute(input *[][]byte, f func(*[]byte)) {
    fmt.Println(*input)
    // get number of groups
    numColumns := make([]byte, len(*input))
    permuteInner(&numColumns, 0, input, f)
}

func testXorFunc(cipherText *[]byte) func(*[]byte) {

    return func(input *[]byte) {
        ciph := *cipherText
        output := make([]byte, len(*input))
        for i, ib := range *input {
            v := int32(ib ^ ciph[i])
            v = v << 0x18
            v = v >> 0x18
            output[i] = byte(v)
        }

        fmt.Println("'", string(*input), "'", string(output), "'")
    }
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

    keygroups := make([][]byte, 0)
    permutations := 1
    for i, g := range groups {
        possibleKeys := calcPossibleKeys(g)
        keygroups = append(keygroups, possibleKeys)
        fmt.Println(i, possibleKeys, len(possibleKeys))
        permutations *= len(possibleKeys)
    }
    fmt.Println("perms", permutations)

    ps := keygroups[:5]
    f := testXorFunc(&inputBytes)
    permute(&ps, f)
}
