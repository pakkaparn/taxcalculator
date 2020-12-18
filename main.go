package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pakkaparn/taxcalculator/taxcal"
)

// Reader struct
type Reader struct {
	reader *bufio.Reader
}

// IO interface
type IO interface {
	getInput() string
}

func (r Reader) getInput() string {
	text, _ := r.reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	return text
}

func main() {
	r := Reader{}

	r.reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Tax Calcualtor")
		fmt.Println("--------------")

		income := askForAnnualIncome(r)
		tax := taxcal.Calculate(income)

		fmt.Printf("You have to pay %.2f", tax)
		fmt.Println("")
		fmt.Println("--------------")
	}
}

func askForAnnualIncome(r IO) float64 {
	fmt.Println("What is you annual income?")
	fmt.Print("-> ")

	text := r.getInput()
	income, err := strconv.ParseFloat(text, 2)

	if err != nil {
		fmt.Printf("Sorry, '%v' is not a number", text)
		fmt.Println("")
		return askForAnnualIncome(r)
	}

	return income
}
