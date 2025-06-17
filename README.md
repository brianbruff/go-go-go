# go-go-go
A pluggable Go math engine with operation libraries

## Overview
This project implements a pluggable system in Go for mathematical operations, demonstrating dynamic plugin loading and runtime discovery of implementations.

## Architecture

### Components
- **go-go-math-op**: Base library defining the `MathOperation` interface
- **go-go-math-op-add**: Plugin implementing addition operation
- **go-go-math-op-multiply**: Plugin implementing multiplication operation  
- **go-go-math-engine**: Main process that loads plugins dynamically

### Key Features
- Runtime plugin discovery via configuration file
- Loose coupling - engine only depends on base interface
- Extensible - new operations can be added without modifying the engine
- Dynamic loading using Go's plugin package

## Interface Definition
All math operations must implement the `MathOperation` interface:
```go
type MathOperation interface {
    PerformOP(a, b float64) float64
    GetName() string
}
```

## Building and Running

### Build All Components
```bash
./build.sh
```

### Run the Math Engine
```bash
cd go-go-math-engine
./go-go-math-engine ../config.json
```

### Configuration
The `config.json` file specifies available operation plugins:
```json
{
  "operations": [
    {
      "name": "add",
      "path": "../go-go-math-op-add/add.so"
    },
    {
      "name": "multiply", 
      "path": "../go-go-math-op-multiply/multiply.so"
    }
  ]
}
```

## Usage Example
```
Loading math operations...
✓ Loaded operation: add
✓ Loaded operation: multiply

=== Go-Go Math Engine ===

Available operations:
- add
- multiply

--- New Calculation ---
Enter first number: 5
Enter second number: 3
Select operation: add
Result: 5.00 add 3.00 = 8.00
```

## Adding New Operations
1. Create a new directory (e.g., `go-go-math-op-subtract`)
2. Implement the `MathOperation` interface
3. Export a `NewOperation() mathop.MathOperation` function
4. Build as a plugin: `go build -buildmode=plugin -o subtract.so`
5. Add to `config.json`

## Requirements
- Go 1.23+
- Linux/Unix environment (Go plugins require CGO and are not supported on Windows)
