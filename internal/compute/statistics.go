package compute

import (
	"fmt"

	"github.com/ianhecker/correlate/internal/matrix"
)

type BasicStatistics struct {
	ID                string
	IsUSD             bool
	Max               float64
	Mean              float64
	Median            float64
	Min               float64
	StandardDeviation float64
}

func MakeBasicStatistics(
	id string,
	isUSD bool,
	numbers ...float64,
) BasicStatistics {
	return BasicStatistics{
		ID:                id,
		IsUSD:             isUSD,
		Max:               Max(numbers...),
		Mean:              Mean(numbers...),
		Median:            Median(numbers...),
		Min:               Min(numbers...),
		StandardDeviation: StandardDeviation(numbers...),
	}
}

func (b *BasicStatistics) Strings() []string {
	var format = ""

	if b.IsUSD {
		format = "$"
	}

	return []string{
		b.ID,
		fmt.Sprintf(format+"%.10f", b.Max),
		fmt.Sprintf(format+"%.10f", b.Mean),
		fmt.Sprintf(format+"%.10f", b.Median),
		fmt.Sprintf(format+"%.10f", b.Min),
		fmt.Sprintf(format+"%.10f", b.StandardDeviation),
	}
}

type Statistics [][]string

func MakeStatistics(
	headers []string,
	data ...[]string,
) Statistics {
	var stats = [][]string{}

	stats = append(stats, headers)
	stats = append(stats, data...)

	return matrix.TransposeString(stats)
}
