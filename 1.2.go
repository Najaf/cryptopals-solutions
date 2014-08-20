package main

import (
	"encoding/hex"
	"fmt"
	"github.com/Najaf/cryptopals-solutions/util"
)

func main() {
	encoded_a := "1c0111001f010100061a024b53535009181c"
	encoded_b := "686974207468652062756c6c277320657965"

	a, _ := hex.DecodeString(encoded_a)
	b, _ := hex.DecodeString(encoded_b)

	generated_result := hex.EncodeToString(util.FixedXOR(a, b))
	expected_result := "746865206b696420646f6e277420706c6179"

	fmt.Printf("Expected:\t%s\n", expected_result)
	fmt.Printf("Generated:\t%s\n", generated_result)

	if generated_result == expected_result {
		fmt.Println("Generated and expected result match")
	}

}
