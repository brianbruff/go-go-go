package mathop

// MathOperation defines the interface that all math operation plugins must implement
type MathOperation interface {
	// PerformOP performs a mathematical operation on two numbers and returns the result
	PerformOP(a, b float64) float64
	// GetName returns the name of the operation (e.g., "add", "multiply")
	GetName() string
}