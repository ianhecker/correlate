package oracle

type Transactions interface {
	MakeStatistics() [][]string
}

type Oracle interface {
	ParseMatrixIntoTransactions([][]string) (Transactions, error)
	Template() [][]string
}
