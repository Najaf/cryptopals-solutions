package hamming

import (
	"github.com/Najaf/cryptopals-solutions/xor"
)

func Distance(a, b string) int {
  distance := 0
  xor := xor.Fixed([]byte(a), []byte(b))

  for _, thisByte := range xor {
    for thisByte > 0 {
      distance += 1
      thisByte &= thisByte - 1
    }
  }

  return distance
}
