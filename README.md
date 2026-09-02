# Linear Stats

A small Go command-line program that derives a linear regression line and the
Pearson correlation coefficient from a one-dimensional dataset.

For each accepted value, its zero-based position is `x` and the value itself
is `y`. Blank and non-integer lines are skipped before those positions are
assigned.

## Run

Provide a text file containing one integer per line:

```bash
go run . examples/data.txt
```

Output for the included `1, 2, 4` dataset:

```text
Linear Regression Line: y = 1.500000x + 0.833333
Pearson Correlation Coefficient: 0.9819805061
```

The program exits with an error when the file cannot be read, contains fewer
than two accepted values, or has constant values for which correlation is
undefined.

## How it works

`calculateRegressionSums` makes two passes over the values. The first finds the
means of `x` and `y`; the second accumulates covariance and squared deviations.
Those shared sums feed both results:

```text
slope     = covariance(x, y) / deviation²(x)
intercept = mean(y) - slope * mean(x)
Pearson r = covariance(x, y) / sqrt(deviation²(x) * deviation²(y))
```

The implementation uses Go's standard library only.

## Test

```bash
go test ./...
go test -race ./...
go vet ./...
```

The test suite covers exact CLI output, a non-perfect correlation, ignored
input lines, insufficient and constant data, and missing files.

## Project structure

- `main.go` — file parsing, CLI validation, and output formatting.
- `statistics.go` — regression and correlation calculations.
- `main_test.go` — end-to-end command tests.
- `examples/data.txt` — minimal reproducible input.

## Author

`skamprogiannis` designed and implemented the parser, statistics, CLI, tests,
and documentation. The curated public history removes an unrelated project
snapshot that preceded this code while preserving the original metadata of
the substantive Linear Stats commits.

## Status

This Zone01 Athens exercise implements the required statistics and explicitly
handles undefined inputs. It is intentionally a file-oriented CLI rather than
a reusable statistics package.
