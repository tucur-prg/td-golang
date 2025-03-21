package main

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/cipher"
	"encoding/hex"
	"encoding/base64"
	"fmt"
)

func main() {
	text := "test message"

	// key size: 16 or 24 or 32
	keyString := "000000000000000000000000000000"
//	keyString := "645E739A7F9F162725C1533DC2C5E827"
//	keyString := "645E739A7F9F162725C1533DC2C5E827645E739A7F9F1627"
//	keyString := "645E739A7F9F162725C1533DC2C5E827645E739A7F9F162725C1533DC2C5E827"

	plain := []byte(text)

	fmt.Println("Key:", len(keyString), ":", keyString)

	key, err := hex.DecodeString(keyString)
	if err != nil {
		panic(err.Error())
	}

	encrypted, err := Encrypt(plain, key)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Encrypted:", encrypted)
	fmt.Println(base64.StdEncoding.EncodeToString(encrypted))

	decrypted, err := Decrypt(encrypted, key)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Decrypted:", string(decrypted))
}

func Encrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nil, nonce, data, nil)
	cipherText = append(nonce, cipherText...)

	return cipherText, nil
}

func Decrypt(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := data[:gcm.NonceSize()]

	plainByte, err := gcm.Open(nil, nonce, data[gcm.NonceSize():], nil)
	if err != nil {
		return nil, err
	}

	return plainByte, nil
}
