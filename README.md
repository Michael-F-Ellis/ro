# ro
*An alternative for routine Go error handling*

## Usage
```
import "github.com/Michael-F-Ellis/ro"

func myfunc() (err error) {
    defer ro.RecoverOn(&err)

    err = SomeFunctionCall()
    ro.ReturnOn(err)

    // Do more stuff
    // ...

    return
}
```

## How it works
- `ReturnOn` panics if `err` is not `nil`.
- `RecoverOn` recovers from the panic raised by `ReturnOn` and the function exits with whatever error value would have been returned normally. `RecoverOn` does not interfere with panics arising outside of `ReturnOn`.

## Benefits
1. Reduces boilerplate lines by 2/3 for common case where your function simply returns errors without
internal handling, i.e.
```
    err = SomeFunctionCall()
    if err != nil {
        return
    }

    err = OtherFunctionCall()
    if err != nil {
        return
    }
```
        becomes 
```
    err = SomeFunctionCall()
    ro.ReturnOn(err)

    err = SomeFunctionCall()
    ro.ReturnOn(err)
```
2.  Does not interfere with normal error checking.  You can still use `if err != nil {...}` wherever needed -- even within a function that also uses ReturnOn.
## Tradeoffs
There's no free lunch.

### Coding discipline: 
You need to remember to defer `RecoverOn` at the top of your function body with a pointer to an error var used consistently within the function. Otherwise the panic won't be handled when an error occurs.
### Debugging:
Can't set a break on the `return` statement of a specific `ReturnOn` call
### Performance:
For the no error case, the overhead of the ReturnOn call will cost a few nanoseconds compared to testing inline for `err != nil`.

When handling an error, the panic recovery will add a small fraction of a microsecond on typical hardware.

The test file for this package has a benchmark suite comparing both cases to normal error handling.  Output from my recent vintage MacBook Pro is below.

```
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
BenchmarkNormalWithoutErr-8     	1000000000	         0.9096 ns/op
BenchmarkReturnOnWithoutErr-8   	206988118	         5.384 ns/op
BenchmarkNormalWithErr-8        	1000000000	         0.8980 ns/op
BenchmarkReturnOnWithErr-8      	 4848568	       241.6 ns/op
```

### Customizing
The RecoverOn and ReturnOn functions are quite simple (see ro.go)

```
func RecoverOn(err *error) {
	if *err != nil {
		_ = recover()
	}
}

func ReturnOn(err error) {
	if err != nil {
		panic(nil)
	}
}
```
If your application design is amenable to a standarized way annotating, wrapping or dispatching errors, you do that in the body of a custom RecoverOn func before calling recover().