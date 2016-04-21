package main

import (
    "fmt"
    "strings"
    "errors"
)

const uppers = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lowers = "abcdefghijklmnopqrstuvwxyz"
const others = "0123456789=●"

type Pair struct {
    Row int
    Col int
}

type PlayFairGrid struct {
    Grid *[][]string
    PositionsCache *map[string]Pair
}
func (self *PlayFairGrid) Build(key string) {
    self.Grid = BuildGrid(BuildGridString(key))
    self.PositionsCache = CacheLetterPositions(self.Grid)
}
func (self *PlayFairGrid) TransformTo(a string, b string) (aa string, bb string, err error) {
    if a == b { return "", "", errors.New("Cannot transform with identical characters " + a) }
    apos, ok := (*self.PositionsCache)[a]
    if !ok { return "", "", errors.New(a + " was not found in the grid") }
    bpos, ok := (*self.PositionsCache)[b]
    if !ok { return "", "", errors.New(b + " was not found in the grid") }

    var aapos, bbpos Pair
    if apos.Row == bpos.Row {
        aapos = Pair{apos.Row, apos.Col + 1}
        if aapos.Col >= 8 {aapos.Col = 0}
        bbpos = Pair{bpos.Row, bpos.Col + 1}
        if bbpos.Col >= 8 {bbpos.Col = 0}
    } else if apos.Col == bpos.Col {
        aapos = Pair{apos.Row + 1, apos.Col}
        if aapos.Row >= 8 {aapos.Row = 0}
        bbpos = Pair{bpos.Row + 1, bpos.Col}
        if bbpos.Row >= 8 {bbpos.Row = 0}
    } else {
        aapos = Pair{apos.Row, bpos.Col}
        bbpos = Pair{bpos.Row, apos.Col}
    }
    aa = (*self.Grid)[aapos.Row][aapos.Col]
    bb = (*self.Grid)[bbpos.Row][bbpos.Col]
    return aa, bb, nil
}
func (self *PlayFairGrid) TransformFrom(a string, b string) (aa string, bb string, err error) {
    if a == b { return "", "", errors.New("Cannot transform with identical characters " + a) }
    apos, ok := (*self.PositionsCache)[a]
    if !ok { return "", "", errors.New(a + " was not found in the grid") }
    bpos, ok := (*self.PositionsCache)[b]
    if !ok { return "", "", errors.New(b + " was not found in the grid") }

    var aapos, bbpos Pair
    if apos.Row == bpos.Row {
        aapos = Pair{apos.Row, apos.Col - 1}
        if aapos.Col < 0 {aapos.Col = 7}
        bbpos = Pair{bpos.Row, bpos.Col - 1}
        if bbpos.Col < 0 {bbpos.Col = 7}
    } else if apos.Col == bpos.Col {
        aapos = Pair{apos.Row - 1, apos.Col}
        if aapos.Row < 0 {aapos.Row = 7}
        bbpos = Pair{bpos.Row - 1, bpos.Col}
        if bbpos.Row < 0 {bbpos.Row = 7}
    } else {
        aapos = Pair{apos.Row, bpos.Col}
        bbpos = Pair{bpos.Row, apos.Col}
    }
    aa = (*self.Grid)[aapos.Row][aapos.Col]
    bb = (*self.Grid)[bbpos.Row][bbpos.Col]
    return aa, bb, nil
}
func (self *PlayFairGrid) TransformStringTo(cipherText string) string {
    plainText := ""
    var last rune
    for i, r := range cipherText {
        if i % 2 == 1 {
            a, b, err := self.TransformTo(string(last), string(r))
            if err != nil {
                panic(err.Error())
            }
            plainText += a + b
        } else {
            last = r
        }
    }
    return plainText
}
func (self *PlayFairGrid) TransformStringFrom(cipherText string) string {
    plainText := ""
    var last rune
    for i, r := range cipherText {
        if i % 2 == 1 {
            a, b, err := self.TransformFrom(string(last), string(r))
            if err != nil {
                panic(err.Error())
            }
            plainText += a + b
        } else {
            last = r
        }
    }
    return plainText
}
func (self *PlayFairGrid) Print() {
    for _, row := range (*self.Grid) {
        fmt.Println(row)
    }
}

func BuildGridString(key string) string {
    uppercaseKey := strings.ToUpper(key)
    lowercaseKey := strings.ToLower(key)
    uSeenMap := make(map[rune]bool)
    gridString := make([]rune, 0)

    for _, b := range uppercaseKey {
        in, _ := uSeenMap[b]
        if !in {
            gridString = append(gridString, b)
            uSeenMap[b] = true
        }
    }
    for _, b := range uppers {
        in, _ := uSeenMap[b]
        if !in {
            gridString = append(gridString, b)
            uSeenMap[b] = true
        }
    }
    for _, b := range lowercaseKey {
        in, _ := uSeenMap[b]
        if !in {
            gridString = append(gridString, b)
            uSeenMap[b] = true
        }
    }
    for _, b := range lowers {
        in, _ := uSeenMap[b]
        if !in {
            gridString = append(gridString, b)
            uSeenMap[b] = true
        }
    }
    for _, b := range others {
        in, _ := uSeenMap[b]
        if !in {
            gridString = append(gridString, b)
            uSeenMap[b] = true
        }
    }
    return string(gridString)
}

func BuildGrid(gridString string) *[][]string {
    output := make([][]string, 0)
    var row []string
    var i int
    for _, r := range gridString {
        if i % 8 == 0 {
            row = make([]string, 8)
        }
        row[i % 8] = string(r)
        if i % 8 == 7 {
            output = append(output, row)
        }
        i++
    }
    return &output
}

func CacheLetterPositions(grid *[][]string) *map[string]Pair {
    output := make(map[string]Pair, 0)
    for i := 0; i < 8; i++ {
        for j := 0; j < 8; j++ {
            c := (*grid)[i][j]
            output[c] = Pair{i, j}
        }
    }
    return &output
}

func main() {

    grid := PlayFairGrid{}
    grid.Build("glassfish")
    grid.Print()

    cipherText := "eFjdlwKgeFlscbApnQEsny3tnye0frxnlrQ5vliW3Yx=5Al●S1nT4obQHql●Ozl●KqeG5252"
    fmt.Println(cipherText)
    fmt.Println(grid.TransformStringFrom(cipherText))
    fmt.Println(grid.TransformStringTo(cipherText))
}
