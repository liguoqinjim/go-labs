package main

import "fmt"
import "crypto/aes"

func main() {
	bc, err := aes.NewCipher([]byte("key3456789012345"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The block size is %d\n", bc.BlockSize())

	var dst = make([]byte, 16)
	var src = []byte("sensitive1234567")

	bc.Encrypt(dst, src)
	fmt.Println(dst)
}
