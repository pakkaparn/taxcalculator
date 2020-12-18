package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pakkaparn/taxcalculator/taxcal"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Tax Calcualtor")
		fmt.Println("--------------")

		income := askForAnnualIncome()
		tax := taxcal.Calculate(income)

		fmt.Printf("You have to pay %.2f", tax)
		fmt.Println("")
		fmt.Println("--------------")
	}
}

func askForAnnualIncome() float64 {
	fmt.Println("What is you annual income?")
	fmt.Print("-> ")

	text := input()
	income, err := strconv.ParseFloat(text, 2)

	if err != nil {
		fmt.Printf("Sorry, '%v' is not a number", text)
		fmt.Println("")
		return askForAnnualIncome()
	}

	return income
}

func input() string {
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)

	return text
}
