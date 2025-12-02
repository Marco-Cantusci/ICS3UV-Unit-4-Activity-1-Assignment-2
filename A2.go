/**
 * @author Marco Cantusci
 * @verion 1.0.0
 * @date 2025-12-02
 * @fileoverview Right angle triangle made out of numbers.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	// declare variables
	var rowsString string
	var rowsFloat float64

	reader := bufio.NewReader(os.Stdin)

	// input
	fmt.Print("How many rows would you like? ")
	rowsString, _ = reader.ReadString('\n')
	rowsString = strings.TrimSpace(rowsString)
	rowsFloat, _ = strconv.ParseFloat(rowsString, 64)
	var rowsNumber int = int(rowsFloat)

	// process
	for row := 1; row <= rowsNumber; row++ {
		for numbers := 1; numbers <= row; numbers++ {
			fmt.Print(numbers)
		}
		fmt.Printf("\n") // output
	}

	fmt.Println("\nDone.")
}
