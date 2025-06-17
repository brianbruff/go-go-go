package main

import (
	"github.com/brianbruff/go-go-go/go-go-math-op"
)

// MultiplyOperation implements the MathOperation interface for multiplication
type MultiplyOperation struct{}

// PerformOP performs multiplication of two numbers
func (op *MultiplyOperation) PerformOP(a, b float64) float64 {
	return a * b
}

// GetName returns the operation name
func (op *MultiplyOperation) GetName() string {
	return "multiply"
}

// NewOperation creates a new MultiplyOperation instance
// This function will be looked up by the plugin loader
func NewOperation() mathop.MathOperation {
	return &MultiplyOperation{}
}