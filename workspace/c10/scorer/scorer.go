package scorer

import (
    "os"
    "bufio"
    "strconv"
    "math"
)

/*
how to store map of bigram trigrams quadgrams
*/

func importGramFile(file string, n int) *map[int32]float64 {
    staging := make(map[int32]int, 0)
    output := make(map[int32]float64, 0)

    f, err := os.Open(file)
    if err != nil {
        panic(err.Error())
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    total := 0
    for scanner.Scan() {
        line := scanner.Text()
        var v int32
        for i := 0; i < n; i++ {
            c := line[i]
            v = v << 8
            v = v + int32(c)
        }
        count, _ := strconv.ParseInt(line[n+1:], 10, 32)
        staging[v] = int(count)
        total += int(count)
    }
    if err := scanner.Err(); err != nil {
        panic("Scanner hit error " + err.Error())
    }
    output[0] = math.Log(float64(0.01) / float64(total))
    for k, v := range staging {
        output[k] = math.Log(float64(v) / float64(total))
    }
    return &output
}

type BytesScorer struct {
    Bigrams *map[int32]float64
    Trigrams *map[int32]float64
    Quadgrams *map[int32]float64
}

func (self *BytesScorer) Init(bigramFile string, trigramFile string, quadgramFile string) {
    self.Bigrams = importGramFile(bigramFile, 2)
    self.Trigrams = importGramFile(trigramFile, 3)
    self.Quadgrams = importGramFile(quadgramFile, 4)
}

func (self *BytesScorer) BigramScore(input *[]byte) float64 {
    bigramsDict := (*self.Bigrams)
    var value int32
    var score float64
    for i, b := range *input {
        if b > 96 && b < 123 { b -=32 }
        value = value << 8
        value += int32(b)
        value &= 0xffff
        if i >= 1 {
            s, ok := bigramsDict[value]
            if !ok { s, ok = bigramsDict[0] }
            if ok {
                score += s
            }
        }
    }
    return score / float64(len(*input))
}

func (self *BytesScorer) TrigramScore(input *[]byte) float64 {
    bigramsDict := (*self.Trigrams)
    var value int32
    var score float64
    for i, b := range *input {
        if b > 96 && b < 123 { b -=32 }
        value = value << 8
        value += int32(b)
        value &= 0xffffff
        if i >= 2 {
            s, ok := bigramsDict[value]
            if !ok { s, ok = bigramsDict[0] }
            if ok {
                score += s
            }
        }
    }
    return score / float64(len(*input))
}

func (self *BytesScorer) QuadgramScore(input *[]byte) float64 {
    bigramsDict := (*self.Quadgrams)
    var value int32
    var score float64
    for i, b := range *input {
        if b > 96 && b < 123 { b -= 32 }
        value = value << 8
        value += int32(b)
        if i >= 3 {
            s, ok := bigramsDict[value]
            if !ok { s, ok = bigramsDict[0] }
            if ok {
                score += s
            }
        }
    }
    return score / float64(len(*input))
}

func (self *BytesScorer) CompositeScore(input *[]byte) float64 {
    total := float64(0)
    total += self.BigramScore(input)
    total += self.TrigramScore(input)
    total += self.QuadgramScore(input)
    return total
}
