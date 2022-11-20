package crypto

/*
#cgo linux windows freebsd openbsd solaris pkg-config: libssl libcrypto
#cgo linux freebsd openbsd solaris CFLAGS: -Wno-deprecated-declarations
#cgo darwin CFLAGS: -I/usr/local/opt/openssl@1.1/include -I/usr/local/opt/openssl/include -Wno-deprecated-declarations
#cgo darwin LDFLAGS: -L/usr/local/opt/openssl@1.1/lib -L/usr/local/opt/openssl/lib -lssl -lcrypto
#cgo windows CFLAGS: -DWIN32_LEAN_AND_MEAN
#include "openssl/engine.h"
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
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
	engine, err := EngineById(name)
	if err != nil {
		return err
	}

	err = engine.SetDefault()
	if err != nil {
		return err
	}

	return nil
}

type Engine struct {
	e *C.ENGINE
}

// Had to re-write this function to get Engine's ID
// as original function Engine struct with e lower case
// which is not accessible from outside (not exported)
func EngineById(name string) (*Engine, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	e := &Engine{
		e: C.ENGINE_by_id(cname),
	}
	if e.e == nil {
		return nil, fmt.Errorf("engine %s missing", name)
	}
	if C.ENGINE_init(e.e) == 0 {
		C.ENGINE_free(e.e)
		return nil, fmt.Errorf("engine %s not initialized", name)
	}
	runtime.SetFinalizer(e, func(e *Engine) {
		C.ENGINE_finish(e.e)
		C.ENGINE_free(e.e)
	})
	return e, nil
}

func (e *Engine) SetDefault() error {
	// Init
	if C.ENGINE_init(e.e) == 0 {
		C.ENGINE_free(e.e)
		return fmt.Errorf("engine %s not initialized", e.e)
	}

	if C.ENGINE_set_default(e.e, C.ENGINE_METHOD_ALL) == 0 {
		return fmt.Errorf("engine %s not set as default", e.e)
	}
	return nil
}
