package main

import (
	"encoding/hex"
	"fmt"
	"log"

	"github.com/juanli16/cryptopals/internal/util"
)

const (
	challengeInput1 = "1c0111001f010100061a024b53535009181c"
	challengeInput2 = "686974207468652062756c6c277320657965"
	challengeOutput = "746865206b696420646f6e277420706c6179"
)

func main() {
	a, err := util.HexToBytes(challengeInput1)
	if err != nil {
		log.Fatal(err)
	}

	b, err := util.HexToBytes(challengeInput2)
	if err != nil {
		log.Fatal(err)
	}

	dst := make([]byte, len(a))
	err = util.FixedLenXor(dst, a, b)
	if err != nil {
		log.Fatal(err)
	}

	answer := hex.EncodeToString(dst)
	if answer != challengeOutput {
		log.Fatal("failed set1 challenge 2")
	}

	fmt.Printf("set1/challenge2 answer: %v\n", answer)
}
