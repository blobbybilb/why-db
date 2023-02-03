package filestore

import (
	"os"
	"whydb/config"
	"whydb/types"
)

func CreateDirIfNotExists(path string) {
	_, err1 := os.Stat(config.DataDir)
	if os.IsNotExist(err1) {
		_ = os.Mkdir(config.DataDir, 0755)
	}

	path = config.DataDir + path
	_, err2 := os.Stat(path)
	if os.IsNotExist(err2) {
		_ = os.Mkdir(path, 0755)
	}
}

func NewFileStore() types.Store {
	return types.Store{
		Set: func(cat string, key string, data string) error {
			filePath := config.DataDir + cat + "/" + key

			CreateDirIfNotExists(cat)
			_ = os.WriteFile(filePath, []byte(data), 0644)

			return nil
		},
		Get: func(cat string, key string) (string, error) {
			filePath := config.DataDir + cat + "/" + key
			text, _ := os.ReadFile(filePath)
			return string(text), nil
		},
		Add: func(cat string, key string, data string) error {
			filePath := config.DataDir + cat + "/" + key

			CreateDirIfNotExists(cat)
			appendFile, _ := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
			_, _ = appendFile.WriteString(data)
			_ = appendFile.Close()

			return nil
		},
		Del: func(cat string, key string) error {
			filePath := config.DataDir + cat + "/" + key

			_ = os.Remove(filePath)
			return nil
		},
	}
}
