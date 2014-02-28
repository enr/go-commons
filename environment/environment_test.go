package environment

import (
	"os"
	"testing"
)

type envvar struct {
	key   string
	value string
}

var vars = []envvar{
	{"TestGetenvEitherCase", "TestGetenvEitherCase_value"},
	{"TESTGETENVEITHERCASE", "TestGetenvEitherCase_value"},
	{"testgetenveithercase", "TestGetenvEitherCase_value"},
}

func TestGetenvEitherCase(t *testing.T) {

	for _, env := range vars {
		os.Setenv(env.key, env.value)
		res := GetenvEitherCase("TestGetenvEitherCase")
		if res != env.value {
			t.Errorf(`Env %s, got %s, expected %s`, env.key, res, env.value)
		}
		os.Setenv(env.key, "")
	}
}

func TestGetenvEitherCase_emptykey(t *testing.T) {
	res := GetenvEitherCase("")
	if res != "" {
		t.Errorf(`Env "", got "%s", expected ""`, res)
	}
}