# 🎄 Advent of Code 2025

Multi-language solutions for Advent of Code 2025 in Go, Python, and TypeScript.

## Quick Start

### Prerequisites
- Go 1.21+
- Python 3.10+
- Node.js 18+ with pnpm

### Running Solutions

Using Make:
```bash
make python-d1p1    # Python - Day 1 Part 1
make python-d1p2    # Python - Day 1 Part 2
make go-d1p1        # Go - Day 1 Part 1
make go-d1p2        # Go - Day 1 Part 2
make ts-d1p1        # TypeScript - Day 1 Part 1
make ts-d1p2        # TypeScript - Day 1 Part 2
```

Or directly:
```bash
# Python
cd 2025/python && python main.py d1p1

# Go
cd 2025/go && go run main.go d1p1

# TypeScript
cd 2025/ts && pnpm start d1p1
```

## Project Structure

```
2025/
├── python/      # Python solutions
├── go/          # Go solutions
└── ts/          # TypeScript solutions
```

## Language Implementations

- **Python**: Type-safe commands with Enum pattern
- **Go**: String-based typed commands
- **TypeScript**: Compiled TypeScript with tsx runner

Each language follows the same command-line interface for consistency.
