package util

import (
	"encoding/base64"
	"encoding/hex"
)

func HexToBase64(hexString string) (string, error) {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hexBytes), nil
}
