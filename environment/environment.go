package environment

import (
	"github.com/mitchellh/go-homedir"
	"os"
	"os/exec"
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

func whichExecutable(exe string) (string, error) {
	path, err := exec.LookPath(exe)
	if err != nil {
		if _, ok := err.(*exec.Error); ok {
			return "", nil
		}
		return "", err
	}
	return path, nil
}

func Which(exe string) string {
	path, _ := whichExecutable(exe)
	return path
}
