package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Bank transaction between sender and receiver
	senderBalance := 10000
	chn := make(chan int)
	receiverBalance := 500
	go moneySender(&senderBalance, chn)
	go moneyReceiver(&receiverBalance, chn)
	fmt.Scanln()
}


func moneySender(balance *int, sndMoney chan<- int) {

	randomNum := rand.Intn(1000) + 400
	for *balance >= randomNum {
		*balance -= randomNum
		fmt.Println("Send money to receiver:", randomNum)
		fmt.Println("Balance of sender:", *balance)
		sndMoney <- randomNum
		randomNum = rand.Intn(1000) + 400
		time.Sleep(500)
	}

	if *balance < randomNum {
		close(sndMoney)
		fmt.Println("Transaction aborted...")
	}
}

func moneyReceiver(balance *int, rcvMoney <-chan int) {
	for {
		money, opened := <-rcvMoney
		if !opened {
			break
		}
		*balance += money
		fmt.Println("Get sum from receiver in size:", money)
		fmt.Println("Balance of receiver:", *balance)
		time.Sleep(500)
	}

}

