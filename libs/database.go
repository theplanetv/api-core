package libs

import "os"

var DB_STRING string

func GetEnv() {
	DB_STRING = os.Getenv("GOOSE_DBSTRING")
}
