package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro ao carregar as variáveis de ambiente do arquivo .env")
		return err
	}
	return nil
}

func GetSecretKey() string {
	return os.Getenv("SECRET_KEY")
}