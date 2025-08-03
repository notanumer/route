# ASCII Hexagon Drawer

A Go program that draws ASCII hexagons using `_`, `/`, and `\` characters.

## Description

This program creates hexagons with customizable dimensions:
- **width**: Number of `_` characters on the top and bottom sides
- **height**: Number of `/` or `\` characters on each of the four side edges

## Input Format

```
t                    // Number of hexagons to draw
width1 height1       // Dimensions for first hexagon
width2 height2       // Dimensions for second hexagon
...
```

## Example

### Input:
```
2
3 2
1 1
```

### Output:
```
 /   \
/     \
/_______\
\_____/
 \___/

/ \
/___\
\_/
```

## Usage

### Run the program:
```bash
go run hexagon.go
```

### Run with input file:
```bash
go run hexagon.go < input.txt
```

### Run tests:
```bash
go test -v
```

### Run benchmarks:
```bash
go test -bench=.
```

## Features

- Clean, idiomatic Go code
- Comprehensive unit tests
- Performance benchmarks
- Proper error handling
- Struct-based design with methods
- Input validation

## Performance

The program is optimized for performance and handles up to 500 test cases efficiently, with each hexagon having dimensions up to 50x50.

## Requirements

- Go 1.16 or later