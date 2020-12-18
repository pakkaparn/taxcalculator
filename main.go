package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pakkaparn/taxcalculator/taxcal"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Tax Calcualtor")
	fmt.Println("--------------")

	for {
		fmt.Println("What is you annual income?")
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		income, err := strconv.ParseFloat(text, 2)

		if err != nil {
			fmt.Printf("Sorry, '%v' is not a number", text)
			fmt.Println("")
			continue
		}

		tax := taxcal.Calculate(income)

		fmt.Printf("You have to pay %.2f", tax)
		fmt.Println("")
		fmt.Println("--------------")
	}
}
