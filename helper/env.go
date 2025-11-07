package helper

import (
	"log"

	"github.com/joho/godotenv"
)
func LoadEnv() {
    // Load .env file from the current directory
    err := godotenv.Load()
    if err != nil {
        // Log a warning if the file isn't found, but don't stop execution, 
        // as variables might be set directly in the shell.
        log.Println("⚠️ Warning: No .env file found. Relying on shell environment variables.")
    } else {
        log.Println("✅ Successfully loaded .env file.")
    }
}