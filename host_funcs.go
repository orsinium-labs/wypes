package wypes

// HostFunc is a wrapped host-defined function.
//
// It is constructed with functions from [H0] to [H8] where the number is
// how many arguments it accepts. If you need more, use [Pair].
//
// There is always exactly one result. If you need to return nothing, use [Void].
// If you want to return 2 or more values, use [Pair], but make sure that the guest
// and the runtime support multi-value returns.
type HostFunc struct {
	Params  []Value
	Results []Value
	Call    func(*Store)
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

// H0 defines a [HostFunc] that accepts no arguments.
func H0[Z Lower](
	fn func() Z,
) HostFunc {
	var z Z
	return HostFunc{
		Params:  []Value{},
		Results: []Value{z},
		Call: func(s *Store) {
			fn().Lower(s)
		},
	}
}

// H1 defines a [HostFunc] that accepts 1 high-level argument.
func H1[A Lift[A], Z Lower](
	fn func(A) Z,
) HostFunc {
	var a A
	var z Z
	return HostFunc{
		Params:  []Value{a},
		Results: []Value{z},
		Call: func(s *Store) {
			a := a.Lift(s)
			fn(a).Lower(s)
		},
	}
}

// H2 defines a [HostFunc] that accepts 2 high-level arguments.
func H2[A Lift[A], B Lift[B], Z Lower](
	fn func(A, B) Z,
) HostFunc {
	var a A
	var b B
	var z Z
	return HostFunc{
		Params:  []Value{a, b},
		Results: []Value{z},
		Call: func(s *Store) {
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b).Lower(s)
		},
	}
}

// H3 defines a [HostFunc] that accepts 3 high-level arguments.
func H3[A Lift[A], B Lift[B], C Lift[C], Z Lower](
	fn func(A, B, C) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c},
		Results: []Value{z},
		Call: func(s *Store) {
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c).Lower(s)
		},
	}
}

// H4 defines a [HostFunc] that accepts 4 high-level arguments.
func H4[A Lift[A], B Lift[B], C Lift[C], D Lift[D], Z Lower](
	fn func(A, B, C, D) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d},
		Results: []Value{z},
		Call: func(s *Store) {
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d).Lower(s)
		},
	}
}

// H5 defines a [HostFunc] that accepts 5 high-level arguments.
func H5[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], Z Lower](
	fn func(A, B, C, D, E) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e},
		Results: []Value{z},
		Call: func(s *Store) {
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e).Lower(s)
		},
	}
}

// H6 defines a [HostFunc] that accepts 6 high-level arguments.
func H6[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], F Lift[F], Z Lower](
	fn func(A, B, C, D, E, F) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var f F
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e, f},
		Results: []Value{z},
		Call: func(s *Store) {
			f := f.Lift(s)
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e, f).Lower(s)
		},
	}
}

// H7 defines a [HostFunc] that accepts 7 high-level arguments.
func H7[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], F Lift[F], G Lift[G], Z Lower](
	fn func(A, B, C, D, E, F, G) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var f F
	var g G
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e, f, g},
		Results: []Value{z},
		Call: func(s *Store) {
			g := g.Lift(s)
			f := f.Lift(s)
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e, f, g).Lower(s)
		},
	}
}

// H8 defines a [HostFunc] that accepts 8 high-level arguments.
func H8[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], F Lift[F], G Lift[G], H Lift[H], Z Lower](
	fn func(A, B, C, D, E, F, G, H) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var f F
	var g G
	var h H
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e, f, g, h},
		Results: []Value{z},
		Call: func(s *Store) {
			h := h.Lift(s)
			g := g.Lift(s)
			f := f.Lift(s)
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e, f, g, h).Lower(s)
		},
	}
}

// H9 defines a [HostFunc] that accepts 9 high-level arguments.
func H9[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], F Lift[F], G Lift[G], H Lift[H], I Lift[I], Z Lower](
	fn func(A, B, C, D, E, F, G, H, I) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var f F
	var g G
	var h H
	var i I
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e, f, g, h, i},
		Results: []Value{z},
		Call: func(s *Store) {
			i := i.Lift(s)
			h := h.Lift(s)
			g := g.Lift(s)
			f := f.Lift(s)
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e, f, g, h, i).Lower(s)
		},
	}
}

