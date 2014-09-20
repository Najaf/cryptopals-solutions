package charfreq

import (
	"sort"
)

// Map from ascii bytes => character frequency score
// Pinched from: https://en.wikipedia.org/wiki/Letter_frequency#Relative_frequencies_of_letters_in_the_English_language
var charScores = map[byte]int{
	0x61: 8,  // a
	0x62: 2,  // b
	0x63: 3,  // c
	0x64: 4,  // d
	0x65: 12, // e
	0x66: 2,  // f
	0x67: 2,  // g
	0x68: 6,  // h
	0x69: 7,  // i
	0x6a: 0,  // j
	0x6b: 1,  // k
	0x6c: 4,  // l
	0x6d: 2,  // m
	0x6e: 7,  // n
	0x6f: 8,  // o
	0x70: 2,  // p
	0x71: 0,  // q
	0x72: 6,  // r
	0x73: 6,  // s
	0x74: 9,  // t
	0x75: 3,  // u
	0x76: 1,  // v
	0x77: 2,  // w
	0x78: 0,  // x
	0x79: 2,  // y
	0x7a: 0,  // z
	0x20: 5,  // space
}

func CharacterFrequncyScore(input []byte) int {
	score := 0
	for _, thisByte := range input {
		score = score + charScores[thisByte]
	}
	return score
}

type ByCharacterFrequency [][]byte

func (s ByCharacterFrequency) Len() int {
	return len(s)
}

func (s ByCharacterFrequency) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByCharacterFrequency) Less(i, j int) bool {
	return CharacterFrequncyScore(s[i]) > CharacterFrequncyScore(s[j])
}

func BestCharFrequencyScore(input [][]byte) ([]byte, int) {
	plaintexts := make([][]byte, len(input))
	copy(plaintexts, input)
	sort.Sort(ByCharacterFrequency(plaintexts))

	winner := plaintexts[0]
	return winner, CharacterFrequncyScore(winner)
}
