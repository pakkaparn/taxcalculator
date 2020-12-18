package taxcal

// Calculate will return annual tax
func Calculate(income float64) float64 {
	deducted := deduct(income)

	netIncome := income - deducted

	if netIncome <= 150000 {
		return 0.00
	}

	return (netIncome - 150000) * 0.05
}

func deduct(income float64) float64 {
	deduct := 0.00
	deducts := make(map[string]float64)

	deducts["personal deduction"] = 60000.00

	for _, amount := range deducts {
		deduct += amount
	}

	return deduct
}
