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
	for keySize, averageHammy := range ctext.AverageNormalizedHammingDistances(2, 40) {
		fmt.Printf("%d %f\n", keySize, averageHammy)
	}

}
