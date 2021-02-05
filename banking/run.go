package banking

import (
	"fmt"
	"log"
)

func BankingProject() {
	silvaAccount := NewAccount("silva")
	silvaAccount.Balance()

	silvaAccount.Deposit(10)
	// this would automatically call String()
	fmt.Println(silvaAccount)

	silvaAccount.Withdraw(5)
	fmt.Println(silvaAccount)

	err := silvaAccount.Withdraw(20)
	if err != nil {
		log.Println(err)
	}

}
