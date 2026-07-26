package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func runProgram(t *testing.T, input string) (string, error) {
	t.Helper()

	path := filepath.Join(t.TempDir(), "data.txt")
	if err := os.WriteFile(path, []byte(input), 0o600); err != nil {
		t.Fatal(err)
	}

	return runProgramAtPath(t, path)
}

func runProgramAtPath(t *testing.T, path string) (string, error) {
	t.Helper()

	command := exec.Command("go", "run", ".", path)
	output, err := command.CombinedOutput()
	return string(output), err
}

func TestProgramPrintsRequiredStatistics(t *testing.T) {
	output, err := runProgram(t, "1\n2\n3\n")
	if err != nil {
		t.Fatalf("program failed: %v\n%s", err, output)
	}

	const expected = "Linear Regression Line: y = 1.000000x + 1.000000\n" +
		"Pearson Correlation Coefficient: 1.0000000000\n"
	if output != expected {
		t.Errorf("output = %q, want %q", output, expected)
	}
}

func TestProgramPrintsNonPerfectStatistics(t *testing.T) {
	output, err := runProgram(t, "1\n2\n4\n")
	if err != nil {
		t.Fatalf("program failed: %v\n%s", err, output)
	}

	const expected = "Linear Regression Line: y = 1.500000x + 0.833333\n" +
		"Pearson Correlation Coefficient: 0.9819805061\n"
	if output != expected {
		t.Errorf("output = %q, want %q", output, expected)
	}
}

func TestProgramSkipsBlankAndNonIntegerLines(t *testing.T) {
	output, err := runProgram(t, "1\nignored\n\n2\n3\n")
	if err != nil {
		t.Fatalf("program failed: %v\n%s", err, output)
	}

	const expected = "Linear Regression Line: y = 1.000000x + 1.000000\n" +
		"Pearson Correlation Coefficient: 1.0000000000\n"
	if output != expected {
		t.Errorf("output = %q, want %q", output, expected)
	}
}

func TestProgramRejectsUndefinedStatistics(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "fewer than two values",
			input: "1\n",
			want:  "at least two values are required",
		},
		{
			name:  "constant values",
			input: "5\n5\n5\n",
			want:  "Pearson correlation coefficient is undefined for constant values",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output, err := runProgram(t, test.input)
			if err == nil {
				t.Fatalf("program succeeded with output: %s", output)
			}
			if !strings.Contains(output, test.want) {
				t.Errorf("output = %q, want error containing %q", output, test.want)
			}
		})
	}
}

func TestProgramRejectsMissingInputFile(t *testing.T) {
	output, err := runProgramAtPath(t, filepath.Join(t.TempDir(), "missing.txt"))
	if err == nil {
		t.Fatalf("program succeeded with output: %s", output)
	}
	if !strings.Contains(output, "open input file") {
		t.Errorf("output = %q, want input-file error", output)
	}
}
