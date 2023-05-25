package env

import (
	"fmt"
	"os"
)

var baseUrlName, baseUrlValue string = "BASE_URL", "127.0.0.1:8080"

func Init() {
	os.Setenv(baseUrlName, baseUrlValue)
	fmt.Println("Listening on route:", os.Getenv(baseUrlName), ". [If it's 127.0.0.1:, it means 'localhost:']")
}

func GetBaseUrl() string {
	return os.Getenv(baseUrlName)
}
