package oracle

// import (
// 	"strconv"

// 	"github.com/ianhecker/correlate/internal/compute"
// )

// type OldTransaction struct {
// 	ID      string
// 	Date    string
// 	Time    float64
// 	C1      float64
// 	C2      float64
// 	C1ToUSD float64
// 	C2ToUSD float64
// }

// func (t *OldTransaction) Unmarshal(in []string) error {
// 	t.ID = in[0]
// 	t.Date = in[1]

// 	var err error
// 	t.Time, err = strconv.ParseFloat(in[2], 64)
// 	if err != nil {
// 		return err
// 	}

// 	t.C1, err = strconv.ParseFloat(in[3], 64)
// 	if err != nil {
// 		return err
// 	}

// 	t.C2, err = strconv.ParseFloat(in[4], 64)
// 	if err != nil {
// 		return err
// 	}

// 	t.C1ToUSD, err = strconv.ParseFloat(in[5], 64)
// 	if err != nil {
// 		return err
// 	}

// 	t.C2ToUSD, err = strconv.ParseFloat(in[6], 64)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// type OldTransactions struct {
// 	Time           compute.BasicStatistics
// 	C1             compute.BasicStatistics
// 	C2             compute.BasicStatistics
// 	CostOfC1InUSD  compute.BasicStatistics
// 	CostOfC2InUSD  compute.BasicStatistics
// 	TotalCostInUSD compute.BasicStatistics
// }

// func MakeOldTransactions(
// 	headers []string,
// 	txns ...OldTransaction,
// ) OldTransactions {
// 	time := make([]float64, len(txns))
// 	c1 := make([]float64, len(txns))
// 	c2 := make([]float64, len(txns))
// 	costOfC1InUSD := make([]float64, len(txns))
// 	costOfC2InUSD := make([]float64, len(txns))
// 	totalCostInUSD := make([]float64, len(txns))

// 	for i := 0; i < len(txns); i++ {
// 		time[i] = txns[i].Time
// 		c1[i] = txns[i].C1
// 		c2[i] = txns[i].C2
// 		costOfC1InUSD[i] = txns[i].C1 * txns[i].C1ToUSD
// 		costOfC2InUSD[i] = txns[i].C2 * txns[i].C2ToUSD
// 		totalCostInUSD[i] = costOfC1InUSD[i] + costOfC2InUSD[i]
// 	}

// 	return OldTransactions{
// 		Time:           compute.MakeBasicStatistics(headers[2], time...),
// 		C1:             compute.MakeBasicStatistics(headers[3], c1...),
// 		C2:             compute.MakeBasicStatistics(headers[4], c2...),
// 		CostOfC1InUSD:  compute.MakeBasicStatistics(headers[5], costOfC1InUSD...),
// 		CostOfC2InUSD:  compute.MakeBasicStatistics(headers[6], costOfC2InUSD...),
// 		TotalCostInUSD: compute.MakeBasicStatistics("Total Cost in USD", totalCostInUSD...),
// 	}
// }

// func (t OldTransactions) MakeStatistics() [][]string {
// 	return compute.MakeStatistics(
// 		[]string{"", "Max", "Mean", "Median", "Min", "Standard Deviation"},
// 		t.Time.Strings(),
// 		t.C1.Strings(),
// 		t.C2.Strings(),
// 		t.CostOfC1InUSD.Strings(),
// 		t.CostOfC2InUSD.Strings(),
// 		t.TotalCostInUSD.Strings(),
// 	)
// }
