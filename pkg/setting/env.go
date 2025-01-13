package setting

import (
	"log"
	"os"

	"github.com/suisbuds/miao/pkg/errcode"
)

func CheckEnv() {

	// doppler mode
	if os.Args[0] != "doppler" || os.Args[1] != "run" || os.Args[2] != "--" {
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
	if os.Getenv("PORT") == "" {
		log.Fatal(errcode.EnvVarNotSet)
	}
}

func SetEnv(v interface{}) {

	// fmt.Printf("v's type: %T\n", v)

	if dbSetting, ok := v.(**DatabaseSetting); ok {
		if (*dbSetting).Password == "${DB_PASSWORD}" {
			(*dbSetting).Password = os.Getenv("DB_PASSWORD") // 解引用二重指针
		}
		if (*dbSetting).UserName == "${USERNAME}" {
			(*dbSetting).UserName = os.Getenv("USERNAME")
		}
		if (*dbSetting).Port == "${PORT}" {
			(*dbSetting).Port = os.Getenv("PORT")
		}
	}

	if jwtSetting, ok := v.(**JWTSettingS); ok {
		if (*jwtSetting).Secret == "${SECRET" {
			(*jwtSetting).Secret = os.Getenv("SECRET")
		}
		if (*jwtSetting).Issuer == "${ISSUER}" {
			(*jwtSetting).Issuer = os.Getenv("ISSUER")
		}
	}
}
