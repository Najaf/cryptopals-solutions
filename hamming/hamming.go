package hamming

import (
	"github.com/Najaf/cryptopals-solutions/xor"
)

func Distance(a, b []byte) float32 {
  distance := 0
  xor := xor.Fixed(a, b)

  for _, thisByte := range xor {
    for thisByte > 0 {
      distance += 1
      thisByte &= thisByte - 1
    }
  }

  return float32(distance)
}
