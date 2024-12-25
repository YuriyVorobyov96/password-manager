package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func Encrypt(text, masterPassword string) (string, error) {
	block, err := aes.NewCipher([]byte(masterPassword))

	if err != nil {
		return "", err
	}

	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return Encode(cipherText), nil
}

func Hash(data string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(data), 10)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CheckHash(data, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))

	return err == nil
}
