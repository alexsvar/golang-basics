package cloud

// CloudDb struct
type CloudDb struct {
	url string
}

// Функция-конструктор создания CloudDb
func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		url: url,
	}
}

// Чтение
func (db *CloudDb) Read() ([]byte, error) {
	return []byte{}, nil
}

// Запись
func (db *CloudDb) Write(content []byte) {
}
