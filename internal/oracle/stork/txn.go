package stork

import "strconv"

type Txn struct {
	ID             string
	Date           string
	Duration       float64
	GasUsed        float64
	GasPriceInGwei float64
	GasFeeInGwei   float64
	GasFeeInEther  float64
	EtherToUSD     float64
}

func (t *Txn) Unmarshal(in []string) error {
	t.ID = in[0]
	t.Date = in[1]

	var err error
	t.Duration, err = strconv.ParseFloat(in[2], 64)
	if err != nil {
		return err
	}

	t.GasUsed, err = strconv.ParseFloat(in[3], 64)
	if err != nil {
		return err
	}

	t.GasPriceInGwei, err = strconv.ParseFloat(in[4], 64)
	if err != nil {
		return err
	}

	t.GasFeeInGwei, err = strconv.ParseFloat(in[5], 64)
	if err != nil {
		return err
	}

	t.GasFeeInEther, err = strconv.ParseFloat(in[6], 64)
	if err != nil {
		return err
	}

	t.EtherToUSD, err = strconv.ParseFloat(in[7], 64)
	if err != nil {
		return err
	}
	return nil
}
