package files

type MpVault struct {
	filename string
}

func NewMpVault(name string) *MpVault {
	return &MpVault{
		filename: name,
	}
}

func (db *MpVault) Read() ([]byte, error) {
	return readFile(db.filename)
}

func (db *MpVault) Write(content []byte) {
	writeFile(content, db.filename)
}

func (db *MpVault) Remove() {
	removeFile(db.filename)
}
