package transactiontypes

import "fmt"

type Transaction string

const (
	TransactionBuy  Transaction = "buy"
	TransactionSell Transaction = "sell"
)

func (s Transaction) String() string {
	return string(s)
}

func ParseTransaction(s string) (Transaction, error) {
	switch s {
	case string(TransactionBuy):
		return TransactionBuy, nil
	case string(TransactionSell):
		return TransactionSell, nil
	default:
		return "", fmt.Errorf("invalid transaction type: %s", s)
	}
}
