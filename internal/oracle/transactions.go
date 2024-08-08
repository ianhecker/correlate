package oracle

type Transactions struct {
	Time          BasicStatistics
	C1            BasicStatistics
	C2            BasicStatistics
	CostOfC1InUSD BasicStatistics
	CostOfC2InUSD BasicStatistics
}

func MakeTransactions(
	headers []string,
	txns ...Transaction,
) Transactions {
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
		Time:          MakeBasicStatistics(headers[2], time...),
		C1:            MakeBasicStatistics(headers[3], c1...),
		C2:            MakeBasicStatistics(headers[4], c2...),
		CostOfC1InUSD: MakeBasicStatistics(headers[5], costOfC1InUSD...),
		CostOfC2InUSD: MakeBasicStatistics(headers[6], costOfC2InUSD...),
	}
}

func (t Transactions) Statistics() [][]string {
	return MakeStatistics(
		[]string{"", "Mean", "Median", "Max", "Min", "Standard Deviation"},
		t.Time.Strings(),
		t.C1.Strings(),
		t.C2.Strings(),
		t.CostOfC1InUSD.Strings(),
		t.CostOfC2InUSD.Strings(),
	)
}
