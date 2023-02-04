package filestore

import (
	"fmt"
	"log"
	"os"

	"whydb/config"
	"whydb/types"
)

func createDirIfNotExists(path string) {
	_, err1 := os.Stat(config.DataDir)
	if os.IsNotExist(err1) {
		err3 := os.Mkdir(config.DataDir, 0755)
		if err3 != nil {
			log.Fatalf("Error creating data directory %v", err3)
		}
	} else if err1 != nil {
		log.Fatalf("Error creating data directory %v", err1)
	}

	path = config.DataDir + path
	_, err2 := os.Stat(path)
	if os.IsNotExist(err2) {
		_ = os.Mkdir(path, 0755)
	} else if err2 != nil {
		log.Fatalf("Error creating directory %v", err2)
	}

}

func NewFileStore() types.Store {
	return types.Store{
		Set: func(cat string, key string, data string) error {
			filePath := config.DataDir + cat + "/" + key
			createDirIfNotExists(cat)
			err := os.WriteFile(filePath, []byte(data), 0644)
			if err != nil {
				fmt.Printf("Error writing file %v", err)
				return fmt.Errorf("why?db: error setting key")
			}
			return nil
		},
		Get: func(cat string, key string) (string, error) {
			filePath := config.DataDir + cat + "/" + key
			text, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Error reading file %v", err)
				return "", fmt.Errorf("why?db: error getting key")
			}
			return string(text), nil
		},
		Add: func(cat string, key string, data string) error {
			filePath := config.DataDir + cat + "/" + key

			createDirIfNotExists(cat)
			appendFile, err1 := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
			if err1 != nil {
				fmt.Printf("Error opening file %v", err1)
				return fmt.Errorf("why?db: error adding to key")
			}
			_, err2 := appendFile.WriteString(data)
			if err2 != nil {
				fmt.Printf("Error appending to file %v", err2)
				return fmt.Errorf("why?db: error adding to key")
			}
			_ = appendFile.Close()

			return nil
		},
		Del: func(cat string, key string) error {
			filePath := config.DataDir + cat + "/" + key

			err := os.Remove(filePath)
			if err != nil {
				fmt.Printf("Error deleting file %v", err)
				return fmt.Errorf("why?db: error deleting key")
			}

			return nil
		},
	}
}
