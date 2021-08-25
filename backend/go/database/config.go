package database

import "fmt"

//Config to maintain DB configuration properties
type Config struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name)
	return connectionString
}
