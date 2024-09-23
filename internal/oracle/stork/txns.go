package stork

import (
	"github.com/ianhecker/correlate/internal/compute"
)

type Txns struct {
	Duration       compute.BasicStatistics
	GasUsed        compute.BasicStatistics
	GasPriceInGwei compute.BasicStatistics
	GasFeeInGwei   compute.BasicStatistics
	GasFeeInEther  compute.BasicStatistics
	EtherToUSD     compute.BasicStatistics
	CostInUSD      compute.BasicStatistics
}

func MakeTxns(
	headers []string,
	txns ...Txn,
) Txns {
	duration := make([]float64, len(txns))
	gasUsed := make([]float64, len(txns))
	gasPriceInGwei := make([]float64, len(txns))
	gasFeeInGwei := make([]float64, len(txns))
	gasFeeInEther := make([]float64, len(txns))
	etherToUSD := make([]float64, len(txns))
	costInUSD := make([]float64, len(txns))

	for i := 0; i < len(txns); i++ {
		duration[i] = txns[i].Duration
		gasUsed[i] = txns[i].GasUsed
		gasPriceInGwei[i] = txns[i].GasPriceInGwei
		gasFeeInGwei[i] = txns[i].GasFeeInGwei
		gasFeeInEther[i] = txns[i].GasFeeInEther
		etherToUSD[i] = txns[i].EtherToUSD
		costInUSD[i] = gasFeeInEther[i] * etherToUSD[i]
	}

	return Txns{
		Duration:       compute.MakeBasicStatistics("DURATION", duration...),
		GasUsed:        compute.MakeBasicStatistics("GAS USED", gasUsed...),
		GasPriceInGwei: compute.MakeBasicStatistics("GAS PRICE IN GWEI", gasPriceInGwei...),
		GasFeeInGwei:   compute.MakeBasicStatistics("GAS FEE IN GWEI", gasFeeInGwei...),
		GasFeeInEther:  compute.MakeBasicStatistics("GAS FEE IN ETHER", gasFeeInEther...),
		EtherToUSD:     compute.MakeBasicStatistics("ETHER TO USD", etherToUSD...),
		CostInUSD:      compute.MakeBasicStatistics("COST IN USD", costInUSD...),
	}
}

func (t Txns) MakeStatistics() [][]string {
	return compute.MakeStatistics(
		[]string{"", "Max", "Mean", "Median", "Min", "Standard Deviation"},
		t.Duration.Strings(),
		t.GasUsed.Strings(),
		t.GasPriceInGwei.Strings(),
		t.GasFeeInGwei.Strings(),
		t.GasFeeInEther.Strings(),
		t.EtherToUSD.Strings(),
		t.CostInUSD.Strings(),
	)
}
