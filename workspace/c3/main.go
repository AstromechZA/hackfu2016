package main

import (
    "fmt"
    "encoding/hex"
    "./crackers"
)


func main() {

    contentBytes, _ := hex.DecodeString("130e5a0d1b095a1b5a1e1b08115a1b141e5a090e150817035a14131d120e")

    xorFunc := func(a byte, b byte) byte {
        return a ^ b
    }

    bestOutput, bestValue := crackers.CrackSingleByte(&contentBytes, xorFunc)
    fmt.Println("best byte value:", bestValue)
    fmt.Println("best output:", string(*bestOutput))

    bestOutput2, bestValue2 := crackers.CrackMultiByte(&contentBytes, 1, xorFunc)
    fmt.Println("best byte value:", bestValue2)
    fmt.Println("best output:", string(*bestOutput2))

    contentBytes, _ = hex.DecodeString("9887702584b28e6c71b7bb997e7195bf817a3884bf98353889fa9f7d34c7bc8a7625c7ae837425cbfa9e7b258eb6cb73308ea8876c7195bf88703f93b69239718eaecb623094fa9b673e85bb897928c798997c2586b3853222c7b88e6625c7b18e6525c7a98e762382aec535058fb398353894fa89703286af98707188bccb613982fa98703295bf886c7194af99673e92b48f7c3f80fa8a793dc7ae83707186b99f7c278eae827022c7b98a67238ebf8f353e89fa83702382fa8f60238eb48c350688a8877171b0bb99350590b5cb623094fa84737191b39f743dc7b386653e95ae8a7b3282fa9f7a7188af99353f86ae827a3f86b6cb663484af997c259efa8a7b35c7af8761388abb9f707191b388613e95a3c5")

    bestScores := crackers.GetBestKeySizes(&contentBytes, 1, 20, xorFunc)

    for _, s := range *bestScores {
        fmt.Println("checking key length", s.KeySize)
        bestOutput2, bestValue2 = crackers.CrackMultiByte(&contentBytes, s.KeySize, xorFunc)
        fmt.Println("best output:", string(*bestOutput2))
    }


}
