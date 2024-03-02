//go:build !nowazero
// +build !nowazero

package wypes

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

func DefineWazero(runtime wazero.Runtime, modules Modules) {
	for modName, funcs := range modules {
		mb := runtime.NewHostModuleBuilder(modName)
		for funcName, funcDef := range funcs {
			fb := mb.NewFunctionBuilder()
			fb = fb.WithGoModuleFunction(
				wazeroAdaptHostFunc(funcDef),
				funcDef.ParamValueTypes(),
				funcDef.ResultValueTypes(),
			)
			mb = fb.Export(funcName)
		}
	}
}

func wazeroAdaptHostFunc(hf HostFunc) api.GoModuleFunction {
	return api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
		adaptedStack := SliceStack(stack)
		store := Store{
			Memory:  mod.Memory(),
			Stack:   &adaptedStack,
			Refs:    map[uint32]any{},
			Context: ctx,
		}
		hf.Call(store)
	})
}
