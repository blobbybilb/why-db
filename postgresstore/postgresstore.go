package postgresstore

import (
	"database/sql"
	"fmt"
	"whydb/config"
	"whydb/types"

	"log"

	_ "github.com/lib/pq"
)

var connStr = fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=require", config.PostgresUser, config.PostgresPass, config.PostgresHost, config.PostgresPort, config.PostgresDB)

var db *sql.DB

func createTableIfNotExists(path string) {
	_, err1 := db.Exec(
		fmt.Sprintf(
			"SELECT 'CREATE DATABASE %s' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '%s')",
			config.PostgresDB, config.PostgresDB))
	if err1 != nil {
		log.Fatalf("Error creating database %v", err1)
	}

	_, err2 := db.Exec("CREATE TABLE IF NOT EXISTS " + path + " (key TEXT PRIMARY KEY, data TEXT)")
	if err2 != nil {
		log.Fatalf("Error creating table")
	}
}

func NewPostgresStore() types.Store {
	var err error

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database")
	}

	return types.Store{
		Set: func(cat string, key string, data string) error {
			createTableIfNotExists(cat)

			_, err := db.Exec("INSERT INTO "+cat+" (key, data) VALUES ($1, $2) ON CONFLICT (key) DO UPDATE SET data = ($2)", key, data)
			if err != nil {
				fmt.Printf("Error inserting data %v", err)
				return fmt.Errorf("why?db: error setting key")
			}

			return nil
		},
		Get: func(cat string, key string) (string, error) {
			var text string
			err := db.QueryRow("SELECT data FROM "+cat+" WHERE key = $1", key).Scan(&text)
			if err != nil {
				fmt.Printf("Error getting data")
				return "", fmt.Errorf("why?db: error getting key")
			}

			return text, nil
		},
		Add: func(cat string, key string, data string) error {
			createTableIfNotExists(cat)

			// TODO: more effecient way
			var text string
			err1 := db.QueryRow("SELECT data FROM "+cat+" WHERE key = $1", key).Scan(&text)
			if err1 != nil {
				fmt.Printf("Error getting data %v", err1)
				return fmt.Errorf("why?db: error adding to key")
			}
			_, err2 := db.Exec("INSERT INTO "+cat+" (key, data) VALUES ($1, $2) ON CONFLICT (key) DO UPDATE SET data = ($2)", key, text+data)
			if err2 != nil {
				fmt.Print("Error inserting data %v", err)
				return fmt.Errorf("why?db: error adding to key")
			}

			return nil
		},
		Del: func(cat string, key string) error {
			_, err := db.Exec("DELETE FROM "+cat+" WHERE key = $1", key)
			if err != nil {
				fmt.Printf("Error deleting data %v", err)
				return fmt.Errorf("why?db: error deleting key")
			}
			return nil
		},
	}
}
