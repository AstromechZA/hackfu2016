package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "path/filepath"
    "./scorer"
    "container/heap"
    "time"
)

func isAllowed(b int32) bool {
    return b > 96 && b < 123
}
func isAllowedKey(b byte) bool {
    return (b > 32 && b < 127)
}

func calcPossibleKeys(group []byte) []byte {
    output := make([]byte, 0)
    var k byte
    var allowed bool

    for k = 33; k < 127; k++ {
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

func permute(input *[][]byte, startPoint int, f func(*[]byte)) {
    // get number of groups
    target := make([]byte, len(*input))
    permuteInner(&target, startPoint, input, f)
}

func testXorFunc(cipherText *[]byte, scrr *scorer.BytesScorer, sink *TextScoreHeap, si int) func(*[]byte) {

    return func(input *[]byte) {
        // first deref everything
        ciph := *cipherText
        inparr := *input
        // now make output working array
        output := make([]byte, len(ciph))
        for i, ib := range ciph {
            ind := (i - si)
            if ind >= 0 && (ind % 16) < len(inparr) {
                v := int32(ib ^ inparr[ind % 16])
                v = v << 0x18
                v = v >> 0x18
                output[i] = byte(v)
            } else {
                output[i] = byte(' ')
            }
        }

        s := scrr.QuintgramScore(&output)
        keystring := string(*input)
        outstring := string(output)

        o := TextScore{ s, keystring, outstring }
        heap.Push(sink, &o)
    }
}

type TextScore struct {
    Score float64
    Key string
    Text string
}
type TextScoreHeap []*TextScore
func (self TextScoreHeap) Len() int { return len(self) }
func (self TextScoreHeap) Less(i, j int) bool { return self[i].Score > self[j].Score }
func (self TextScoreHeap) Swap(i, j int) { self[i], self[j] = self[j], self[i] }
func (self *TextScoreHeap) Push(x interface{}) {
    *self = append(*self, x.(*TextScore))
}
func (self *TextScoreHeap) Pop() interface{} {
    old := *self
    n := len(old)
    x := old[n - 1]
    *self = old[:n - 1]
    return x
}

func main() {

    startTime := time.Now()
    // first identify this file path
    thisFile, _ := filepath.Abs(os.Args[0])
    inputBytes, _ := ioutil.ReadFile(filepath.Join(thisFile, "..", "messages.enc"))
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

    keygroups := make([][]byte, 0)
    permutations := 1
    for _, g := range groups {
        possibleKeys := calcPossibleKeys(g)
        keygroups = append(keygroups, possibleKeys)
        permutations *= len(possibleKeys)
    }

    scrr := scorer.BytesScorer{}
    scrr.Init(
        filepath.Join(thisFile, "..", "scorer", "bigrams.txt"),
        filepath.Join(thisFile, "..", "scorer", "trigrams.txt"),
        filepath.Join(thisFile, "..", "scorer", "quadgrams.txt"),
        filepath.Join(thisFile, "..", "scorer", "quintgrams.txt"))

    scoreHeap := &TextScoreHeap{}
    heap.Init(scoreHeap)

    si := 10
    ei := 16
    fmt.Println("permuting for characters", si, "to", ei)
    kg := keygroups[si:ei]
    fmt.Println(kg)
    f := testXorFunc(&inputBytes, &scrr, scoreHeap, si)
    permute(&kg, 0, f)

    fmt.Println("top", 200)
    for i := 0; i < 200; i++ {
        item := heap.Pop(scoreHeap).(*TextScore)
        fmt.Printf("%8.4f '%s' '%s'\n", item.Score, item.Text, item.Key)
    }

    endTime := time.Now()
    fmt.Println("elapsed", endTime.Sub(startTime))
}
