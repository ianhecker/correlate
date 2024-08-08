package oracle

import "strconv"

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
