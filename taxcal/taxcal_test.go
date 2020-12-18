package taxcal

import "testing"

func TestCalculateWithTwoHundredAndTenGrandToReturnZero(t *testing.T) {
	tax := Calculate(210000)

	if tax != 0 {
		t.Error("Calculate(210000) should return 0 but return", tax)
	}
}
func TestCalculateWithTwoHundredAndTwentyGrandToReturnZero(t *testing.T) {
	tax := Calculate(220000.00)

	if tax != 500 {
		t.Error("Calculate(220000) should return 500 but return", tax)
	}
}
func TestDeductWithOneHundredFifttyGrandToReturnSixtyGrand(t *testing.T) {
	deducted := deduct(150000.00)

	if deducted != 60000.0 {
		t.Error("deduct(150000.00) should return 60000.00 but return", deducted)
	}
}
