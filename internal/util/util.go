package util

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
)

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
