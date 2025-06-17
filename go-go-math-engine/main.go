package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"plugin"
	"strconv"
	"strings"

	"github.com/brianbruff/go-go-go/go-go-math-op"
)

// Config represents the configuration structure for available operations
type Config struct {
	Operations []OperationConfig `json:"operations"`
}

// OperationConfig represents individual operation configuration
type OperationConfig struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// Engine represents the math engine
type Engine struct {
	operations map[string]mathop.MathOperation
	config     Config
}

// NewEngine creates a new math engine instance
func NewEngine() *Engine {
	return &Engine{
		operations: make(map[string]mathop.MathOperation),
	}
}

// LoadConfig loads the configuration from the specified file
func (e *Engine) LoadConfig(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&e.config); err != nil {
		return fmt.Errorf("failed to decode config: %w", err)
	}

	return nil
}

// LoadOperations loads all operations specified in the config
func (e *Engine) LoadOperations() error {
	fmt.Println("Loading math operations...")

	for _, opConfig := range e.config.Operations {
		err := e.loadOperation(opConfig)
		if err != nil {
			fmt.Printf("Failed to load operation %s: %v\n", opConfig.Name, err)
			continue
		}
		fmt.Printf("✓ Loaded operation: %s\n", opConfig.Name)
	}

	if len(e.operations) == 0 {
		return fmt.Errorf("no operations were successfully loaded")
	}

	return nil
}

// loadOperation loads a single operation plugin
func (e *Engine) loadOperation(opConfig OperationConfig) error {
	// Load the plugin
	p, err := plugin.Open(opConfig.Path)
	if err != nil {
		return fmt.Errorf("failed to open plugin: %w", err)
	}

	// Look up the NewOperation symbol
	symNewOperation, err := p.Lookup("NewOperation")
	if err != nil {
		return fmt.Errorf("failed to find NewOperation symbol: %w", err)
	}

	// Assert that it's a function that returns MathOperation
	newOperationFunc, ok := symNewOperation.(func() mathop.MathOperation)
	if !ok {
		return fmt.Errorf("NewOperation symbol is not a valid function")
	}

	// Create the operation instance
	operation := newOperationFunc()

	// Store the operation
	e.operations[opConfig.Name] = operation
	return nil
}

// ListOperations lists all available operations
func (e *Engine) ListOperations() {
	fmt.Println("\nAvailable operations:")
	for name := range e.operations {
		fmt.Printf("- %s\n", name)
	}
}

// getNumbers prompts the user for two numbers
func (e *Engine) getNumbers() (float64, float64, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter first number: ")
	line1, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, err
	}
	num1, err := strconv.ParseFloat(strings.TrimSpace(line1), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid first number: %w", err)
	}

	fmt.Print("Enter second number: ")
	line2, err := reader.ReadString('\n')
	if err != nil {
		return 0, 0, err
	}
	num2, err := strconv.ParseFloat(strings.TrimSpace(line2), 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid second number: %w", err)
	}

	return num1, num2, nil
}

// getOperation prompts the user to select an operation
func (e *Engine) getOperation() (mathop.MathOperation, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Select operation: ")
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	opName := strings.TrimSpace(line)
	operation, exists := e.operations[opName]
	if !exists {
		return nil, fmt.Errorf("operation '%s' not found", opName)
	}

	return operation, nil
}

// Run starts the engine's interactive mode
func (e *Engine) Run() error {
	fmt.Println("\n=== Go-Go Math Engine ===")
	e.ListOperations()

	for {
		fmt.Println("\n--- New Calculation ---")

		// Get two numbers from user
		num1, num2, err := e.getNumbers()
		if err != nil {
			fmt.Printf("Error getting numbers: %v\n", err)
			continue
		}

		// Get operation from user
		operation, err := e.getOperation()
		if err != nil {
			fmt.Printf("Error selecting operation: %v\n", err)
			continue
		}

		// Perform the operation
		result := operation.PerformOP(num1, num2)

		// Display result
		fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operation.GetName(), num2, result)

		// Ask if user wants to continue
		fmt.Print("\nPerform another calculation? (y/n): ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if strings.ToLower(strings.TrimSpace(line)) != "y" {
			break
		}
	}

	fmt.Println("Thank you for using Go-Go Math Engine!")
	return nil
}

func main() {
	engine := NewEngine()

	// Load configuration
	configPath := "config.json"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}

	if err := engine.LoadConfig(configPath); err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Load operations from plugins
	if err := engine.LoadOperations(); err != nil {
		fmt.Printf("Error loading operations: %v\n", err)
		os.Exit(1)
	}

	// Run the engine
	if err := engine.Run(); err != nil {
		fmt.Printf("Engine error: %v\n", err)
		os.Exit(1)
	}
}