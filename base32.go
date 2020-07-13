package base32crockford

import (
	"fmt"
	"math/big"
)

func Decode(bs []byte) ([]byte, error) {
	_, err := bytesToInt(bs)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func bytesToInt(bs []byte) (*big.Int, error) {
	bitsLen := len(bs) * 8
	if bitsLen%5 != 0 {
		return nil, fmt.Errorf("invalid bit length %d", bitsLen)
	}

	iv := big.NewInt(0)
	for _, b := range bs {
		iv.Lsh(iv, 8)
		iv.Add(iv, big.NewInt(int64(b)))
	}
	return iv, nil
}
