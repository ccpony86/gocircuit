package server

import (
	"time"
	"circuit/use/circuit"
)

func (Main) Main(diskpath string, forgetafter time.Duration) (circuit.XPerm, error) {
	srv, err := New(diskpath, forgetafter)
	if err != nil {
		return nil, err
	}
	circuit.Daemonize(func() { <-(chan int)(nil) })
	return circuit.PermRef(srv), nil
}

// Main is a ...
type Main struct{}

func init() {
	circuit.RegisterFunc(Main{})
	circuit.RegisterValue(&Server{})
}
