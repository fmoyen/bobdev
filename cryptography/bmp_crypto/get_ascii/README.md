# ASCII Code Converter

A Go program that reads a text file and outputs the ASCII codes (decimal and hexadecimal) for each character.

## Features

- Reads text files and processes each character
- Outputs ASCII codes in both decimal and hexadecimal format
- Handles standard ASCII characters (0-127)
- Special formatting for control characters (newlines, tabs, etc.)
- Supports output to console or file
- Skips non-ASCII characters with summary reporting

## Prerequisites

- Go 1.16 or higher installed on your system
- Download from: https://golang.org/dl/

## Installation

1. Clone or download this repository
2. Navigate to the project directory

```bash
cd get_ascii
```

## Usage

### Basic Syntax

```bash
go run main.go <input_file> [output_file]
```

### Arguments

- `input_file` - Path to the text file to process (required)
- `output_file` - Path to save the output (optional, prints to console if not specified)

### Examples

#### Example 1: Output to Console

```bash
go run main.go sample.txt
```

This will display the ASCII codes for each character in `sample.txt` directly in the console.

#### Example 2: Output to File

```bash
go run main.go sample.txt output.txt
```

This will write the ASCII codes to `output.txt` and display a summary in the console.

#### Example 3: Using Your Own File

```bash
go run main.go myfile.txt results.txt
```

## Output Format

Each character is displayed in the following format:

```
Character 'A' -> Decimal: 65, Hex: 0x41
```

### Special Characters

The program handles special characters with appropriate escape sequences:

- Newline: `Character '\n' -> Decimal: 10, Hex: 0x0A`
- Tab: `Character '\t' -> Decimal: 9, Hex: 0x09`
- Space: `Character ' ' -> Decimal: 32, Hex: 0x20`
- Backslash: `Character '\\' -> Decimal: 92, Hex: 0x5C`
- Single quote: `Character '\'' -> Decimal: 39, Hex: 0x27`
- Control characters: `Character '\x00' -> Decimal: 0, Hex: 0x00`

## Sample Output

For the input text "Hi!":

```
Character 'H' -> Decimal: 72, Hex: 0x48
Character 'i' -> Decimal: 105, Hex: 0x69
Character '!' -> Decimal: 33, Hex: 0x21
```

## Building an Executable

To create a standalone executable:

### Windows

```bash
go build -o ascii_converter.exe main.go
```

Then run:

```bash
.\ascii_converter.exe sample.txt
```

### Linux/Mac

```bash
go build -o ascii_converter main.go
```

Then run:

```bash
./ascii_converter sample.txt
```

## Error Handling

The program handles various error conditions:

- **Missing arguments**: Displays usage information
- **File not found**: Shows clear error message with filename
- **Output file creation error**: Reports the issue and exits
- **Non-ASCII characters**: Skips them and reports count in summary

## Testing

A sample test file (`sample.txt`) is included with various ASCII characters:

- Uppercase and lowercase letters
- Numbers
- Special characters and punctuation
- Tabs and newlines

Test the program with:

```bash
go run main.go sample.txt
```

## Technical Details

- **Language**: Go (Golang)
- **ASCII Range**: 0-127 (standard ASCII)
- **Character Encoding**: Processes bytes directly
- **File I/O**: Uses `os.ReadFile` for input, `os.Create` for output

## License

This project is provided as-is for educational and practical use.

## Author

Created as a utility tool for ASCII code conversion and analysis.