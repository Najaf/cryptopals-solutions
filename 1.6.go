package main

import (
	"fmt"
	"github.com/Najaf/cryptopals-solutions/ciphertext"
	"io/ioutil"
	"os"
)

func main() {
	// Crack open the ciphertext
	f, _ := os.Open("challenge-data/6.txt")
	encodedCiphertext, _ := ioutil.ReadAll(f)

	ctext := ciphertext.NewCiphertextFromBase64(encodedCiphertext, 0)
	keySize := ctext.GuessRepeatingXorKeySize(2, 40)
}
