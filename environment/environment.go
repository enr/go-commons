package environment

import (
	"os"
	"github.com/mitchellh/go-homedir"
	"strings"
)

func GetenvEitherCase(k string) string {
	if k == "" {
		return ""
	}
	if v := os.Getenv(k); v != "" {
		return v
	}
	if v := os.Getenv(strings.ToUpper(k)); v != "" {
		return v
	}
	return os.Getenv(strings.ToLower(k))
}

func UserHome() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return home, nil
}
