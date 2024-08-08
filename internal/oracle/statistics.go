package oracle

import (
	"fmt"

	"github.com/ianhecker/correlate/internal/compute"
	"github.com/ianhecker/correlate/internal/matrix"
)

type BasicStatistics struct {
	ID                string
	Mean              float64
	Median            float64
	Max               float64
	Min               float64
	StandardDeviation float64
}

func MakeBasicStatistics(
	id string,
	numbers ...float64,
) BasicStatistics {
	return BasicStatistics{
		ID:                id,
		Mean:              compute.Mean(numbers...),
		Median:            compute.Median(numbers...),
		Max:               compute.Max(numbers...),
		Min:               compute.Min(numbers...),
		StandardDeviation: compute.StandardDeviation(numbers...),
	}
}

func (b BasicStatistics) Strings() []string {
	return []string{
		b.ID,
		fmt.Sprintf("%f", b.Mean),
		fmt.Sprintf("%f", b.Median),
		fmt.Sprintf("%f", b.Max),
		fmt.Sprintf("%f", b.Min),
		fmt.Sprintf("%f", b.StandardDeviation),
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
