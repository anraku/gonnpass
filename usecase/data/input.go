package data

type InputData struct {
	KeywordAND []string
	KeywordOR  []string
	Order      Order
	Count      int
}

type Order int

const (
	UpdateOrder Order = iota + 1
	StartOrder
	NewOrder
)
