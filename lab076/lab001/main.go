// Go provides built-in support for [base64
// encoding/decoding](http://en.wikipedia.org/wiki/Base64).

package main

// This syntax imports the `encoding/base64` package with
// the `b64` name instead of the default `base64`. It'll
// save us some space below.
import b64 "encoding/base64"
import "fmt"

func main() {
	data := "abc123!?$*&()'-=@~"
	fmt.Println("原始string:", data)

	// Go supports both standard and URL-compatible
	// base64. Here's how to encode using the standard
	// encoder. The encoder requires a `[]byte` so we
	// cast our `string` to that type.
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("string->base64(standard):", sEnc)

	// Decoding may return an error, which you can check
	// if you don't already know the input to be
	// well-formed.
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("base64->string(standard):", string(sDec))
	fmt.Println()

	// This encodes/decodes using a URL-compatible base64
	// format.
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("string->base64(url):", uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println("base64->string(url):", string(uDec))
}
