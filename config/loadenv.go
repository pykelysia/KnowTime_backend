package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pykelysia/pyketools"
)

func LoadEnv(p string) {
	if err := godotenv.Load(p); err != nil {
		pyketools.Fatalf(".env not loaded: %v", err)
	}

	ChatModelName = os.Getenv("CHAT_MODEL_NAME")
	ImageModelName = os.Getenv("IMAGE_MODEL_NAME")
	BaseURL = os.Getenv("BASE_URL")
	APIKey = os.Getenv("API_KEY")

	JwtKey = []byte(os.Getenv("JWT_KEY"))
}
