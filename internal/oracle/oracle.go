package oracle

import (
	"fmt"
	"strconv"
)

type Transactions struct {
	ID     []string
	Date   []string
	Time   []float64
	C1     []float64
	C2     []float64
	C1_USD []float64
	C2_USD []float64
}

func (t *Transactions) Unmarshal(array [][]string) error {
	transposed := transpose(array)

	t.ID = transposed[0]
	t.Date = transposed[1]

	var err error
	t.Time, err = t.ParseStringToFloat64(transposed[2])
	if err != nil {
		return err
	}

	t.C1, err = t.ParseStringToFloat64(transposed[3])
	if err != nil {
		return err
	}

	t.C2, err = t.ParseStringToFloat64(transposed[4])
	if err != nil {
		return err
	}

	t.C1_USD, err = t.ParseStringToFloat64(transposed[5])
	if err != nil {
		return err
	}

	t.C2_USD, err = t.ParseStringToFloat64(transposed[6])
	if err != nil {
		return err
	}
	return nil
}

func (t Transactions) CostOfC1InUSD() []float64 {
	var costOfC1InUSD = make([]float64, len(t.C1))

	for i := 0; i < len(t.C1); i++ {
		costOfC1InUSD[i] = t.C1[i] * t.C1_USD[i]
	}

	return costOfC1InUSD
}

func (t Transactions) CostOfC2InUSD() []float64 {
	var costOfC2InUSD = make([]float64, len(t.C2))

	for i := 0; i < len(t.C2); i++ {
		costOfC2InUSD[i] = t.C2[i] * t.C2_USD[i]
	}

	return costOfC2InUSD
}

func transpose(a [][]string) [][]string {
	newArr := make([][]string, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			newArr[j] = append(newArr[j], a[i][j])
		}
	}
	return newArr
}

func (d *Transactions) ParseStringToFloat64(in []string) ([]float64, error) {
	var out = make([]float64, len(in))

	for i := 0; i < len(in); i++ {
		f, err := strconv.ParseFloat(in[i], 64)

		if err != nil {
			return nil, err
		}
		out[i] = f
	}
	return out, nil
}

func ParseCSV(data [][]string) (Transactions, error) {
	if len(data) < 2 {
		return Transactions{}, fmt.Errorf("Data is not long enough")
	}

	var txns Transactions
	err := txns.Unmarshal(data[1:])
	if err != nil {
		return Transactions{}, err
	}
	return txns, nil
}
