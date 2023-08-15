package env

import (
	"github.com/Xavier577/gin-project-starter/pkg/number"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	Production  = "production"
	Staging     = "staging"
	Development = "development"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}

type Value interface {
	string | number.Number
}

func Get[T Value](key string) T {
	var a interface{} = os.Getenv(key)
	return a.(T)
}
