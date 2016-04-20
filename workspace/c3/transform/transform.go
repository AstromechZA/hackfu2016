package transform

// Map maps a function across all input bytes in an array and returns a new
// array
func Map(input *[]byte, f func(int, byte) byte) *[]byte {
    output := make([]byte, len(*input))
    for i, v := range *input {
        output[i] = f(i, v)
    }
    return &output
}


func XorByteTransform(input *[]byte, x byte) *[]byte {
    inner := func(p int, b byte) byte {
        return b ^ x
    }
    return Map(input, inner)
}

func XorBytesTransform(input *[]byte, x *[]byte) *[]byte {
    xl := len(*x)
    inner := func(i int, b byte) byte {
        return b ^ (*x)[i % xl]
    }
    return Map(input, inner)
}
