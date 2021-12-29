package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/juanli16/cryptopals/internal/util"
)

const challengeInput = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

func main() {
	input, err := util.HexToBytes(challengeInput)
	if err != nil {
		log.Fatal(err)
	}

	// ascii characters are 7 bits
	chars := make([]byte, 1<<7)
	for i := range chars {
		chars[i] = byte(rune(i))
	}

	// range over all possible ascii characters to decode the xor cipher
	// and compute the resulting letter frequencies
	// compute the chi squared score with the English letter frequencies,
	// the decoded string that has the least
	// chi squared score is the most likely answer.
	minChi2 := float64(100)
	mostLikelyAnswer := ""
	decoded := make([]byte, len(input))
	for _, c := range chars {
		key := bytes.Repeat([]byte{c}, len(input))
		util.FixedLenXor(decoded, input, key)
		decodedString := string(decoded)
		score := util.Score(util.CountLetterFreq(decodedString))
		if score < minChi2 {
			minChi2 = score
			mostLikelyAnswer = decodedString
		}
	}

	fmt.Printf("most likely answer: %v\n", mostLikelyAnswer)
}
