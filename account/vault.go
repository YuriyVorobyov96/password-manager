package account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

const dataFileName = "data.json"

type Vault struct {
	Accounts  []Account `json:"accounts"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (vault *Vault) ToBytes() ([]byte, error) {
	return json.Marshal(vault)
}

func NewVault() *Vault {
	data, err := files.ReadFile(dataFileName)

	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault

	err = json.Unmarshal(data, &vault)

	if err != nil {
		color.Red("Can't read data")

		return &Vault{
			Accounts:  []Account{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()

	data, err := vault.ToBytes()

	if err != nil {
		color.Red("Can't write data")
	}

	files.WriteFile(data, dataFileName)
}

func (vault *Vault) FindByUrl(searchString string) {
	if len(vault.Accounts) == 0 {
		color.Yellow("Can't find accounts")
	}

	for _, value := range vault.Accounts {
		if strings.Contains(strings.ToLower(value.Url), strings.ToLower(searchString)) {
			color.Green("{ Login: %s, Password: %s, URL: %s }\n", value.Login, value.Password, value.Url)
		}
	}
}