// H10 defines a [HostFunc] that accepts 10 high-level arguments.
func H10[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], F Lift[F], G Lift[G], H Lift[H], I Lift[I], J Lift[J], Z Lower](
	fn func(A, B, C, D, E, F, G, H, I, J) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var f F
	var g G
	var h H
	var i I
	var j J
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e, f, g, h, i, j},
		Results: []Value{z},
		Call: func(s *Store) {
			j := j.Lift(s)
			i := i.Lift(s)
			h := h.Lift(s)
			g := g.Lift(s)
			f := f.Lift(s)
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e, f, g, h, i, j).Lower(s)
		},
	}
}

// H11 defines a [HostFunc] that accepts 11 high-level arguments.
func H11[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], F Lift[F], G Lift[G], H Lift[H], I Lift[I], J Lift[J], K Lift[K], Z Lower](
	fn func(A, B, C, D, E, F, G, H, I, J, K) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var f F
	var g G
	var h H
	var i I
	var j J
	var k K
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e, f, g, h, i, j, k},
		Results: []Value{z},
		Call: func(s *Store) {
			k := k.Lift(s)
			j := j.Lift(s)
			i := i.Lift(s)
			h := h.Lift(s)
			g := g.Lift(s)
			f := f.Lift(s)
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e, f, g, h, i, j, k).Lower(s)
		},
	}
}

// H12 defines a [HostFunc] that accepts 12 high-level arguments.
func H12[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], F Lift[F], G Lift[G], H Lift[H], I Lift[I], J Lift[J], K Lift[K], L Lift[L], Z Lower](
	fn func(A, B, C, D, E, F, G, H, I, J, K, L) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var f F
	var g G
	var h H
	var i I
	var j J
	var k K
	var l L
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e, f, g, h, i, j, k, l},
		Results: []Value{z},
		Call: func(s *Store) {
			l := l.Lift(s)
			k := k.Lift(s)
			j := j.Lift(s)
			i := i.Lift(s)
			h := h.Lift(s)
			g := g.Lift(s)
			f := f.Lift(s)
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e, f, g, h, i, j, k, l).Lower(s)
		},
	}
}

// H13 defines a [HostFunc] that accepts 13 high-level arguments.
func H13[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], F Lift[F], G Lift[G], H Lift[H], I Lift[I], J Lift[J], K Lift[K], L Lift[L], M Lift[M], Z Lower](
	fn func(A, B, C, D, E, F, G, H, I, J, K, L, M) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var f F
	var g G
	var h H
	var i I
	var j J
	var k K
	var l L
	var m M
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e, f, g, h, i, j, k, l, m},
		Results: []Value{z},
		Call: func(s *Store) {
			m := m.Lift(s)
			l := l.Lift(s)
			k := k.Lift(s)
			j := j.Lift(s)
			i := i.Lift(s)
			h := h.Lift(s)
			g := g.Lift(s)
			f := f.Lift(s)
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e, f, g, h, i, j, k, l, m).Lower(s)
		},
	}
}

// H14 defines a [HostFunc] that accepts 14 high-level arguments.
func H14[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], F Lift[F], G Lift[G], H Lift[H], I Lift[I], J Lift[J], K Lift[K], L Lift[L], M Lift[M], N Lift[N], Z Lower](
	fn func(A, B, C, D, E, F, G, H, I, J, K, L, M, N) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var f F
	var g G
	var h H
	var i I
	var j J
	var k K
	var l L
	var m M
	var n N
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e, f, g, h, i, j, k, l, m, n},
		Results: []Value{z},
		Call: func(s *Store) {
			n := n.Lift(s)
			m := m.Lift(s)
			l := l.Lift(s)
			k := k.Lift(s)
			j := j.Lift(s)
			i := i.Lift(s)
			h := h.Lift(s)
			g := g.Lift(s)
			f := f.Lift(s)
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e, f, g, h, i, j, k, l, m, n).Lower(s)
		},
	}
}

