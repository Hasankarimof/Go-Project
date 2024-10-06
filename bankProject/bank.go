package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {

		presentOperations()

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
			fileops.WriteValueToFile(accountBalanceFile, accountBalance)
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
			fileops.WriteValueToFile(accountBalanceFile, accountBalance)
		} else {
			fmt.Println("Thank you !")
			break
		}

	}

}
