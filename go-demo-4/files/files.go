package files

import (
	"demo/app-4/output"
	"fmt"
	"os"
)

// JsonDB Struct
type JsonDb struct {
	filename string
}

// Функция-конструктор JsonDb
func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename: name,
	}
}

// Чтение файла
func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Запись файла
func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.PrintError(err)
		return
	}
	fmt.Println("Запись успешна")
}
