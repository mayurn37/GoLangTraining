package main

import "fmt"

/* Note : We can use one struct with member 'isAccount' to show if that is a current acc or saving acc,
but just to show inheritance, I have taken two structs */

//Account ...
type Account interface {
	depositMoney(float32)
	checkBalance()
	withdrawMoney(float32)
	getInterest()
}

//SavingsAccount ...
type SavingsAccount struct {
	accountNo    int
	interestRate float32
	balance      float32
}

//CurrentAccount ...
type CurrentAccount struct {
	accountNo int
	balance   float32
}

func (currentAcc *CurrentAccount) checkBalance() {
	fmt.Printf("Balance of Current  account is: Rs. %.2f\n", currentAcc.balance)
}
func (currentAcc *CurrentAccount) withdrawMoney(amount float32) {

	if currentAcc.balance-amount < 1000 {
		fmt.Printf("Cannot Withdraw Rs. %.2f. Minimum Blanace Rs. 1000 is required.\n", amount)
	} else {
		currentAcc.balance = currentAcc.balance - amount
		fmt.Printf("Amount Withdrawn: Rs. %.2f, Avalable balance: Rs. %.2f\n", amount, currentAcc.balance)
	}
}
func (currentAcc *CurrentAccount) getInterest() {
	fmt.Printf("current account: No interest\n")
}

func (currentAcc *CurrentAccount) depositMoney(amount float32) {
	currentAcc.balance = currentAcc.balance + amount
	fmt.Printf("Amount Deposited: Rs. %.2f, Available balance:Rs. %.2f\n", amount, currentAcc.balance)
}

func (savingsAcc *SavingsAccount) checkBalance() {
	fmt.Printf("Balance of Savings  account is: %.2f\n", savingsAcc.balance)
}
func (savingsAcc *SavingsAccount) withdrawMoney(amount float32) {
	if savingsAcc.balance-amount < 1000 {
		fmt.Printf("Cannot Withdraw Rs. %.2f. Minimum Blanace Rs. 1000 is required.\n", amount)
	} else {
		savingsAcc.balance = savingsAcc.balance - amount
		fmt.Printf("Amount Withdrawn: Rs. %.2f, Avalable balance: Rs. %.2f\n", amount, savingsAcc.balance)
	}
}
func (savingsAcc *SavingsAccount) getInterest() {
	fmt.Printf("Interest on Rs. %.2f is Rs. %.2f\n", savingsAcc.balance, savingsAcc.balance*savingsAcc.interestRate/100)
}

func (savingsAcc *SavingsAccount) depositMoney(amount float32) {
	savingsAcc.balance = savingsAcc.balance + amount
	fmt.Printf("Amount Deposited: Rs. %.2f, Available balance:Rs. %.2f\n", amount, savingsAcc.balance)
}

func main() {

	var acc1 Account = &SavingsAccount{accountNo: 12, interestRate: 4, balance: 12000.0}
	var acc2 Account = &CurrentAccount{accountNo: 11, balance: 12000.0}
	acc1.checkBalance()
	acc1.getInterest()
	acc1.depositMoney(1290.0)
	acc1.checkBalance()
	acc2.withdrawMoney(12400)

}

