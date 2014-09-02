package main

import (
	"encoding/hex"
	"fmt"
	"github.com/Najaf/cryptopals-solutions/charfreq"
	"github.com/Najaf/cryptopals-solutions/util"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("challenge-data/4.txt")
	s, _ := ioutil.ReadAll(f)
	ciphertexts := strings.Split(string(s), "\n")

	var plaintexts [][]byte

	for _, encoded_ciphertext := range ciphertexts {
		ciphertext, _ := hex.DecodeString(encoded_ciphertext)
		for i := 0; i < 256; i++ {
			result := util.SingleByteXOR(ciphertext, byte(i))
			plaintexts = append(plaintexts, result)
		}
	}

	winner, _ := charfreq.BestCharFrequencyScore(plaintexts)
	fmt.Printf("%s\n", winner)
}
