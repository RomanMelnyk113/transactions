package tests

import (
	"juni/task/types"
	"time"
)

var transaction_0 = types.Transaction{
	ID:          "314706840032118839",
	Description: "Paypal *ny The New York Times Co.",
	Category:    "banks",
	Amount:      -2.1,
	MadeOn:      time.Date(2020, 9, 17, 0, 0, 0, 0, time.UTC),
}

var transaction_1 = types.Transaction{
	ID:          "314707597087213848",
	Description: "Credit Card",
	Category:    "payment_methods",
	Amount:      1.86,
	MadeOn:      time.Date(2020, 9, 17, 0, 0, 0, 0, time.UTC),
}
var transaction_2 = types.Transaction{
	ID:          "314707597087213849",
	Description: "To Euro",
	Category:    "payment_methods",
	Amount:      -1.86,
	MadeOn:      time.Date(2020, 9, 17, 0, 0, 0, 0, time.UTC),
}
var transaction_3 = types.Transaction{
	ID:          "314707596835555592",
	Description: "From British Pound",
	Category:    "payment_methods",
	Amount:      2,
	MadeOn:      time.Date(2020, 9, 17, 0, 0, 0, 0, time.UTC),
}
var transaction_4 = types.Transaction{
	ID:          "314707596835555593",
	Description: "The New York Times Co.",
	Category:    "payment_methods",
	Amount:      -2,
	MadeOn:      time.Date(2020, 9, 17, 0, 0, 0, 0, time.UTC),
}
