package main

import (
	"errors"
	"math"
)

var (
	errInsufficientValues   = errors.New("at least two values are required")
	errUndefinedCorrelation = errors.New("Pearson correlation coefficient is undefined for constant values")
)

type regressionSums struct {
	meanX             float64
	meanY             float64
	covariance        float64
	xDeviationSquared float64
	yDeviationSquared float64
}

func getLinearRegressionLine(values []int) (slope, intercept float64, err error) {
	sums, err := calculateRegressionSums(values)
	if err != nil {
		return 0, 0, err
	}

	slope = sums.covariance / sums.xDeviationSquared
	intercept = sums.meanY - slope*sums.meanX
	return slope, intercept, nil
}

func getPearsonCorrelationCoefficient(values []int) (float64, error) {
	sums, err := calculateRegressionSums(values)
	if err != nil {
		return 0, err
	}

	denominator := math.Sqrt(sums.xDeviationSquared * sums.yDeviationSquared)
	if denominator == 0 {
		return 0, errUndefinedCorrelation
	}

	return sums.covariance / denominator, nil
}

func calculateRegressionSums(values []int) (regressionSums, error) {
	if len(values) < 2 {
		return regressionSums{}, errInsufficientValues
	}

	var sumX, sumY float64
	for index, value := range values {
		sumX += float64(index)
		sumY += float64(value)
	}

	sums := regressionSums{
		meanX: sumX / float64(len(values)),
		meanY: sumY / float64(len(values)),
	}
	for index, value := range values {
		xDeviation := float64(index) - sums.meanX
		yDeviation := float64(value) - sums.meanY

		sums.covariance += xDeviation * yDeviation
		sums.xDeviationSquared += xDeviation * xDeviation
		sums.yDeviationSquared += yDeviation * yDeviation
	}

	return sums, nil
}
