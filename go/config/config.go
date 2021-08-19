package config

import "os"

var (
	DB_URL      = os.Getenv("MONGO_URL")
	DB_NAME     = os.Getenv("MONGO_DATABASE")
	DB_USER     = os.Getenv("MONGO_USERNAME")
	DB_PASSWORD = os.Getenv("MONGO_PASSWORD")
	ENVIRONMENT = os.Getenv("ENVIRONMENT")
	SECRET_KEY  = os.Getenv("SECRET_KEY")
)
