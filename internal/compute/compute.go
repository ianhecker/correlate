package compute

import (
	"math"
	"sort"
)

func Mean(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))

}

func Median(numbers ...float64) float64 {
	numbers2 := make([]float64, len(numbers))
	copy(numbers2, numbers)

	sort.Float64s(numbers2)

	var median float64
	length := len(numbers2)
	if length == 0 {
		return 0
	} else if length%2 == 0 {
		median = (numbers2[length/2-1] + numbers2[length/2]) / 2
	} else {
		median = numbers2[length/2]
	}

	return median
}

func Min(numbers ...float64) float64 {
	numbers2 := make([]float64, len(numbers))
	copy(numbers2, numbers)

	sort.Float64s(numbers2)

	return numbers2[0]
}

func Max(numbers ...float64) float64 {
	numbers2 := make([]float64, len(numbers))
	copy(numbers2, numbers)

	sort.Float64s(numbers2)

	return numbers2[len(numbers2)-1]
}

func StandardDeviation(numbers ...float64) float64 {
	mean := Mean(numbers...)

	var sum float64
	for _, num := range numbers {
		sum += math.Pow(num-mean, 2)
	}

	variance := sum / float64(len(numbers))
	return math.Sqrt(variance)
}
