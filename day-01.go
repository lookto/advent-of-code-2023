package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the file
	file, err := os.Open("day-01-input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Initialize a  2D slice to store the characters
	var matrix [][]interface{}

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			fmt.Println("error reading line")
			break
		}

		// Split the line into characters and store them in the matrix
		var row []interface{}
		for _, char := range line {
			// Convert int32 to rune for correct character representation
			row = append(row, rune(char))
		}
		matrix = append(matrix, row)
	}

	// Print the matrix
	for _, row := range matrix {
		for _, char := range row {
			fmt.Printf("%c ", char)
		}
		fmt.Println()
	}

}
