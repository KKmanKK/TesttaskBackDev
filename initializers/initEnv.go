package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erorr loading env")
	}
}
