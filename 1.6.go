package main

import (
	"fmt"
	"github.com/Najaf/cryptopals-solutions/ciphertext"
	"github.com/Najaf/cryptopals-solutions/xor"
	"io/ioutil"
	"os"
)

func main() {
	// Crack open the ciphertext
	f, _ := os.Open("challenge-data/6.txt")
	encodedCiphertext, _ := ioutil.ReadAll(f)

	ctext := ciphertext.NewCiphertextFromBase64(encodedCiphertext, 0)
	ctext.BlockSize = ctext.GuessRepeatingXorKeySize(2, 40)
	key := ctext.BreakRepeatingXorKey()
	fmt.Println(string(xor.RepeatingKey(ctext.Data, key)))
}
