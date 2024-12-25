package account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

const vaultFileName = "data.json"

type Vault struct {
	Accounts  []Account `json:"accounts"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	data, err := files.ReadFile(vaultFileName)

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
	vault.save()
	color.Green("Successfully add account")
}

func (vault *Vault) FindByUrl(searchString string) {
	if len(vault.Accounts) == 0 {
		color.Yellow("Can't find accounts")

		return
	}

	for _, value := range vault.Accounts {
		if strings.Contains(strings.ToLower(value.Url), strings.ToLower(searchString)) {
			color.Green("{ Login: %s, Password: %s, URL: %s }\n", value.Login, value.Password, value.Url)
		}
	}
}

func (vault *Vault) RemoveByUrl(url string) {
	if len(vault.Accounts) == 0 {
		color.Yellow("Nothing to delete")

		return
	}

	accounts := []Account{}

	for _, value := range vault.Accounts {
		if value.Url != url {
			accounts = append(accounts, value)
		}
	}

	if len(accounts) < len(vault.Accounts) {
		vault.Accounts = accounts
		vault.save()
		color.Green("Successfully delete accounts")

		return
	}

	color.Yellow("Nothing to delete")
}

func (vault *Vault) Restart() {
	files.RemoveFile(vaultFileName)

	newVault := NewVault()

	vault.Accounts = newVault.Accounts
	vault.CreatedAt = newVault.CreatedAt
	vault.UpdatedAt = newVault.UpdatedAt

	vault.save()
}

func (vault *Vault) toBytes() ([]byte, error) {
	return json.Marshal(vault)
}

func (vault *Vault) save() {
	vault.UpdatedAt = time.Now()

	data, err := vault.toBytes()

	if err != nil {
		color.Red("Can't write data")
	}

	files.WriteFile(data, vaultFileName)
}
