package stork

import (
	"fmt"

	"github.com/ianhecker/correlate/internal/oracle"
)

type Stork struct{}

func (c Stork) ParseMatrixIntoTransactions(data [][]string) (oracle.Transactions, error) {
	if len(data) < 8 {
		return Txns{}, fmt.Errorf("Data is not long enough")
	}

	var headers = data[0]
	var txns []Txn

	for i := 1; i < len(data); i++ {

		var txn Txn

		err := txn.Unmarshal(data[i])
		if err != nil {
			return Txns{}, err
		}
		txns = append(txns, txn)
	}

	return MakeTxns(headers, txns...), nil
}

func (c Stork) Template() [][]string {
	return [][]string{}
}
