# dailyDSA-logic

This repository contains daily coding problems and their solutions in both Go and Python. The goal is to practice and improve problem-solving skills by implementing various algorithms and logic-based challenges.

## Structure

- **golang/**: Contains Go solutions for daily DSA problems.
    - `main.go`: Entry point for running Go solutions.
    - Individual `.go` files for each problem (e.g., `check_even_or_odd.go`, `closest_number.go`, etc.)
    - `go.mod`: Go module definition.
- **python/**: Contains Python solutions for daily DSA problems.
    - `main.py`: Entry point for running Python solutions.

## How to Use

### Go
1. Navigate to the `golang` directory:
   ```bash
   cd golang
   ```
2. Run the main file:
   ```bash
   go run main.go
   ```

> **Note:** Individual Go files (such as `check_even_or_odd.go`, `closest_number.go`, etc.) are not meant to be run directly. They are imported as modules and used within `main.go`. To execute a specific problem, modify `main.go` to call the desired function from the relevant file.

### Python
1. Navigate to the `python` directory:
   ```bash
   cd python
   ```
2. Run the main file:
   ```bash
   python3 main.py
   ```

## Contributing

Feel free to add new problems or improve existing solutions. Please follow the existing folder structure and naming conventions.

## License

This project is licensed under the MIT License.
