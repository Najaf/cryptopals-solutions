package main

import (
	"github.com/Najaf/cryptopals-solutions/util"
  "encoding/hex"
  "fmt"
)

func singleByteXOR(input []byte, key byte) []byte {
  key_bytes := make([]byte, len(input))
  for i, _ := range key_bytes {
    key_bytes[i] = key
  }

  return util.FixedXOR(input, key_bytes)
}

func main() {
  ciphertext, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

  for i := 0; i < 256; i++ {
    result := singleByteXOR(ciphertext, byte(i))
    fmt.Printf("%s \n", result)
  }
}
