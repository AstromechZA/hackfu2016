package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
    "os"
    "sort"
)

func PrintableByteMetric(b int16, c int) int {
    switch {
    // lowercase
    case 97 <= b && b <= 122:
        return c * 1
    // space
    case b == 32 || b == 10:
        return c * 3
    // uppercase
    case 65 <= b && b <= 90:
        return c * 5
    case b == 33 || b == 44 || b == 46 || b == 34 || b == 39 || b == 63:
        return c * 5
    case 32 <= b && b <= 126:
        return c * 10
    default:
        return c * 1000
    }
}

// PrintableMetric determines a score of how printable the given letter
// frequency is. From best to worst: lowercase, spaces, uppercase,
// punctuation, everything else.
func PrintableMetric(input *map[int16]int) float32 {
    total := 0
    totalCounts := 0
    for k, v := range *input {
        totalCounts += v
        total += PrintableByteMetric(k, v)
    }
    return float32(total) / float32(totalCounts)
}

// LetterFrequency calculates the frequency of characters in the given array
func LetterFrequency(letters *[]int16) *map[int16]int {
    dest := make(map[int16]int)
    for _, b := range *letters {
        i, _ := dest[b]
        dest[b] = i + 1
    }
    return &dest
}

// Map maps a function across all input bytes in an array and returns a new
// array
func Map(input *[]int16, f func(int, int16) int16) *[]int16 {
    output := make([]int16, len(*input))
    for i, v := range *input {
        output[i] = f(i, v)
    }
    return &output
}

func CrackSingleByte(cipherText *[]int16, f func(a int16, b int16) int16) (*[]int16, int16) {

    bestI := int16(0)
    bestScore := float32(1000)
    var i int16
    for i = 0; i < 0x7fff; i++ {

        tfunc := func(p int, b int16) int16 {
            return f(i, b)
        }

        transformedBytes := Map(cipherText, tfunc)
        freq := LetterFrequency(transformedBytes)
        score := PrintableMetric(freq)
        if score < bestScore {
            bestI = i
            bestScore = score
        }
    }

    tfunc := func(p int, b int16) int16 {
        return f(bestI, b)
    }

    return Map(cipherText, tfunc), bestI
}

func Xor(a int16, b int16) int16 {
    return a ^ b
}

func NotXor(a int16, b int16) int16 {
    return ^Xor(a, b)
}

func HammingWeight(a int16) int {
    b := (a & 0x5555) + ((a >> 1) & 0x5555)
    c := (b & 0x3333) + ((b >> 2) & 0x3333)
    d := (c & 0x0F0F) + ((c >> 4) & 0x0F0F)
    return int((d & 0x00FF) + ((d >> 8) & 0x00FF))
}

func HammingDistance(a int16, b int16) int {
    return HammingWeight(Xor(a, b))
}

func NAvgHD(cipherText *[]int16, keyLength int) float32 {

    // split the cipher text into groups of 'keyLength' * 2 ints
    // each group is 2 int16s
    groupLength := keyLength * 2
    iters := len(*cipherText) / groupLength
    totalHD := 0
    for i := 0; i < iters; i++ {
        group1StartIndex := i * groupLength
        group2StartIndex := group1StartIndex + keyLength
        for j := 0; j < keyLength; j++ {
            totalHD += HammingDistance((*cipherText)[group1StartIndex+j], (*cipherText)[group2StartIndex+j])
        }
    }
    // average
    AvgHD := float32(totalHD) / float32(iters)
    // normalize
    return AvgHD / float32(keyLength)
}

// keylengthscore sorting
type KeyLengthScore struct {
    KeyLength int
    Score float32
}
type ByScore []KeyLengthScore
func (s ByScore) Len() int { return len(s) }
func (s ByScore) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByScore) Less(i, j int) bool { return s[i].Score < s[j].Score }

func FindBestKeyLengths(cipherText *[]int16, top int) {
    scores := make([]KeyLengthScore, 0)
    for i := 2; i < top; i++ {
        s := NAvgHD(cipherText, i)
        scores = append(scores, KeyLengthScore{i, s})
    }
    sort.Sort(ByScore(scores))

    fmt.Println("Printing top 10 key lengths:")
    for i := 0; i < 10; i++ {
        fmt.Println(scores[i])
    }
}

func CrackSingleByteInPlace(cipherText *[]int16, plainText *[]int16, keyLength int, keyByte int, f func(a int16, b int16) int16) int16 {

    lenT := len(*cipherText)
    bestPrintableScore := float32(100000000)
    var bestKeyValue int16
    var key int16
    for key = 0; key < 0x7FFF; key++ {

        incidentCounts := make(map[int16]int)
        for i := keyByte; i < lenT; i += keyLength {
            n := f(key, (*cipherText)[i])
            x, _ := incidentCounts[n]
            incidentCounts[n] = x + 1
        }
        printableScore := 0
        printableCounts := 0
        for k, v := range incidentCounts {
            printableCounts += v
            printableScore += PrintableByteMetric(k, v)
        }
        printableScoreAvg := float32(printableScore) / float32(printableCounts)

        if printableScoreAvg < bestPrintableScore {
            bestPrintableScore = printableScoreAvg
            bestKeyValue = key
            fmt.Println("best score for byte", keyByte, "is now", bestPrintableScore, "key=", bestKeyValue)
        }
    }
    for i := keyByte; i < lenT; i += keyLength {
        (*plainText)[i] = f(bestKeyValue, (*cipherText)[i])
    }
    return bestKeyValue
}


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
    shorts := make([]int16, len(content)/2)
    for i := 0; i < len(content); i+=2 {
        shorts[i/2] = (int16(content[i]) << 8) | int16(content[i+1])
    }

    fmt.Println("number of shorts:", len(shorts))

    FindBestKeyLengths(&shorts, 100)

    probablyKeyLength := 13

    fmt.Println("We assume", probablyKeyLength, "is best")

    // build destination for cracked output
    crackedOutput := make([]int16, len(shorts))

    bestKeyBytes := make([]int16, probablyKeyLength)
    for i := 0; i < probablyKeyLength; i++ {
        bestKeyBytes[i] = CrackSingleByteInPlace(&shorts, &crackedOutput, probablyKeyLength, i, NotXor)
    }

    fmt.Println(bestKeyBytes)

    missingChars := make(map[int16]int)

    for _, r := range crackedOutput {
        if r < 128 {
            fmt.Print(string(r))
        } else {
            fmt.Print(string(r - 320))
            i, _ := missingChars[r]
            missingChars[r] = i + 1
        }
    }
}
