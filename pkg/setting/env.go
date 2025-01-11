package setting

import (
	"log"
	"os"

	"github.com/suisbuds/miao/pkg/errcode"
)

func CheckEnv() {
	
	// doppler mode
	if len(os.Args) < 3 || os.Args[0]!="doppler"|| os.Args[1] != "run" || os.Args[2] != "--" {
        return
    }
	if os.Getenv("DB_PASSWORD") == "" {
		log.Fatal(errcode.EnvVarNotSet)
	}
	if os.Getenv("USERNAME") == "" {
		log.Fatal(errcode.EnvVarNotSet)
	}
	if os.Getenv("SECRET") == "" {
		log.Fatal(errcode.EnvVarNotSet)
	}
	if os.Getenv("ISSUER") == "" {
		log.Fatal(errcode.EnvVarNotSet)
	}
}
