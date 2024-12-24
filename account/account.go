package account

import (
	"crypto/rand"
	"errors"
	"math/big"
	"net/url"
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

func NewAccount(login, password, urlString string) (*Account, error) {
	if len(login) == 0 {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)

	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	acc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if len(acc.Password) == 0 {
		acc.generatePassword(12)
	}

	return acc, nil
}

func (acc *Account) generatePassword(n int) {
	password := make([]rune, n)

	for i := range password {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))

		if err != nil {
			panic(err)
		}

		password[i] = letters[nBig.Int64()]
	}

	acc.Password = string(password)
}