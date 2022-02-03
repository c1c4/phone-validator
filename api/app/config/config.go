package config

import (
	"os"
)

var (
	DATABASE = ""
	ENV      = "DEV"
)

func LoadEnv() {
	ENV = os.Getenv("ENV")

	if ENV != "TEST" {
		DATABASE = os.Getenv("DATABASE")
	} else {
		DATABASE = os.Getenv("TEST_DATABASE")
	}
}
