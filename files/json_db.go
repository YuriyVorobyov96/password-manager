package files

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) ReadFile() ([]byte, error) {
	return readFile(db.filename)
}

func (db *JsonDb) WriteFile(content []byte) {
	writeFile(content, db.filename)
}

func (db *JsonDb) RemoveFile() {
	removeFile(db.filename)
}
