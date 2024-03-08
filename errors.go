package wypes

import "errors"

var (
	ErrRefNotFound = errors.New("HostRef with the given ID is not found in Refs")
	ErrMemRead     = errors.New("Memory.Read is out of bounds")
	ErrMemWrite    = errors.New("Memory.Write is out of bounds")
	ErrRefCast     = errors.New("Reference returned by Refs.Get is not of the type expected by HostRef")
)
