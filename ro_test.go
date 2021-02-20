package ro

import (
	"fmt"
	"testing"
)

// The following error vars are used minimize irrelevant overhead during
// comparative benchmarking.
var theError error                           // changes at test time
var anError = fmt.Errorf("an error message") // does not change at test time
var noError error                            // does not change at test time

func TestReturnOn(t *testing.T) {
	theError = anError
	err := ReturnOnMayErr()
	if err != anError {
		t.Errorf("expected %v, got %v", anError, err)
	}
	theError = noError
	err = ReturnOnMayErr()
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

// ReturnOnMayErr uses the ReturnOn mechanism to handle the current value of
// theError
func ReturnOnMayErr() (err error) {
	defer RecoverOn(&err)
	err = theError
	ReturnOn(err)
	return
}

// NormalMayErr uses the normal Go error syntax to handle the current value of
// theError
func NormalMayErr() (err error) {
	err = theError
	if err != nil {
		return
	}
	return
}

// Benchmarks comparing ReturnOn to normal error handling

var e error // external target to prevent overoptimized benchmarks

func BenchmarkNormalWithoutErr(b *testing.B) {
	theError = noError
	for i := 0; i < b.N; i++ {
		e = NormalMayErr()
	}
}

func BenchmarkReturnOnWithoutErr(b *testing.B) {
	theError = noError
	for i := 0; i < b.N; i++ {
		e = ReturnOnMayErr()
	}
}

func BenchmarkNormalWithErr(b *testing.B) {
	theError = anError
	for i := 0; i < b.N; i++ {
		e = NormalMayErr()
	}
}

func BenchmarkReturnOnWithErr(b *testing.B) {
	theError = anError
	for i := 0; i < b.N; i++ {
		e = ReturnOnMayErr()
	}
}
