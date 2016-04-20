package frequency

// LetterFrequency calculates the frequency of characters in the given array
func LetterFrequency(letters *[]byte) *map[byte]int {
    dest := make(map[byte]int)
    for _, b := range *letters {
        i, _ := dest[b]
        dest[b] = i + 1
    }
    return &dest
}

// NGramFrequency calculates the frequency of ngrams in the given byte array
func NGramFrequency(letters *[]byte, n int) (*map[string]int, int) {
    dest := make(map[string]int)
    gramCount := len(*letters) - (n - 1)
    for i := 0; i < gramCount; i++ {
        key := string((*letters)[i:i+n])
        i, _ := dest[key]
        dest[key] = i + 1
    }
    return &dest, gramCount
}

// NormalizedFrequency takes in a frequency map and normalizes each entry
func NormalizedFrequency(input *map[string]int, totalCount int) *map[string]float32 {
    output := make(map[string]float32)
    for k, v := range *input {
        output[k] = float32(v) / float32(totalCount)
    }
    return &output
}

// PrintableMetric determines a score of how printable the given letter
// frequency is. From best to worst: lowercase, spaces, uppercase,
// punctuation, everything else.
func PrintableMetric(input *map[byte]int) float32 {
    total := 0
    totalCounts := 0
    for k, v := range *input {
        totalCounts += v
        switch {
        // lowercase
        case 97 <= k && k <= 122:
            total += v * 1
        // space
        case k == 32 || k == 10:
            total += v * 3
        // uppercase
        case 65 <= k && k <= 90:
            total += v * 5
        case k == 33 || k == 44 || k == 46 || k == 34 || k == 39 || k == 63:
            total += v * 5
        case 32 <= k && k <= 126:
            total += v * 10
        default:
            total += v * 1000
        }
    }
    return float32(total) / float32(totalCounts)
}
