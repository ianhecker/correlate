package oracle

import (
	"fmt"
)

func ParseMatrixIntoTransactions(data [][]string) (Transactions, error) {
	if len(data) < 2 {
		return Transactions{}, fmt.Errorf("Data is not long enough")
	}

	var headers = data[0]
	var txns []Transaction

	for i := 1; i < len(data); i++ {

		var txn Transaction

		err := txn.Unmarshal(data[i])
		if err != nil {
			return Transactions{}, err
		}
		txns = append(txns, txn)
	}

	return MakeTransactions(headers, txns...), nil
}
