package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"

	"github.com/fatih/color"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/pbkdf2"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encrypt(text, masterPassword string) (string, error) {
	key := getKeyByPassword(masterPassword)
	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return encode(cipherText), nil
}

func Decrypt(text, masterPassword string) (string, error) {
	key := getKeyByPassword(masterPassword)
	block, err := aes.NewCipher(key)

	if err != nil {
		return "", err
	}

	cipherText := decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
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

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)

	if err != nil {
		color.Red("Can't decode password")
		panic(err)
	}

	return data
}

func getKeyByPassword(data string) []byte {
	return pbkdf2.Key([]byte(data), bytes, 4096, 32, sha1.New)
}
