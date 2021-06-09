package test

import (
	"github.com/bcowtech/host"
)

type MockComponent struct {
}

func (c *MockComponent) Runner() host.Runner {
	return &MockComponentRunner{
		prefix: "MockComponent",
	}
}
