package main

/*

The idea with this script is to randomly swap elements in a playfair cipher and
attempt to evolve the playfair cipher table using a quintgram fitness function.

There are 2 primary difficulties here:
- fast transforms
    we need quick way of swapping elements and swapping their related transforms
    without having to build and manipulate strings and generate new objects
- minimal value copies
    we should be only copying values at specific points

*/

import (
    "fmt"
    "os"
    "math"
    "bufio"
    "strconv"
    "path/filepath"
    "math/rand"
    "time"
)

/*
The playfair square is 8x8=64 items. there are 64*64 transforms to manage here.
we can write them as 4096 entries in a map short->short
and then we just use a 64 entry byte[] as the square and swap items in it as
required -- wait no this doesn't work. stop. rethink. the key is letters, not
positions.

ok then we do our swaps on a map[char(letter)] -> short(position)
by swapping the values of 2 keys, we swap their positions

*/

func BuildOriginalLetterPositions() (*[]byte, *[]byte) {
    letterToPos := make([]byte, 127) // oportunity for optimisation - we could use a sparse array here for true O(1)
    posToLetter := make([]byte, 64)

    alphabetsrc := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789=.")
    alphabetdst := make([]byte, len(alphabetsrc))
    {
        perm := rand.Perm(len(alphabetsrc))
        for i, v := range perm {
            alphabetdst[v] = alphabetsrc[i]
        }
    }

    for i, r := range alphabetdst {
        letterToPos[int(r)] = byte(i)
        posToLetter[i] = byte(r)
    }
    return &letterToPos, &posToLetter
}

func BuildTransformsMap() *[]int16 {
    output := make([]int16, 64 * 64)
    var a, b, c, d int16
    for a = 0; a < 64; a++ {
        for b = 0; b < 64; b++ {
            if a != b {
                // same row
                if (a & 56) == (b & 56) {
                    c = a & 56
                    d = b & 56
                    if a == c { c |= 7 } else { c = a - 1 }
                    if b == d { d |= 7 } else { d = b - 1 }
                // same column
                } else if (a & 7) == (b & 7) {
                    c = a & 56
                    d = b & 56
                    if c == 0 {c = 56} else { c = a - 8 }
                    if d == 0 {d = 56} else { d = b - 8 }
                    c |= (a & 7)
                    d |= (b & 7)
                // different
                } else {
                    c = a & 56
                    d = b & 56
                    c |= (b & 7)
                    d |= (a & 7)
                }
                output[(a << 6) | b] = (c << 6) | d
            } else {
                output[(a << 6) | b] = -1
            }
        }
    }
    return &output
}

func Transform(a byte, b byte, letterToPos *[]byte, posToLetter *[]byte, transforms *[]int16) (byte, byte) {
    lp := *letterToPos
    pl := *posToLetter

    p := (int16(lp[a]) << 6) | int16(lp[b])
    p = (*transforms)[p]

    return pl[ (p >> 6) & 63 ], pl[p & 63]
}

func TransformAll(src *[]byte, dst *[]byte, letterToPos *[]byte, posToLetter *[]byte, transforms *[]int16) {
    lp := *letterToPos
    pl := *posToLetter
    s := *src
    d := *dst

    l := len(s)
    var p int16
    for i := 0; i < l; i+=2 {
        p = (int16(lp[s[i]]) << 6) | int16(lp[s[i+1]])
        p = (*transforms)[p]
        d[i] = pl[ (p >> 6) & 63 ]
        d[i+1] = pl[p & 63]
    }
}

