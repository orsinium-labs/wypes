package main

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log"

	"github.com/orsinium-labs/wypes"
	"github.com/tetratelabs/wazero"
)

//go:embed guest.wasm
var source []byte

func addI32(a wypes.Int32, b wypes.Int32) wypes.Int32 {
	return a - b
}

func main() {
	err := run()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func run() error {
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	modules := wypes.Modules{
		"env": {
			"add_i32": wypes.H2(addI32),
		},
	}
	err := modules.DefineWazero(r, nil)
	if err != nil {
		return fmt.Errorf("define host functions: %v", err)
	}
	m, err := r.Instantiate(ctx, source)
	if err != nil {
		return fmt.Errorf("instantiate module: %v", err)
	}
	entry := m.ExportedFunction("run")
	if entry == nil {
		return errors.New("guest function run is not defined or not exported")
	}
	res, err := entry.Call(ctx)
	if err != nil {
		return fmt.Errorf("call exported function: %v", err)
	}
	fmt.Printf("Result: (4 + 5) * 2 = %v", res[0])
	return nil
}
