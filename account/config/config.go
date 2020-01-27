package config

import (
    "os"
    // "strconv"
    // "strings"
)

type DatabaseConfig struct {
    Defaultport string
	DBSource   string
	DBSourceWithCred string
	Hosts string
	Username string
	Password string
	Database string
	CollectionUsers string
	CollectionTokens string
}

type JWTConfig struct {
    SecretKey   string
}

type Config struct {
    Database    DatabaseConfig
    Jwtsecret JWTConfig
    // UserRoles []string
    // MaxUsers  int
}

// New returns a new Config struct
func New() *Config {
    return &Config{
	Database: DatabaseConfig{
	    Defaultport: getEnv("DEFAULT_PORT", ""),
		DBSource:   getEnv("DB_SOURCE", ""),
		DBSourceWithCred:   getEnv("DB_SOURCE_WITH_CRED", ""),
		Hosts:   getEnv("HOSTS", ""),
		Username:   getEnv("USERNAME", ""),
		Password:   getEnv("PASSWORD", ""),
		Database:   getEnv("DATABASE", ""),
		CollectionUsers:   getEnv("COLLECTION_USERS", ""),
		CollectionTokens:   getEnv("COLLECTION_TOKENS", ""),
	},
	Jwtsecret: JWTConfig{
		SecretKey:   getEnv("SECRETKEY", ""),
	},
	// DebugMode: getEnvAsBool("DEBUG_MODE", true),
	// UserRoles: getEnvAsSlice("USER_ROLES", []string{"admin"}, ","),
	// MaxUsers:  getEnvAsInt("MAX_USERS", 1),
    }
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
	return value
    }

    return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
// func getEnvAsInt(name string, defaultVal int) int {
//     valueStr := getEnv(name, "")
//     if value, err := strconv.Atoi(valueStr); err == nil {
// 	return value
//     }

//     return defaultVal
// }

// Helper to read an environment variable into a bool or return default value
// func getEnvAsBool(name string, defaultVal bool) bool {
//     valStr := getEnv(name, "")
//     if val, err := strconv.ParseBool(valStr); err == nil {
// 	return val
//     }

//     return defaultVal
// }

// Helper to read an environment variable into a string slice or return default value
// func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
//     valStr := getEnv(name, "")

//     if valStr == "" {
// 	return defaultVal
//     }

//     val := strings.Split(valStr, sep)

//     return val
// }