func BuildSNGramScoringTransformAll(ngramsrcfile string, ngram int) func(*[]byte, *[]byte, *[]byte, *[]byte, *[]int16)float64 {
    f, err := os.Open(ngramsrcfile)
    if err != nil {
        panic(err.Error())
    }
    defer f.Close()

    ngramScoreMap := make(map[int64]float64, 0)
    {
        staging := make(map[int64]int, 0)
        scanner := bufio.NewScanner(f)
        total := 0
        for scanner.Scan() {
            line := scanner.Text()
            var v int64
            for i := 0; i < ngram; i++ {
                c := (line[i] + 32) & 0xff
                v = v << 8
                v = v + int64(c)
            }
            count, _ := strconv.ParseInt(line[ngram+1:], 10, 32)
            staging[v] = int(count)
            total += int(count)
        }
        if err := scanner.Err(); err != nil {
            panic("Scanner hit error " + err.Error())
        }
        ngramScoreMap[0] = math.Log(float64(1) / float64(total))
        for k, v := range staging {
            ngramScoreMap[k] = math.Log(float64(v) / float64(total))
        }
    }

    return func(src *[]byte, dst *[]byte, letterToPos *[]byte, posToLetter *[]byte, transforms *[]int16) float64 {
        lp := *letterToPos
        pl := *posToLetter
        s := *src
        d := *dst

        worstNGramScore := float64(len(s) - ngram + 1) * ngramScoreMap[0]
        worstCaseScore := float64(len(s))

        l := len(s)
        var p int16
        var v int64
        var casescore float64
        var ngramscore float64
        for i := 0; i < l; i+=2 {
            p = (int16(lp[s[i]]) << 6) | int16(lp[s[i+1]])
            p = (*transforms)[p]
            d[i] = pl[ (p >> 6) & 63 ]
            d[i+1] = pl[p & 63]

            if d[i] > 96 && d[i] < 123 {
                casescore += 1
            }

            if d[i + 1] > 96 && d[i + 1] < 123 {
                casescore += 1
            }

            if i >= ngram {
                v = ((v & 0xffffffff) << 8) | int64(d[i])
                s, ok := ngramScoreMap[v]
                if !ok { s, ok = ngramScoreMap[0] }
                ngramscore += s
            }
            if i+1 >= ngram {
                v = ((v & 0xffffffff) << 8) | int64(d[i + 1])
                s, ok := ngramScoreMap[v]
                if !ok { s, ok = ngramScoreMap[0] }
                ngramscore += s
            }
        }

        ns := (1 - (ngramscore / worstNGramScore))
        cs := (casescore / worstCaseScore)

        return 0.8 * ns + 0.2 * cs
    }
}

func DoArbitrarySwap(letterToPos *[]byte, posToLetter *[]byte, temperature float64) (*[]byte, *[]byte) {
    lp := *letterToPos
    pl := *posToLetter

    r := rand.Intn(40)

    if r == 1 {
        newlp := make([]byte, len(lp))
        copy(newlp, lp)
        newpl := make([]byte, len(pl))
        copy(newpl, pl)

        p1 := byte(rand.Intn(8))
        p2 := p1
        for p1 == p2 {p2 = byte(rand.Intn(8))}
        p1 <<= 3
        p2 <<= 3
        var i byte
        for i = 0; i < 8; i++ {
            la := pl[p1 + i]
            lb := pl[p2 + i]
            newlp[lb] = p1 + i
            newlp[la] = p2 + i
            newpl[p1 + i] = lb
            newpl[p2 + i] = la
        }
        lp = newlp
        pl = newpl

    } else if r == 2 {
        newlp := make([]byte, len(lp))
        copy(newlp, lp)
        newpl := make([]byte, len(pl))
        copy(newpl, pl)
        p1 := byte(rand.Intn(8))
        p2 := p1
        for p1 == p2 {p2 = byte(rand.Intn(8))}
        var i byte
        for i = 0; i < 64; i += 8 {
            la := pl[p1 + i]
            lb := pl[p2 + i]
            newlp[lb] = p1 + i
            newlp[la] = p2 + i
            newpl[p1 + i] = lb
            newpl[p2 + i] = la
        }
        lp = newlp
        pl = newpl
    } else if r == 3 {
        newlp := make([]byte, len(lp))
        copy(newlp, lp)
        newpl := make([]byte, len(pl))
        copy(newpl, pl)
        var i byte
        for i = 0; i < 64; i ++ {
            p1 := i
            p2 := (p1 & 56) | (7 - (p1 & 7))
            la := pl[p1]
            lb := pl[p2]
            newlp[lb] = p1
            newlp[la] = p2
            newpl[p1] = lb
            newpl[p2] = la
        }
        lp = newlp
        pl = newpl
    } else if r == 4 {
        newlp := make([]byte, len(lp))
        copy(newlp, lp)
        newpl := make([]byte, len(pl))
        copy(newpl, pl)
        var i byte
        for i = 0; i < 64; i ++ {
            p1 := i
            p2 := 63 - i
            la := pl[p1]
            lb := pl[p2]
            newlp[lb] = p1
            newlp[la] = p2
            newpl[p1] = lb
            newpl[p2] = la
        }
        lp = newlp
        pl = newpl
    } else if r == 5 {
        newlp := make([]byte, len(lp))
        copy(newlp, lp)
        newpl := make([]byte, len(pl))
        copy(newpl, pl)
        var i byte
        for i = 0; i < 64; i ++ {
            p1 := i
            p2 := (p1 & 7) | (56 - (p1 & 56))
            la := pl[p1]
            lb := pl[p2]
            newlp[lb] = p1
            newlp[la] = p2
            newpl[p1] = lb
            newpl[p2] = la
        }
        lp = newlp
        pl = newpl
    } else {

        newlp := make([]byte, len(lp))
        copy(newlp, lp)
        newpl := make([]byte, len(pl))
        copy(newpl, pl)

        p1 := byte(rand.Intn(64))
        p2 := p1
        for p1 == p2 {p2 = byte(rand.Intn(64))}

        la := pl[p1]
        lb := pl[p2]
        newlp[lb] = p1
        newlp[la] = p2
        newpl[p1] = lb
        newpl[p2] = la

        lp = newlp
        pl = newpl
    }

    return &lp, &pl
}



