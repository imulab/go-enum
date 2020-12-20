# go-enum
Object as the next best thing before Go supports a more useable enum.

![Default](https://github.com/imulab/go-enum/workflows/Default/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/imulab/go-enum)](https://goreportcard.com/report/github.com/imulab/go-enum)
[![Version](https://img.shields.io/badge/version-0.1.0-blue)](https://img.shields.io/badge/version-0.1.0-blue)

## Motivation

> Go does not have an explicit `enum` keyword to associate string based values with integer based indexes. String values
> are necessary for human consumption while integer based indexes are necessary to save storage spaces. The lack of a
> automatic conversion mechanism makes writing persistence layer quite painful. This small library is to cover the simple
> use case of converting between string values and integer indexes.

## Install

```bash
go get -u github.com/imulab/go-enum
```

## Usage

The simple use case is to create an enum for single value purposes:

```go
options := enum.New("red", "green", "blue")

// Get indexes (which could be saved to database)
options.Index("red")   // 1
options.Index("green") // 2
options.Index("blue")  // 3

// Restore values from indexes
options.Value(1) // red
options.Value(2) // green
options.Value(3) // blue
```

The second use case is to create an enum for multiple value purposes:

```go
multiSelect := enum.NewComposite("one", "two", "three")

// Get indexes, same as before, but in the order of 2
options.Index("one")   // 1
options.Index("two")   // 2
options.Index("three") // 4

// Compute a bitmap to represent multiple values together
multiSelect.BitMap("one", "three") // 5

// Hydrate/Restore values from bitmap
multiSelect.Hydrate(5) // ["one", "three"]
```

For details, please consult [Go Doc](https://pkg.go.dev/github.com/imulab/go-enum).

