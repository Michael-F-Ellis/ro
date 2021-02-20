// Package ro reduces error handling boilerplate at the expense of a little cpu
// time and requires some coding disipline to use correctly.
package ro

// ReturnOn panics if e is not nil on the expectation that a deferred recovery
// function will catch and handle it.
func ReturnOn(e error) {
	if e != nil {
		panic(nil)
	}
}

// RecoverOn recovers from panics raised by ReturnOn.  Other panics will
// propagate without recovery. Invoke it with defer at the top of a function
// with a pointer to the error var that will be used within the function.
func RecoverOn(err *error) {
	if *err != nil {
		_ = recover()
	}
}