func main() {

    rand.Seed(time.Now().UnixNano())

    fmt.Println("loading quintgrams...")
    qf := filepath.Join(os.Args[0], "..", "quintgrams.txt")
    transformAllScoringFunc := BuildSNGramScoringTransformAll(qf, 5)
    fmt.Println("2")

    fmt.Println("building original transforms and positions")
    parentLP, parentPL := BuildOriginalLetterPositions()
    tm := BuildTransformsMap()

    in := []byte("eFjdlwKgeFlscbApnQEsny3tnye0frxnlrQ5vliW3Yx=5Al.S1nT4obQHql.Ozl.KqeG5252")
    out := make([]byte, 72)

    bestSeen := ""
    bestSeenFitness := float64(0)
    bestLP := parentLP
    bestPL := parentPL

    for {
        parentFitness := transformAllScoringFunc(&in, &out, parentLP, parentPL, tm)
        childLP := parentLP
        childPL := parentPL
        childFitness := parentFitness

        // many many rounds
        rounds := 1000000
        //temperature starts at 1 and tends to 0 in even steps
        temperature := float64(1)
        temperatureDrop := temperature * (1 + 3 * rand.Float64()) / float64(rounds)

        for i := rounds; i > 0; i-- {
            childLP, childPL = DoArbitrarySwap(parentLP, parentPL, temperature)
            // the lower the fitness, the better the score
            childFitness = transformAllScoringFunc(&in, &out, childLP, childPL, tm)
            dF := childFitness - parentFitness
            // always swap if dF is positive (if childfitness > parentfitness)
            if dF > 0 {
                parentPL = childPL
                parentLP = childLP
                parentFitness = childFitness

                if parentFitness > bestSeenFitness {
                    bestSeenFitness = parentFitness
                    bestSeen = string(out)
                    bestPL = parentPL
                    bestLP = parentLP
                    fmt.Println(bestSeenFitness, bestSeen)
                }

            } else if dF < 0 {
                p := temperature * temperature * temperature * temperature
                r := rand.Float64()
                if r < p {
                    parentPL = childPL
                    parentLP = childLP
                    parentFitness = childFitness
                }
            }
            if temperature > 0.1 {
                temperature -= temperatureDrop
            }
        }
    }
    fmt.Println(bestSeenFitness, bestSeen)
    fmt.Println(bestPL, bestLP)
}
