package main

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/cipher"
	"encoding/hex"
	"encoding/base64"
	"bytes"
	"fmt"
)

func main() {
	text := "test message"
	keyString := "645E739A7F9F162725C1533DC2C5E827"

	plain := []byte(text)

	key, err := hex.DecodeString(keyString)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Key:", keyString)

	iv, encrypted, err := Encrypt(plain, key)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("IV:", hex.EncodeToString(iv))
	fmt.Println("Encrypted:", encrypted)
	fmt.Println(base64.StdEncoding.EncodeToString(encrypted))

	decrypted, err := Decrypt(encrypted, key, iv)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Decrypted:", string(decrypted))
}


// Initialization Vector
func GenerateIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	return iv ,nil
}

func Pkcs7Pad(data []byte) []byte {
	length := aes.BlockSize - (len(data) % aes.BlockSize)
	trailing := bytes.Repeat([]byte{byte(length)}, length)
	return append(data, trailing...)
}

func Pkcs7Unpad(data []byte) []byte {
	dataLength := len(data)
	padLength := int(data[dataLength-1])
	return data[:dataLength-padLength]
}

func Encrypt(data, key []byte) (iv []byte, encrypted []byte, err error) {
	iv, err = GenerateIV()
	if err != nil {
		return nil, nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	padded := Pkcs7Pad(data)
	encrypted = make([]byte, len(padded))
	cbcEncrypter := cipher.NewCBCEncrypter(block, iv)
	cbcEncrypter.CryptBlocks(encrypted, padded)
	return iv, encrypted, nil
}

func Decrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(data))
	cbcDecrypter := cipher.NewCBCDecrypter(block, iv)
	cbcDecrypter.CryptBlocks(decrypted, data)
	return Pkcs7Unpad(decrypted), nil
}
