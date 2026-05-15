package store

import "backend/model/home/salary"

var DropdownSalaryData = salary.SalaryConfig{
	Min:            0,
	Max:            300,
	Step:           1,
	Value:          0,
	CurrencyPrefix: "$",
	CurrencySuffix: "k/year",
}
