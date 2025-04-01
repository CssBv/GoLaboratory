package main

import (
	"fmt"
	"os"
	"strconv"
)
const accountBalance = "balance.txt"

func getBalanceFromFile() float64{
	data , _ := os.ReadFile(accountBalance)
	balanceText := string(data)
	balance , _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func main(){

	fmt.Print("Bank App\n")
	var totalBalance = getBalanceFromFile()
	
	for {
		fmt.Print("Choose an option:\n")
		fmt.Print("1. Deposit\n")
		fmt.Print("2. Withdraw\n")
		fmt.Print("3. Check Balance\n")
		fmt.Print("4. Exit\n")
		choice := getUserInput("Enter your choice:")
		
		switch choice {
		case 1:
			var depositAmount float64
			fmt.Print("Enter amount to deposit:\n")
			fmt.Scan(&depositAmount)
			
			if depositAmount < 0 {
				fmt.Print("\nInvalid Amount. Please try again.\n\n")
				continue
			}
			totalBalance += depositAmount
			writeBalanceToFile(totalBalance)
			fmt.Print("Total Balance:", totalBalance, "\n")	
		case 2:
			var withdrawAmount float64
			fmt.Print("Enter a mount to withdraw:\n")
			fmt.Scan(&withdrawAmount)
			
			if withdrawAmount > totalBalance {
				fmt.Print("\nInvalid Amount. Please try again.\n\n")
				continue
			}
			totalBalance -= withdrawAmount
			fmt.Print("Total Balance:", totalBalance, "\n")
		case 3:
			fmt.Print("Your balance:\n")
			fmt.Print("Total Balance:", totalBalance, "\n")
		default:
			fmt.Print("Exiting...\n")
			fmt.Print("Thanks for choosing our bank!\n")
			return	
		}
	}
}

func getUserInput(promptText string) float64 {
	var userInput float64
	fmt.Println(promptText)
	fmt.Scan(&userInput)
	return userInput
}

func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(("balance.txt"), []byte(balanceText), 0644)
}

