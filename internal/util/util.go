package util

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"strings"
)

const englishLetters = "abcdefghijklmnopqrstuvwxyz"

var ErrNotSameLength = errors.New("input byte slices do not have the same length")

func HexToBytes(hexString string) ([]byte, error) {
	return hex.DecodeString(hexString)
}

func HexToBase64(hexString string) (string, error) {
	hexBytes, err := HexToBytes(hexString)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hexBytes), nil
}

func FixedLenXor(dst, a, b []byte) error {
	if len(dst) != len(a) || len(a) != len(b) {
		return ErrNotSameLength
	}

	for i := range dst {
		dst[i] = a[i] ^ b[i]
	}

	return nil
}

func CountLetterFreq(s string) map[string]int {
	s = strings.ToLower(s)
	m := make(map[string]int)
	for _, r := range englishLetters {
		l := string(r)
		m[l] = strings.Count(s, l)
	}

	return m
}

func Score(m map[string]int) float64 {
	// from https://www3.nd.edu/~busiforc/handouts/cryptography/letterfrequencies.html
	engFreq := map[string]float64{
		"a": 8.5, "b": 2.07, "c": 4.53, "d": 3.38, "e": 11.16, "f": 1.81, "g": 2.47, "h": 3.0,
		"i": 7.54, "j": 0.19, "k": 1.1, "l": 5.49, "m": 3.01, "n": 6.65, "o": 7.16, "p": 3.167,
		"q": 0.19, "r": 7.58, "s": 5.74, "t": 6.95, "u": 3.63, "v": 1.0, "w": 1.29, "x": 0.29,
		"y": 1.78, "z": 0.27,
	}

	// total lettters
	total := 0
	for _, v := range m {
		total += v
	}

	// compute likelihood score with chi squared
	// the smaller the chi squared score, the more likely it is the same distribution
	chi2 := float64(0)
	for _, r := range englishLetters {
		l := string(r)
		observed := float64(m[l])
		expected := float64(total) * engFreq[l] / 100
		difference := observed - expected
		chi2 += difference * difference / expected
	}
	return chi2
}
