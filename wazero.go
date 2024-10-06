//go:build !nowazero
// +build !nowazero

package wypes

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

// DefineWazero registers all the host modules in the given wazero runtime.
func (ms Modules) DefineWazero(runtime wazero.Runtime, refs Refs) error {
	if refs == nil {
		refs = NewMapRefs()
	}
	for modName, funcs := range ms {
		err := funcs.DefineWazero(runtime, modName, refs)
		if err != nil {
			return err
		}
	}
	return nil
}

// DefineWazero registers the host module in the given wazero runtime.
func (m Module) DefineWazero(runtime wazero.Runtime, modName string, refs Refs) error {
	var err error
	mb := runtime.NewHostModuleBuilder(modName)
	for funcName, funcDef := range m {
		fb := mb.NewFunctionBuilder()
		fb = fb.WithGoModuleFunction(
			wazeroAdaptHostFunc(funcDef, refs),
			funcDef.ParamValueTypes(),
			funcDef.ResultValueTypes(),
		)
		mb = fb.Export(funcName)
	}
	_, err = mb.Instantiate(context.Background())
	return err
}

func wazeroAdaptHostFunc(hf HostFunc, refs Refs) api.GoModuleFunction {
	return api.GoModuleFunc(func(ctx context.Context, mod api.Module, stack []uint64) {
		adaptedStack := SliceStack(stack)
		store := Store{
			Memory:  mod.Memory(),
			Stack:   &adaptedStack,
			Refs:    refs,
			Context: ctx,
		}
		hf.Call(&store)
	})
}
