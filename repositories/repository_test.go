package repositories

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := InitDatabase(); err != nil {
		panic(err)
	}

	code := m.Run()

	os.Exit(code)
}
