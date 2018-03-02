package services

import (
	"fmt"
	"os"
)

// Self provides serivce lookup for the current process, useful
// when running Inspeqtor from the Makefile,
// otherwise it can't find itself as a service.
type Self struct{}

func (m *Self) Name() string { return "self" }

func (m *Self) Restart(name string) error {
	return fmt.Errorf("Cannot restart myself")
}

func (m *Self) Reload(name string) error {
	return fmt.Errorf("Cannot reload myself")
}

func (m *Self) LookupService(name string) (*ProcessStatus, error) {
	if name == "inspeqtor" {
		return &ProcessStatus{
			Pid:    os.Getpid(),
			Status: Up}, nil
	}
	return nil, &ServiceError{Init: m.Name(), Name: name, Err: ErrServiceNotFound}
}
