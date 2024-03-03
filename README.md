# wypes

Go library to define type-safe host functions in [wazero](https://github.com/tetratelabs/wazero) and other [WebAssembly](https://webassembly.org/) runtimes.

Features:

* 🛡 Type safe
* 🐎 Fast
* 🔨 Works with any WebAssmebly runtime, like [wazero](https://github.com/tetratelabs/wazero) or [wasman](https://github.com/c0mm4nd/wasman)
* 🧠 Handles for you memory operations
* 👉 Manages external references
* 🧼 Simple and clean API
* 🐜 Can be compiled with TinyGo
* 😎 No reflect, no unsafe, only generics and dedication.

## 📦 Installation

```bash
go get github.com/orsinium-labs/wypes
```

## 🔧 Usage

Define a function using provided types:

```go
func addI32(a wypes.Int32, b wypes.Int32) wypes.Int32 {
    return a + b
}
```

Define a mapping of module and function names to function definitions:

```go
modules := wypes.Modules{
    "env": {
        "add_i32": wypes.H2(addI32),
    },
}
```

Link the modules to the runtime. We provide a convenience method to do this for wazero:

```go
err := modules.DefineWazero(r, nil)
```

That's it! Now the wasm module can call the `env.add_i32` function.

## 🛹 Tricks

The library provides lots of useful types that you can use in your functions. Make sure to [check the docs](https://pkg.go.dev/github.com/orsinium-labs/wypes). A few highlights:

1. Context provides access to the context.Context passed into the guest function call in wazero.
1. [Store](https://pkg.go.dev/github.com/orsinium-labs/wypes#Store) provides access to all the state: memory, stack, references.
1. [Duration](https://pkg.go.dev/github.com/orsinium-labs/wypes#Duration) and [Time](https://pkg.go.dev/github.com/orsinium-labs/wypes#Time) to pass time.Duration and time.Time (as UNIX timestamp).
1. [HostRef](https://pkg.go.dev/github.com/orsinium-labs/wypes#HostRef) can hold a reference to the [Refs](https://pkg.go.dev/github.com/orsinium-labs/wypes#Refs) store of host objects.
1. [Void](https://pkg.go.dev/github.com/orsinium-labs/wypes#Void) is used as the return type for functions that return no value.

See [documentation](https://pkg.go.dev/github.com/orsinium-labs/wypes) for more.
