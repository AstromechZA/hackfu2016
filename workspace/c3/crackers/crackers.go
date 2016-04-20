package crackers

import (
    "sort"
    "../transform"
    "../frequency"
)

func CrackSingleByte(cipherText *[]byte, f func(a byte, b byte) byte) (*[]byte, byte) {

    bestI := byte(0)
    bestScore := float32(1000)
    var i byte
    for i = 0; i < 255; i++ {

        tfunc := func(p int, b byte) byte {
            return f(i, b)
        }

        transformedBytes := transform.Map(cipherText, tfunc)
        freq := frequency.LetterFrequency(transformedBytes)
        score := frequency.PrintableMetric(freq)
        if score < bestScore {
            bestI = i
            bestScore = score
        }
    }

    tfunc := func(p int, b byte) byte {
        return f(bestI, b)
    }

    return transform.Map(cipherText, tfunc), bestI
}

func HammingWeight(b byte) int {
    ib := int(b)
    total := 0
    total += (ib & 1)
    total += (ib & 2) >> 1
    total += (ib & 4) >> 2
    total += (ib & 8) >> 3
    total += (ib & 16) >> 4
    total += (ib & 32) >> 5
    total += (ib & 64) >> 6
    total += (ib & 128) >> 7
    return total
}

func GetKeySizeScore(cipherText *[]byte, length int, tf func(a byte, b byte) byte) float32 {
    total := 0
    iters := len(*cipherText) - length
    for i := 0; i < iters; i++ {
        a := (*cipherText)[i]
        b := (*cipherText)[i + length]
        total += HammingWeight(tf(a, b))
    }
    return float32(total) / float32(iters)
}

type KeySizeScore struct {
    KeySize int
    Score float32
}

type ByScore []KeySizeScore
func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

func GetBestKeySizes(cipherText *[]byte, bottom int, top int, tf func(a byte, b byte) byte) *[]KeySizeScore {
    scores := make([]KeySizeScore, top - bottom + 1)
    for i := bottom; i <= top; i++ {
        scores[i - bottom] = KeySizeScore{i, GetKeySizeScore(cipherText, i, tf)}
    }

    sort.Sort(ByScore(scores))

    return &scores
}


func CrackMultiByte(cipherText *[]byte, length int, tf func(a byte, b byte) byte) (*[]byte, *[]byte) {

    bestKey := make([]byte, length)
    dataLength := len(*cipherText) / length

    for i := 0; i < length; i++ {

        data := make([]byte, dataLength)
        for j := 0; j < dataLength; j++ {
            data[j] = (*cipherText)[i + j * length]
        }

        _, bk := CrackSingleByte(&data, tf)
        bestKey[i] = bk
    }

    tfunc := func(p int, b byte) byte {
        return tf(bestKey[p % length], b)
    }

    return transform.Map(cipherText, tfunc), &bestKey

}
