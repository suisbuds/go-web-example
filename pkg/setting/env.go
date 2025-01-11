package setting

import (
	"log"
	"os"

	"github.com/suisbuds/miao/pkg/errcode"
)

func CheckEnv() {
	if os.Getenv("USERNAME") == "" {
		log.Fatal(errcode.EnvVarNotSet)
	}
}
