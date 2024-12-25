package account

import (
	"crypto/rand"
	"errors"
	"math/big"
	"net/url"
	"password/manager/cipher"
	"password/manager/output"
	"time"

	"github.com/fatih/color"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#^")

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) OutputData() {
	color.Cyan(acc.Login, acc.Password, acc.Url, "\n")
}

func NewAccount(login, password, urlString, masterPassword string) (*Account, error) {
	if len(login) == 0 {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	var encryptedPassword string

	if len(password) > 0 {
		encryptedPassword, err = cipher.Encrypt(string(password), masterPassword)

		if err != nil {
			return nil, errors.New("INVALID_PASSWORD")
		}
	}

	acc := &Account{
		Login:     login,
		Password:  encryptedPassword,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if len(acc.Password) == 0 {
		acc.generatePassword(12, masterPassword)
	}

	return acc, nil
}

func (acc *Account) generatePassword(n int, masterPassword string) {
	password := make([]rune, n)

	for i := range password {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))

		if err != nil {
			panic(err)
		}

		password[i] = letters[nBig.Int64()]
	}

	encryptedPassword, err := cipher.Encrypt(string(password), masterPassword)

	if err != nil {
		output.PrintError("Error on password generation")

		return
	}

	acc.Password = encryptedPassword
}
