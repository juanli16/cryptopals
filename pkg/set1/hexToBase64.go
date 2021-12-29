package main

import (
	"fmt"
	"log"

	"github.com/juanli16/cryptopals/internal/util"
)

const (
	challengeInput  = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	challengeOutput = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
)

func main() {
	answer, err := util.HexToBase64(challengeInput)
	if err != nil {
		log.Fatalf("failed to decode hex string: %v", err)
	}
	if answer != challengeOutput {
		log.Fatal("failed to encode hex to base64")
	}

	fmt.Printf("answer: %v\n", answer)
}
