package wypes

type HostFunc struct {
	Params  []Value
	Results []Value
	Call    func(Store)
}

func (f *HostFunc) NumParams() int {
	return countStackValues(f.Params)
}

func (f *HostFunc) NumResults() int {
	return countStackValues(f.Results)
}

func (f *HostFunc) ParamValueTypes() []ValueType {
	return mergeValueTypes(f.Params)
}

func (f *HostFunc) ResultValueTypes() []ValueType {
	return mergeValueTypes(f.Results)
}

func countStackValues(values []Value) int {
	count := 0
	for _, v := range values {
		count += len(v.ValueTypes())
	}
	return count
}

func mergeValueTypes(values []Value) []ValueType {
	res := make([]ValueType, 0, len(values))
	for _, v := range values {
		res = append(res, v.ValueTypes()...)
	}
	return res
}

func H1[A Lift[A], Z Lower](fn func(A) Z) HostFunc {
	var a A
	var z Z
	return HostFunc{
		Params:  []Value{a},
		Results: []Value{z},
		Call: func(s Store) {
			fn(a.Lift(s)).Lower(s)
		},
	}
}

func H2[A Lift[A], B Lift[B], Z Lower](fn func(A, B) Z) HostFunc {
	var a A
	var b B
	var z Z
	return HostFunc{
		Params:  []Value{a, b},
		Results: []Value{z},
		Call: func(s Store) {
			fn(a.Lift(s), b.Lift(s)).Lower(s)
		},
	}
}
