package environment

import (
	"os"
	"os/user"
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
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	home := usr.HomeDir
	return home, nil
}
