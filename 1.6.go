package main

import (
	"fmt"
	"github.com/Najaf/cryptopals-solutions/hamming"
)


func main() {
  a, b := "this is a test", "wokka wokka!!!"
  fmt.Printf("Left:\t%s\n", a)
  fmt.Printf("Right:\t%s\t\n", b)
  fmt.Printf("Hamming Distance: %d\n", hamming.Distance(a, b))
}
