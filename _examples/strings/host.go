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

func hostPrint(s wypes.String) wypes.Void {
	fmt.Println(s.Unwrap())
	return wypes.Void{}
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
			"print": wypes.H1(hostPrint),
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
	entry := m.ExportedFunction("greet")
	if entry == nil {
		return errors.New("guest function greet is not defined or not exported")
	}
	_, err = entry.Call(ctx)
	if err != nil {
		return fmt.Errorf("call exported function: %v", err)
	}
	return nil
}
