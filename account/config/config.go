package config

import (
    "os"
)

type AppConfig struct {
	Defaultport string
	Appname string
}
type MailingConfig struct {
	Host string
	Port string
	Username string
	Password string
	
}
type DatabaseConfig struct {
    // Defaultport string
	DBSource   string
	DBSourceWithCred string
	Hosts string
	Username string
	Password string
	Database string
	CollectionUsers string
	CollectionTokens string
	CollectionActivities string
}

type JWTConfig struct {
    SecretKey   string
}

type Config struct {
	AppConfig AppConfig
    Database    DatabaseConfig
	Jwtsecret JWTConfig
	Mailing MailingConfig
}

// New returns a new Config struct
func New() *Config {
    return &Config{
	AppConfig: AppConfig {
	    Defaultport: getEnv("DEFAULT_PORT", ""),
	},
	Database: DatabaseConfig{
		DBSource:   getEnv("DB_SOURCE", ""),
		DBSourceWithCred:   getEnv("DB_SOURCE_WITH_CRED", ""),
		Hosts:   getEnv("HOSTS", ""),
		Username:   getEnv("USERNAME", ""),
		Password:   getEnv("PASSWORD", ""),
		Database:   getEnv("DATABASE", ""),
		CollectionUsers:   getEnv("COLLECTION_USERS", ""),
		CollectionTokens:   getEnv("COLLECTION_TOKENS", ""),
		CollectionActivities:  getEnv("COLLECTION_ACTIVITIES", ""),
	},
	Mailing: MailingConfig{
		Host : getEnv("MAIL_HOST", ""),
		Port : getEnv("MAIL_PORT", ""),
		Username: getEnv("MAIL_USERNAME", ""),
		Password: getEnv("MAIL_PASSWORD", ""),
	},
	Jwtsecret: JWTConfig{
		SecretKey:   getEnv("SECRETKEY", ""),
	},
    }
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
	return value
    }

    return defaultVal
}