# Password Manager

This is a simple password manager with encryption for local use

## How it works

1) Application retrieves or creates a Vault to store account data
2) User creates a Master Password to access the Vault
3) The Bcrypt hash of the Master Password is stored on the disc
4) When user adds a new account, the password will be encrypted with Master Password by converting it with pbkdf2 to key and using it in an AES/Base64 transforming
5) When user receives the account info, the password will be decrypted to original string