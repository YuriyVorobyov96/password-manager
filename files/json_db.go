package files

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	return readFile(db.filename)
}

func (db *JsonDb) Write(content []byte) {
	writeFile(content, db.filename)
}

func (db *JsonDb) Remove() {
	removeFile(db.filename)
}
