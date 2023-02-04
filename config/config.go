package config

const (
	Address = "0.0.0.0"
	Port    = 6010

	StoreType = "file" // "file" or "postgres"

	// file store config
	DataDir = "./data/"

	// postgres store config
	PostgresHost = "<DB URL>"
	PostgresPort = 5432
	PostgresUser = "<username>"
	PostgresPass = "<password>"
	PostgresDB   = "<DB name>"
)
