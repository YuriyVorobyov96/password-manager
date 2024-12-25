package account

import (
	"encoding/json"
	"password/manager/cipher"
	"password/manager/files"
	"strings"
	"time"

	"github.com/fatih/color"
)

const VaultFileName = "data.json"

type Vault struct {
	Accounts  []Account `json:"accounts"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db files.JsonDb
}

func NewVault(db *files.JsonDb) *VaultWithDb {
	data, err := db.ReadFile()

	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}

	var vault Vault

	err = json.Unmarshal(data, &vault)

	if err != nil {
		color.Red("Can't read data")

		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			db: *db,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db: *db,
	}
}

func (vault *VaultWithDb) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
	color.Green("Successfully add account")
}

func (vault *VaultWithDb) FindByUrl(searchString, masterPassword string) {
	if len(vault.Accounts) == 0 {
		color.Yellow("Can't find accounts")

		return
	}

	for _, value := range vault.Accounts {
		if strings.Contains(strings.ToLower(value.Url), strings.ToLower(searchString)) {
			decryptedPassword, err := cipher.Decrypt(value.Password, masterPassword)

			if err != nil {
				color.Red("Can't decrypt password")
				panic(err)
			}

			color.Green("{ Login: %s, Password: %s, URL: %s }\n", value.Login, decryptedPassword, value.Url)
		}
	}
}

func (vault *VaultWithDb) RemoveByUrl(url string) {
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

func (vault *VaultWithDb) Restart() {
	vault.db.RemoveFile()

	newVault := NewVault(&vault.db)

	vault.Accounts = newVault.Accounts
	vault.CreatedAt = newVault.CreatedAt
	vault.UpdatedAt = newVault.UpdatedAt
	vault.db = newVault.db

	vault.save()
}

func (vault *Vault) toBytes() ([]byte, error) {
	return json.Marshal(vault)
}

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()

	data, err := vault.Vault.toBytes()

	if err != nil {
		color.Red("Can't write data")
	}

	vault.db.WriteFile(data)
}
