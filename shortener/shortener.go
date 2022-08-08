package shortener

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"

	"github.com/itchyny/base58-go"
)

func ShortenUrl(longUrl string) string {
	b := createSha256(longUrl)
	num := new(big.Int).SetBytes(b).Uint64()
	hash := base58Encode([]byte(fmt.Sprintf("%d", num)))

	return hash[:8]
}

func createSha256(input string) []byte {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hash.Sum(nil)
}

func base58Encode(bytes []byte) string {
	enc := base58.BitcoinEncoding
	encoded, err := enc.Encode(bytes)
	if err != nil {
		log.Fatalln("err encoding: ", err)
	}

	return string(encoded)
}
