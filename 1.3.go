package main

import (
	"encoding/hex"
	"fmt"
	"github.com/Najaf/cryptopals-solutions/util"
	"github.com/Najaf/cryptopals-solutions/charfreq"
)

func main() {
	ciphertext, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	var plaintexts [][]byte

	for i := 0; i < 256; i++ {
		result := util.SingleByteXOR(ciphertext, byte(i))
		plaintexts = append(plaintexts, result)
	}

  winner, _ := charfreq.BestCharFrequencyScore(plaintexts)

	fmt.Printf("%s\n", winner)
}
