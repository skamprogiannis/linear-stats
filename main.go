package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go run . <data-file>")
		os.Exit(1)
	}

	values, err := readValues(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	slope, intercept, err := getLinearRegressionLine(values)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	correlation, err := getPearsonCorrelationCoefficient(values)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", correlation)
}

func readValues(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open input file: %w", err)
	}
	defer file.Close()

	var values []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		value, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		values = append(values, value)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}

	return values, nil
}
