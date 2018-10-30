package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
)

func Encryptor(s string) (string, error) {
	pass := []byte(NewPassword(s))

	cipherBlock, err := aes.NewCipher([]byte(tokenPetter))
	if err != nil {
		log.Println("error cipherBlock")
		return "", err
	}

	endcodeRawString := base64.StdEncoding.EncodeToString(pass)
	cipherText := make([]byte, aes.BlockSize+len(endcodeRawString))

	t1 := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, t1); err != nil {
		log.Println("error Readfull")
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(cipherBlock, t1)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], []byte(endcodeRawString))
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decryptor(s string) (string, error) {

	cipherBlock, err := aes.NewCipher([]byte(tokenPetter))
	if err != nil {
		return "", err
	}
	decodedString, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	if len(decodedString) < aes.BlockSize {
		return "", errors.New("cipher text is too short")
	}
	iv := decodedString[:aes.BlockSize]
	decodedString = decodedString[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(cipherBlock, iv)
	cfb.XORKeyStream(decodedString, decodedString)
	decryptedString, err := base64.StdEncoding.DecodeString(string(decodedString))
	if err != nil {
		return "", err
	}
	return string(decryptedString), nil
}

func NewPassword(s string) string {
	return "!@#$%1401" + s + "95*&^()"
}
