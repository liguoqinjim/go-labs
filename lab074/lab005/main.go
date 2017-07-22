package main

import (
	"fmt"

	"golang.org/x/crypto/blowfish"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
)

func encrypt(pt, key []byte) []byte {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Pad(pt) // padd last block of plaintext if block size less than block cipher size
	if err != nil {
		panic(err.Error())
	}
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)
	return ct
}

func decrypt(ct, key []byte) []byte {
	block, err := blowfish.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBDecrypter(block)
	pt := make([]byte, len(ct))
	mode.CryptBlocks(pt, ct)
	padder := padding.NewPkcs5Padding()
	pt, err = padder.Unpad(pt) // unpad plaintext after decryption
	if err != nil {
		panic(err.Error())
	}
	return pt
}

func example() {
	pt := []byte("Some plain text")
	key := []byte("a_very_secret_key")

	ct := encrypt(pt, key)
	fmt.Printf("Ciphertext: %x\n", ct)

	recovered_pt := decrypt(ct, key)
	fmt.Printf("Recovered plaintext: %s\n", recovered_pt)
}

func main() {
	example()
}
