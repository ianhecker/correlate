package oracle

import (
	"fmt"
	"strconv"

	"github.com/ianhecker/correlate/internal/correlate"
)

type Transactions struct {
	Time          BasicStatistics
	C1            BasicStatistics
	C2            BasicStatistics
	CostOfC1InUSD BasicStatistics
	CostOfC2InUSD BasicStatistics
}

func MakeTransactions(txns ...Transaction) Transactions {
	time := make([]float64, len(txns))
	c1 := make([]float64, len(txns))
	c2 := make([]float64, len(txns))
	costOfC1InUSD := make([]float64, len(txns))
	costOfC2InUSD := make([]float64, len(txns))

	for i := 0; i < len(txns); i++ {
		time[i] = txns[i].Time
		c1[i] = txns[i].C1
		c2[i] = txns[i].C2
		costOfC1InUSD[i] = txns[i].C1 * txns[i].C1ToUSD
		costOfC2InUSD[i] = txns[i].C2 * txns[i].C2ToUSD
	}

	return Transactions{
		Time:          MakeBasicStatistics("Time", time...),
		C1:            MakeBasicStatistics("C1", c1...),
		C2:            MakeBasicStatistics("C2", c2...),
		CostOfC1InUSD: MakeBasicStatistics("Cost of C1 In USD", costOfC1InUSD...),
		CostOfC2InUSD: MakeBasicStatistics("Cost of C2 In USD", costOfC2InUSD...),
	}
}

func (t Transactions) TimeStatsToStrings() []string {
	return t.Time.Strings()
}

func (t Transactions) C1StatsToStrings() []string {
	return t.C1.Strings()
}

func (t Transactions) C2StatsToStrings() []string {
	return t.C2.Strings()
}

func (t Transactions) CostOfC1InUSDStatsToStrings() []string {
	return t.CostOfC1InUSD.Strings()
}

func (t Transactions) CostOfC2InUSDStatsToStrings() []string {
	return t.CostOfC2InUSD.Strings()
}

type Transaction struct {
	ID      string
	Date    string
	Time    float64
	C1      float64
	C2      float64
	C1ToUSD float64
	C2ToUSD float64
}

func (t *Transaction) Unmarshal(in []string) error {
	t.ID = in[0]
	t.Date = in[1]

	var err error
	t.Time, err = strconv.ParseFloat(in[2], 64)
	if err != nil {
		return err
	}

	t.C1, err = strconv.ParseFloat(in[3], 64)
	if err != nil {
		return err
	}

	t.C2, err = strconv.ParseFloat(in[4], 64)
	if err != nil {
		return err
	}

	t.C1ToUSD, err = strconv.ParseFloat(in[5], 64)
	if err != nil {
		return err
	}

	t.C2ToUSD, err = strconv.ParseFloat(in[6], 64)
	if err != nil {
		return err
	}
	return nil
}

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
		Mean:              correlate.Mean(numbers...),
		Median:            correlate.Median(numbers...),
		Max:               correlate.Max(numbers...),
		Min:               correlate.Min(numbers...),
		StandardDeviation: correlate.StandardDeviation(numbers...),
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

func ParseCSV(data [][]string) (Transactions, error) {
	if len(data) < 2 {
		return Transactions{}, fmt.Errorf("Data is not long enough")
	}

	var txns []Transaction

	for i := 1; i < len(data); i++ {
		var txn Transaction
		err := txn.Unmarshal(data[i])

		if err != nil {
			return Transactions{}, err
		}
		txns = append(txns, txn)
	}

	return MakeTransactions(txns...), nil
}
