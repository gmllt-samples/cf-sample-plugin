package utils

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// HandleError standardize error messages and exit the program
func HandleError(err error) {
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}

// PrintTable formats and prints a table to the console
func PrintTable(headers []string, rows [][]string) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	// Print the headers
	for _, header := range headers {
		_, err := fmt.Fprintf(writer, "%s\t", header)
		if err != nil {
			return
		}
	}
	_, err := fmt.Fprintln(writer)
	if err != nil {
		return
	}

	// Print the rows
	for _, row := range rows {
		for _, col := range row {
			_, err := fmt.Fprintf(writer, "%s\t", col)
			if err != nil {
				return
			}
		}
		_, err := fmt.Fprintln(writer)
		if err != nil {
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		return
	}
}

// ValidateArgs checks if the right number of arguments are passed
func ValidateArgs(args []string, expected int) {
	if len(args) < expected {
		_, err := fmt.Fprintf(os.Stderr, "Error: Expected at least %d arguments, but got %d\n", expected, len(args))
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
