package oracle

import (
	"fmt"
)

type Transaction struct {
	ID     string
	Date   string
	Time   string
	C1     string
	C2     string
	C1_USD string
	C2_USD string
}

func (d *Transaction) Unmarshal(array []string) error {
	var length = 7

	if len(array) != length {
		return fmt.Errorf("Length of array is not %d: %+v", length, array)
	}

	var i = 0
	d.ID = array[i]
	i++
	d.Date = array[i]
	i++
	d.Time = array[i]
	i++
	d.C1 = array[i]
	i++
	d.C2 = array[i]
	i++
	d.C1_USD = array[i]
	i++
	d.C2_USD = array[i]

	return nil
}

func ParseCSV(data [][]string) ([]Transaction, error) {
	if len(data) < 2 {
		return nil, fmt.Errorf("Data is not long enough")
	}

	txns := make([]Transaction, len(data)-1)

	for i := 0; i < len(data)-1; i++ {
		err := txns[i].Unmarshal(data[i+1])

		if err != nil {
			return nil, err
		}
	}
	return txns, nil
}
