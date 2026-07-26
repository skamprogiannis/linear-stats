package main

import (
	"math"
	"sort"
)

// stats contains rounded descriptive statistics for a set of integers.
type stats struct {
	average           int
	median            int
	variance          int
	standardDeviation int
}

// calculateStats returns rounded descriptive statistics for nums.
func calculateStats(nums []int) stats {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	average := float64(sum) / float64(len(nums))

	sorted := append([]int(nil), nums...)
	sort.Ints(sorted)
	mid := len(sorted) / 2
	median := float64(sorted[mid])
	if len(sorted)%2 == 0 {
		median = float64(sorted[mid-1]+sorted[mid]) / 2
	}

	squaredDeviationSum := 0.0
	for _, n := range nums {
		deviation := float64(n) - average
		squaredDeviationSum += deviation * deviation
	}
	variance := squaredDeviationSum / float64(len(nums))

	return stats{
		average:           int(math.Round(average)),
		median:            int(math.Round(median)),
		variance:          int(math.Round(variance)),
		standardDeviation: int(math.Round(math.Sqrt(variance))),
	}
}

func getPearsonCorrelationCoefficient(values []int) float64 {
	if len(values) < 2 {
		return math.NaN()
	}

	var sumX, sumY float64
	for index, value := range values {
		sumX += float64(index)
		sumY += float64(value)
	}

	meanX := sumX / float64(len(values))
	meanY := sumY / float64(len(values))

	var covariance, xDeviationSquared, yDeviationSquared float64
	for index, value := range values {
		xDeviation := float64(index) - meanX
		yDeviation := float64(value) - meanY

		covariance += xDeviation * yDeviation
		xDeviationSquared += xDeviation * xDeviation
		yDeviationSquared += yDeviation * yDeviation
	}

	denominator := math.Sqrt(xDeviationSquared * yDeviationSquared)
	if denominator == 0 {
		return math.NaN()
	}

	return covariance / denominator
}
