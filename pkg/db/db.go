package db

import (
	"encoding/json"
	"os"
)

type AppDB struct {
	Admins []int64 `json:"admins"`
}

type DBManager struct {
	filename string
	Config   *AppDB
}

func NewDB(filename string) func() *DBManager {
	cnf := &DBManager{
		filename: filename,
	}
	return func() *DBManager {
		configText, err := cnf.Load()
		if err != nil {
			panic(err)
		}
		cnf.Config = configText

		return cnf
	}
}

func (cm *DBManager) Load() (*AppDB, error) {
	// Проверяем, существует ли файл
	if _, err := os.Stat(cm.filename); os.IsNotExist(err) {
		// Файл не существует, создаем его с дефолтной конфигурацией
		if err := cm.createDefaultConfig(); err != nil {
			return nil, err
		}
		return cm.Config, nil
	}

	file, err := os.Open(cm.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(cm.Config); err != nil {
		return nil, err
	}
	return cm.Config, nil
}

func (cm *DBManager) Save() error {
	file, err := os.Create(cm.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(cm.Config)
}

func (cm *DBManager) Get() *AppDB {
	return cm.Config
}

func (cm *DBManager) Set(config *AppDB) {
	cm.Config = config
}

// createDefaultConfig создает файл с дефолтной конфигурацией
func (cm *DBManager) createDefaultConfig() error {
	cm.Config = &AppDB{
		Admins: []int64{},
	}
	return cm.Save()
}
