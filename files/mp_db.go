package files

type MpDb struct {
	filename string
}

func NewMpDb(name string) *MpDb {
	return &MpDb{
		filename: name,
	}
}

func (db *MpDb) ReadFile() ([]byte, error) {
	return readFile(db.filename)
}

func (db *MpDb) WriteFile(content []byte) {
	writeFile(content, db.filename)
}

func (db *MpDb) RemoveFile() {
	removeFile(db.filename)
}
