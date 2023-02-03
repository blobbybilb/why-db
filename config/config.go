package config

const (
	Address = "0.0.0.0"
	Port    = 6010

	StoreType = "postgres" // "file" or "postgres"

	// file store config
	DataDir = "./data/"

	// postgres store config
	PostgresHost = "<ip, domain, localhost, etc.>"
	PostgresPort = 5432
	PostgresUser = "<username>"
	PostgresPass = "<password>"
	PostgresDB   = "<db_name>"
)
