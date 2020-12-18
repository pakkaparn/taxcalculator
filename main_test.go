package main

import (
	"bufio"
	"bytes"
	"testing"
)

var getInputMock func() string

type IOMock struct {
	inputCount int
}

func (io IOMock) getInput() string {
	return getInputMock()
}

func TestGetInputShouldReturnHello(t *testing.T) {
	var stdin bytes.Buffer
	stdin.Write([]byte("Hello\n"))

	reader := Reader{reader: bufio.NewReader(&stdin)}

	message := reader.getInput()

	if message != "Hello" {
		t.Error("getInput() should return Hello but return", message)
	}
}

func TestAskForAnnualIncomeWithTwoHundredGrandShouldReturnTwoHundredGrand(t *testing.T) {
	io := IOMock{}
	getInputMock = func() string {
		return "200000"
	}

	amount := askForAnnualIncome(io)

	if amount != 200000.00 {
		t.Error("askForAnnualIncome(io) should return 200000.00 but return", amount)
	}
}

func TestAskForAnnualIncomeWithStringShouldCallRecursiveFunction(t *testing.T) {
	io := IOMock{inputCount: 0}

	getInputMock = func() string {
		if io.inputCount < 2 {
			io.inputCount++

			return ""
		}

		return "200000"
	}

	amount := askForAnnualIncome(io)

	if io.inputCount < 2 {
		t.Error("askForAnnualIncome(io) should call getInput 2 times but call only ", amount)
	}

	if amount != 200000.00 {
		t.Error("askForAnnualIncome(io) should return 200000.00 but return", amount)
	}
}
