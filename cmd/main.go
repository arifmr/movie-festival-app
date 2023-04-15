package main

import "os"

func main() {
	// Read dotenv file
	if os.Getenv("APP_ENV") == "local" {
		godotenv.Load()
	}

}
