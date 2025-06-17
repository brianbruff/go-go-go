#!/bin/bash
set -e

echo "Building Go-Go Math Engine and Operation Libraries..."

# Build the add operation plugin
echo "Building add operation plugin..."
cd go-go-math-op-add
go build -buildmode=plugin -o add.so
cd ..

# Build the multiply operation plugin  
echo "Building multiply operation plugin..."
cd go-go-math-op-multiply
go build -buildmode=plugin -o multiply.so
cd ..

# Build the main engine
echo "Building math engine..."
cd go-go-math-engine
go build -o go-go-math-engine
cd ..

echo "✓ All components built successfully!"
echo ""
echo "To run the math engine:"
echo "  cd go-go-math-engine"
echo "  ./go-go-math-engine ../config.json"