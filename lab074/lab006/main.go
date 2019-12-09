package main

import (
	"crypto/aes"
	"fmt"

	"github.com/andreburgaud/crypt2go/ecb"
	"github.com/andreburgaud/crypt2go/padding"
)

func encrypt(pt, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBEncrypter(block)
	padder := padding.NewPkcs7Padding(mode.BlockSize())
	pt, err = padder.Pad(pt) // padd last block of plaintext if block size less than block cipher size
	if err != nil {
		panic(err.Error())
	}
	ct := make([]byte, len(pt))
	mode.CryptBlocks(ct, pt)
	return ct
}

func decrypt(ct, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	mode := ecb.NewECBDecrypter(block)
	pt := make([]byte, len(ct))
	mode.CryptBlocks(pt, ct)
	padder := padding.NewPkcs7Padding(mode.BlockSize())
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
	fmt.Println(string(ct))
	fmt.Printf("Ciphertext: %x\n", ct)

	recovered_pt := decrypt(ct, key)
	fmt.Printf("Recovered plaintext: %s\n", recovered_pt)
}

func example2() {
	pt := []byte("1_admin_1576480302")
	key := []byte("app001_seal_1234")

	ct := encrypt(pt, key)
	fmt.Println(string(ct))
	fmt.Printf("Ciphertext: %x\n", ct)

	recovered_pt := decrypt(ct, key)
	fmt.Printf("Recovered plaintext: %s\n", recovered_pt)
}

func main() {
	//example()
	example2()
}
