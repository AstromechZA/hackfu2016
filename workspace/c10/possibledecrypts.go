package main

import (
    "os"
    "path/filepath"
    "container/heap"
    "io/ioutil"
    "bufio"
    "fmt"
    "./scorer"
)

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

func loadKeyParts(path string) []string {
    output := make([]string, 100)
    f, err := os.Open(path)
    if err != nil {
        panic(err.Error())
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    i := 0
    for scanner.Scan() {
        s := scanner.Text()
        s = s[10:46]
        output[i] = s
        i += 1
        if i >= 100 {
            break
        }
    }
    if err := scanner.Err(); err != nil {
        panic(err.Error())
    }
    return output
}

func main() {
    thisFile, _ := filepath.Abs(os.Args[0])

    scrr := scorer.BytesScorer{}
    scrr.Init(
        filepath.Join(thisFile, "..", "scorer", "bigrams.txt"),
        filepath.Join(thisFile, "..", "scorer", "trigrams.txt"),
        filepath.Join(thisFile, "..", "scorer", "quadgrams.txt"),
        filepath.Join(thisFile, "..", "scorer", "quintgrams.txt"))

    scoreHeap := &TextScoreHeap{}
    heap.Init(scoreHeap)

    inputBytes, _ := ioutil.ReadFile(filepath.Join(thisFile, "..", "messages.enc"))
    inputBytes = inputBytes[:len(inputBytes) - 1]

    fmt.Println("input", inputBytes)

    datas0004 := loadKeyParts(filepath.Join(thisFile, "..", "datas", "0004.txt"))
    datas0410 := loadKeyParts(filepath.Join(thisFile, "..", "datas", "0410.txt"))
    datas1016 := loadKeyParts(filepath.Join(thisFile, "..", "datas", "1016.txt"))

    final := make([]byte, 36)
    for _, s1 := range datas0004 {

        for i := 0; i < 36; i++ {
            j := i % 16
            if j >= 0 && j < 4 {
                final[i] = s1[i]
            }
        }

        for _, s2 := range datas0410 {

            for i := 0; i < 36; i++ {
                j := i % 16
                if j >= 4 && j < 10 {
                    final[i] = s2[i]
                }
            }

            for _, s3 := range datas1016 {

                for i := 0; i < 36; i++ {
                    j := i % 16
                    if j >= 10 && j < 16 {
                        final[i] = s3[i]
                    }
                }

                score := scrr.QuintgramScore(&final)

                heap.Push(scoreHeap, &TextScore{ score, "", string(final) })

            }
        }
    }

    fmt.Println("top", 1000)
    for i := 0; i < 1000; i++ {
        item := heap.Pop(scoreHeap).(*TextScore)
        fmt.Printf("%8.4f '%s' '%s'\n", item.Score, item.Text, item.Key)
    }

}