// H15 defines a [HostFunc] that accepts 15 high-level arguments.
func H15[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], F Lift[F], G Lift[G], H Lift[H], I Lift[I], J Lift[J], K Lift[K], L Lift[L], M Lift[M], N Lift[N], O Lift[O], Z Lower](
	fn func(A, B, C, D, E, F, G, H, I, J, K, L, M, N, O) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var f F
	var g G
	var h H
	var i I
	var j J
	var k K
	var l L
	var m M
	var n N
	var o O
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o},
		Results: []Value{z},
		Call: func(s *Store) {
			o := o.Lift(s)
			n := n.Lift(s)
			m := m.Lift(s)
			l := l.Lift(s)
			k := k.Lift(s)
			j := j.Lift(s)
			i := i.Lift(s)
			h := h.Lift(s)
			g := g.Lift(s)
			f := f.Lift(s)
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o).Lower(s)
		},
	}
}

// H16 defines a [HostFunc] that accepts 16 high-level arguments.
func H16[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], F Lift[F], G Lift[G], H Lift[H], I Lift[I], J Lift[J], K Lift[K], L Lift[L], M Lift[M], N Lift[N], O Lift[O], P Lift[P], Z Lower](
	fn func(A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var f F
	var g G
	var h H
	var i I
	var j J
	var k K
	var l L
	var m M
	var n N
	var o O
	var p P
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p},
		Results: []Value{z},
		Call: func(s *Store) {
			p := p.Lift(s)
			o := o.Lift(s)
			n := n.Lift(s)
			m := m.Lift(s)
			l := l.Lift(s)
			k := k.Lift(s)
			j := j.Lift(s)
			i := i.Lift(s)
			h := h.Lift(s)
			g := g.Lift(s)
			f := f.Lift(s)
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p).Lower(s)
		},
	}
}

// H17 defines a [HostFunc] that accepts 17 high-level arguments.
func H17[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], F Lift[F], G Lift[G], H Lift[H], I Lift[I], J Lift[J], K Lift[K], L Lift[L], M Lift[M], N Lift[N], O Lift[O], P Lift[P], Q Lift[Q], Z Lower](
	fn func(A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var f F
	var g G
	var h H
	var i I
	var j J
	var k K
	var l L
	var m M
	var n N
	var o O
	var p P
	var q Q
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q},
		Results: []Value{z},
		Call: func(s *Store) {
			q := q.Lift(s)
			p := p.Lift(s)
			o := o.Lift(s)
			n := n.Lift(s)
			m := m.Lift(s)
			l := l.Lift(s)
			k := k.Lift(s)
			j := j.Lift(s)
			i := i.Lift(s)
			h := h.Lift(s)
			g := g.Lift(s)
			f := f.Lift(s)
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q).Lower(s)
		},
	}
}

// H18 defines a [HostFunc] that accepts 18 high-level arguments.
func H18[A Lift[A], B Lift[B], C Lift[C], D Lift[D], E Lift[E], F Lift[F], G Lift[G], H Lift[H], I Lift[I], J Lift[J], K Lift[K], L Lift[L], M Lift[M], N Lift[N], O Lift[O], P Lift[P], Q Lift[Q], R Lift[R], Z Lower](
	fn func(A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R) Z,
) HostFunc {
	var a A
	var b B
	var c C
	var d D
	var e E
	var f F
	var g G
	var h H
	var i I
	var j J
	var k K
	var l L
	var m M
	var n N
	var o O
	var p P
	var q Q
	var r R
	var z Z
	return HostFunc{
		Params:  []Value{a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r},
		Results: []Value{z},
		Call: func(s *Store) {
			r := r.Lift(s)
			q := q.Lift(s)
			p := p.Lift(s)
			o := o.Lift(s)
			n := n.Lift(s)
			m := m.Lift(s)
			l := l.Lift(s)
			k := k.Lift(s)
			j := j.Lift(s)
			i := i.Lift(s)
			h := h.Lift(s)
			g := g.Lift(s)
			f := f.Lift(s)
			e := e.Lift(s)
			d := d.Lift(s)
			c := c.Lift(s)
			b := b.Lift(s)
			a := a.Lift(s)
			fn(a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r).Lower(s)
		},
	}
}
