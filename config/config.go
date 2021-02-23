package config

import (
	"os"
)

// IsLocal はローカル環境がどうか返します。
func IsLocal() bool {
	return os.Getenv("ENV") == "local"
}

// IsDev はDev環境がどうか返します。
func IsDev() bool {
	return os.Getenv("ENV") == "dev"
}

// IsPrd はPrd環境がどうか返します。
func IsPrd() bool {
	return os.Getenv("ENV") == "prod"
}
