package main

import (
	"github.com/brianbruff/go-go-go/go-go-math-op"
)

// AddOperation implements the MathOperation interface for addition
type AddOperation struct{}

// PerformOP performs addition of two numbers
func (op *AddOperation) PerformOP(a, b float64) float64 {
	return a + b
}

// GetName returns the operation name
func (op *AddOperation) GetName() string {
	return "add"
}

// NewOperation creates a new AddOperation instance
// This function will be looked up by the plugin loader
func NewOperation() mathop.MathOperation {
	return &AddOperation{}
}