# merideum-go

Reference implementation of Merideum in Golang

# Build

Go version: 1.18
Antlr4 version: 4.9.3

# Getting Started

Make sure Go is installed.

Make sure Antlr4 is installed and aliased: https://github.com/antlr/antlr4/blob/master/doc/getting-started.md

# Making changes to the grammar

Pull the ANTLR4 grammar from the repository: `./scripts/get-grammar.sh`

> Make sure you have permission to run the script: `chmod +x ./scripts/get-grammar.sh`

Generate the Go files: `antlr4 -Dlanguage=Go -visitor -o ./parser Merit.g4`

# Publishing the module

Always publish locally with a SNAPSHOT tag and test with the `merideum-go` implementation.
