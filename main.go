package main

import (
	"fmt"
	"time"
)

type OptFunc func(opts *Opts)

type Opts struct {
	timeout time.Duration
	tls     bool
	path    string
}

type Module struct {
	Opts
}

func defaultOpts() Opts {
	return Opts{
		timeout: 5 * time.Second,
		tls:     false,
		path:    "/default/path/",
	}
}

// withTLS it is OptFunc its self
func withTLS(opts *Opts) {
	opts.tls = true
}

// withTimeout take option value, construct inside and return OptFunc
func withTimeout(timeout time.Duration) OptFunc {
	return func(opts *Opts) {
		opts.timeout = timeout
	}
}

func withPath(path string) OptFunc {
	return func(opts *Opts) {
		opts.path = path
	}
}

func NewModule(opts ...OptFunc) *Module {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}

	return &Module{
		Opts: o,
	}
}

func main() {
	// example with default options
	m1 := NewModule()
	fmt.Printf("%+v\n", m1)

	// example with custom options
	m2 := NewModule(withTLS, withTimeout(50*time.Millisecond), withPath("some/custom/path"))
	fmt.Printf("%+v\n", m2)
}
