package main

import (
	"github.com/tink-crypto/tink-go/v2/aead"
	"github.com/tink-crypto/tink-go/v2/keyset"
	"encoding/base64"
	"fmt"
)

func main() {
	text := "hello message"

	tpl := aead.AES256GCMKeyTemplate()
	fmt.Println(tpl)

	kh, err := keyset.NewHandle(tpl)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(kh)

	a, err := aead.New(kh)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(a)

	ciphertext, err := a.Encrypt([]byte(text), nil)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Cipher:", ciphertext)
	fmt.Println(base64.StdEncoding.EncodeToString(ciphertext))

	plaintext, err := a.Decrypt(ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Plaintext:", string(plaintext), plaintext)
}