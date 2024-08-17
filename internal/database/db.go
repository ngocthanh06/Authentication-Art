package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type EnvDbConfig struct {
	host     string
	port     string
	username string
	password string
	database string
}

func newEnvDbConfig() *EnvDbConfig {
	return &EnvDbConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("PORT"),
		username: os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASS"),
		database: os.Getenv("DB_NAME"),
	}
}

func (c *EnvDbConfig) GetHost() string {
	return c.host
}

func (c *EnvDbConfig) GetPort() string {
	return c.port
}

func (c *EnvDbConfig) GetUserName() string {
	return c.username
}

func (c *EnvDbConfig) GetPassword() string {
	return c.password
}

func (c *EnvDbConfig) GetDatabaseName() string {
	return c.database
}

func getDnsDatabasePostgres(config *EnvDbConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.GetHost(),
		config.GetUserName(),
		config.GetPassword(),
		config.GetDatabaseName(),
		config.GetPort(),
	)
}

func getTypeDatabaseConnection(config *EnvDbConfig) string {
	switch os.Getenv("DB_CONN") {
	case "postgres":
		return getDnsDatabasePostgres(config)
	default:
		return ""
	}
}

var DB *gorm.DB

func ConnectionDatabase() {
	envConfig := newEnvDbConfig()
	dns := getTypeDatabaseConnection(envConfig)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	fmt.Println("Connection database successfully!")
	DB = db
}
