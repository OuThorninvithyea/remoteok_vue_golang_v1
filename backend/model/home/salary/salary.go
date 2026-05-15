package salary

type SalaryConfig struct {
	Min            int    `json:"min"`
	Max            int    `json:"max"`
	Step           int    `json:"step"`
	Value          int    `json:"value"`
	CurrencyPrefix string `json:"currencyPrefix"`
	CurrencySuffix string `json:"currencySuffix"`
}
