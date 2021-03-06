package ciphertext

import (
	"encoding/base64"
	"github.com/Najaf/cryptopals-solutions/charfreq"
	"github.com/Najaf/cryptopals-solutions/hamming"
	"github.com/Najaf/cryptopals-solutions/xor"
)

type Ciphertext struct {
	Data      []byte
	BlockSize int
}

func NewCiphertextFromBase64(data []byte, blockSize int) *Ciphertext {
	decoded, _ := base64.StdEncoding.DecodeString(string(data))
	return NewCiphertext(decoded, blockSize)
}

func NewCiphertext(data []byte, blockSize int) *Ciphertext {
	return &Ciphertext{data, blockSize}
}

func (c Ciphertext) Block(n int) []byte {
	startIndex := n * c.BlockSize
	endIndex := startIndex + c.BlockSize
	return c.Data[startIndex:endIndex]
}

func (c Ciphertext) Blocks(blocks ...int) [][]byte {
	result := make([][]byte, len(blocks))

	for i, n := range blocks {
		result[i] = c.Block(n)
	}

	return result
}

func (c *Ciphertext) BlockCount() int {
	return len(c.Data) / c.BlockSize
}

func (c *Ciphertext) TransposedBlock(byteIndex int) []byte {
	result := make([]byte, c.BlockCount())
	for n := 0; n < c.BlockCount(); n++ {
		result[n] = c.Block(n)[byteIndex]
	}
	return result
}

func (c *Ciphertext) KeyByte(byteIndex int) byte {
	transposedBlock := c.TransposedBlock(byteIndex)
	maxScore, keyByte := 0, byte(0)
	for i := 0; i < 256; i++ {
		plaintext := xor.SingleByte(transposedBlock, byte(i))
		score := charfreq.CharacterFrequncyScore(plaintext)
		if score > maxScore {
			maxScore = score
			keyByte = byte(i)
		}
	}
	return keyByte
}

func (c *Ciphertext) BreakRepeatingXorKey() []byte {
	key := make([]byte, c.BlockSize)
	for i := 0; i < c.BlockSize; i++ {
		key[i] = c.KeyByte(i)
	}
	return key
}

func (c *Ciphertext) averageNormalizedHammingDistances(lower, upper int) map[int]float32 {
	result := make(map[int]float32)
	originalBlockSize := c.BlockSize
	for keySize := lower; keySize <= upper; keySize++ {
		c.BlockSize = keySize
		sampleBlocks := c.Blocks(0, 1, 2, 3)

		firstHammy := hamming.Distance(sampleBlocks[0], sampleBlocks[1]) / float32(keySize)
		secondHammy := hamming.Distance(sampleBlocks[0], sampleBlocks[2]) / float32(keySize)
		thirdHammy := hamming.Distance(sampleBlocks[0], sampleBlocks[3]) / float32(keySize)
		fourthHammy := hamming.Distance(sampleBlocks[1], sampleBlocks[2]) / float32(keySize)
		fifthHammy := hamming.Distance(sampleBlocks[1], sampleBlocks[3]) / float32(keySize)
		sixthHammy := hamming.Distance(sampleBlocks[2], sampleBlocks[3]) / float32(keySize)

		averageHammy := (firstHammy + secondHammy + thirdHammy + fourthHammy + fifthHammy + sixthHammy) / 6.0
		result[keySize] = averageHammy
	}

	c.BlockSize = originalBlockSize
	return result
}

func (c *Ciphertext) GuessRepeatingXorKeySize(lower, upper int) int {
	minHammy := float32(100.0)
	guessedKeySize := 0
	for keySize, hammy := range c.averageNormalizedHammingDistances(lower, upper) {
		if hammy < minHammy {
			minHammy = hammy
			guessedKeySize = keySize
		}
	}
	return guessedKeySize
}
