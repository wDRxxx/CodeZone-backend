package config

import "github.com/joho/godotenv"

// Load loads given .env files
func Load(filenames ...string) error {
	return godotenv.Load(filenames...)
}
