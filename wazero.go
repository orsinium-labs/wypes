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
		store := Store{
			Memory:  mod.Memory(),
			Stack:   &wazeroStackAdapter{stack},
			Refs:    map[uint32]any{},
			Context: ctx,
		}
		hf.Call(store)
	})
}

type wazeroStackAdapter struct {
	raw []uint64
}

func (s *wazeroStackAdapter) Push(v Raw) {
	s.raw = append(s.raw, v)
}

func (s *wazeroStackAdapter) Pop() Raw {
	idx := len(s.raw) - 1
	v := s.raw[idx]
	s.raw = s.raw[:idx]
	return v
}
