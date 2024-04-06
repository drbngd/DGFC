#!/bin/bash

# Check if a file name is provided as an argument
if [ $# -eq 0 ]; then
    echo "Usage: $0 <file_name>"
    exit 1
fi

# Extract the file name from the command line arguments
file_name=$1

# Run the antlr4 command with the provided file name
antlr4 -Dlanguage=Go -o parser "$file_name"