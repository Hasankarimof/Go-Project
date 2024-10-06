package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

// Read a balance file
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to read file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}

// Write a balance to file
func WriteBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {

	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	fmt.Println("Welcome to Go Bank!")

	for {

		fmt.Println("What do  you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		fmt.Println("Your Choice:", choice)

		if choice == 1 {
			fmt.Println("Your balance is ", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			WriteBalanceToFile(accountBalance)
		} else if choice == 3 {
			fmt.Print("How much do you want to Withdraw money? ")
			var WithdrawAmount float64
			fmt.Scan(&WithdrawAmount)
			if WithdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				continue
			} else if WithdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance -= WithdrawAmount
			fmt.Println("Balance updated! Your Balance amount:", accountBalance)
			WriteBalanceToFile(accountBalance)
		} else {
			fmt.Println("Thank you !")
			break
		}

	}

}
