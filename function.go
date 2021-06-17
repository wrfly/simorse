package simorse

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
)

type function struct {
	f    func() error
	name string
}

var (
	funcs = make(map[string]function)
	m     sync.Mutex
)

// Register function to simorse
func Register(alias string, f interface{}) error {
	m.Lock()
	defer m.Unlock()
	_, ok := funcs[alias]
	if ok {
		return fmt.Errorf("function %s already exist", alias)
	}

	ptr := reflect.ValueOf(f).Pointer()
	rFuncName := runtime.FuncForPC(ptr).Name()
	debug("register function: %s=%s", alias, rFuncName)
	funcs[alias] = function{
		f:    f.(func() error),
		name: rFuncName,
	}
	return nil
}

// Call this function
func Call(alias string) error {
	x, ok := funcs[alias]
	if !ok {
		return fmt.Errorf("function %s not found", alias)
	}
	debug("calling function %s", x.name)
	return x.f()
}

// List all the functions registered
func List() {
	for alias, f := range funcs {
		fmt.Printf("alias: %s, func=%s\n", alias, f.name)
	}
}
