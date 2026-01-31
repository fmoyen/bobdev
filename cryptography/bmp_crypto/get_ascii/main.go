package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Check command-line arguments
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	inputFile := os.Args[1]
	var outputFile string
	if len(os.Args) >= 3 {
		outputFile = os.Args[2]
	}

	// Read input file
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading file '%s': %v\n", inputFile, err)
	}

	// Prepare output writer
	var writer io.Writer
	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			log.Fatalf("Error creating output file '%s': %v\n", outputFile, err)
		}
		defer file.Close()
		writer = file
		fmt.Printf("Writing output to file: %s\n", outputFile)
	} else {
		writer = os.Stdout
	}

	// Process each byte and output ASCII codes
	processedCount := 0
	skippedCount := 0

	for _, b := range data {
		// Check if byte is within standard ASCII range (0-127)
		if b > 127 {
			skippedCount++
			continue
		}

		// Format character representation
		charRepr := formatCharacter(b)

		// Write output in the specified format
		fmt.Fprintf(writer, "Character %s -> Decimal: %d, Hex: 0x%02X\n", charRepr, b, b)
		processedCount++
	}

	// Print summary if output was written to file
	if outputFile != "" {
		fmt.Printf("\nProcessing complete!\n")
		fmt.Printf("Total ASCII characters processed: %d\n", processedCount)
		if skippedCount > 0 {
			fmt.Printf("Non-ASCII characters skipped: %d\n", skippedCount)
		}
	}
}

// formatCharacter returns a formatted string representation of a character
func formatCharacter(b byte) string {
	switch b {
	case '\n':
		return "'\\n'"
	case '\r':
		return "'\\r'"
	case '\t':
		return "'\\t'"
	case ' ':
		return "' '"
	case '\\':
		return "'\\\\'"
	case '\'':
		return "'\\''"
	default:
		// For other control characters (0-31, 127), show as hex
		if b < 32 || b == 127 {
			return fmt.Sprintf("'\\x%02X'", b)
		}
		// For printable ASCII characters
		return fmt.Sprintf("'%c'", b)
	}
}

// printUsage displays usage information
func printUsage() {
	fmt.Println("ASCII Code Converter")
	fmt.Println("====================")
	fmt.Println("\nUsage:")
	fmt.Println("  go run main.go <input_file> [output_file]")
	fmt.Println("\nArguments:")
	fmt.Println("  input_file   - Path to the text file to process (required)")
	fmt.Println("  output_file  - Path to save the output (optional, prints to console if not specified)")
	fmt.Println("\nExamples:")
	fmt.Println("  go run main.go sample.txt")
	fmt.Println("  go run main.go sample.txt output.txt")
	fmt.Println("\nOutput Format:")
	fmt.Println("  Character 'A' -> Decimal: 65, Hex: 0x41")
}

// Made with Bob
