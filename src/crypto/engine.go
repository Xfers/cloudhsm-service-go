package crypto

import (
	"fmt"

	"github.com/Xfers/go-openssl"
)

const (
	DefaultEngine = "cloudhsm"
)

func Init() {
	defaultSettings()
}

func defaultSettings() {
	err := loadEngine(DefaultEngine)
	if err != nil {
		fmt.Printf("Failed to load default engine %s, error: %s. ", DefaultEngine, err)
		fmt.Println("Falling back to software implementation")
	} else {
		fmt.Printf("Default engine %s is enabled", DefaultEngine)
	}
}

func loadEngine(name string) error {
	engine, err := openssl.EngineById(name)
	if err != nil {
		return err
	}

	err = engine.SetDefault()
	if err != nil {
		return err
	}

	return nil
}
