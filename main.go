package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// main reads numbers from standard input and prints a line for the Linear Regression Line and
// one for the Pearson Correlation Coefficient.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var history []int

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		history = append(history, num)

		lsrl := getLinearRegressionLine(history)
		pcc := getPearsonCorrelationCoefficient(history)
		fmt.Printf("Linear Regression Line: y = %dx + %d\n", lsrl, lsrl)
		fmt.Printf("Pearson Correlation Coefficient: %d", pcc)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
