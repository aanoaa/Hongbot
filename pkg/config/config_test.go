package config_test

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/aanoaa/hongbot/pkg/config"
)

func TestConfig(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "hongbot")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	cfg1 := mockConfig()
	if err := cfg1.Save(tmpfile.Name()); err != nil {
		t.Error(err)
	}

	cfg2 := mockEmptyConfig()
	if reflect.DeepEqual(cfg1, cfg2) {
		t.Error("Should ne")
	}

	if err := cfg2.Restore(tmpfile.Name()); err != nil {
		log.Fatal(err)
	}

	if !reflect.DeepEqual(cfg1, cfg2) {
		t.Error(err)
	}
}

func mockConfig() config.Config {
	return config.Config{
		Address:  "freenode.net:6667",
		Nick:     "hongbot",
		Pass:     "",
		User:     "",
		Name:     "",
		Channels: []string{"#hongbot"},
	}
}

func mockEmptyConfig() config.Config {
	return config.Config{
		Address:  "",
		Nick:     "",
		Pass:     "",
		User:     "",
		Name:     "",
		Channels: []string{""},
	}
}
