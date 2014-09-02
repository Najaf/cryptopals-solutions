package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "strings"
)

func main() {
  f, _ := os.Open("challenge-data/4.txt")
  s, _ := ioutil.ReadAll(f)
  ciphertexts := strings.Split(string(s), "\n")

  fmt.Printf("%s\n", ciphertexts)
}
