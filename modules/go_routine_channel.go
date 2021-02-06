package modules

import (
	"fmt"
	"math/rand"
	"time"
)

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy - ", i)
		time.Sleep(time.Second)
	}
}

/*
isSexy sends str through channel
*/
func isSexy(person string, c chan string) {
	// random processing time (0~9 sec)
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	// send message through channel
	c <- "sexy " + person
}

func ExperienceChannelAndGoRoutine() {
	people := []string{"silva", "nico", "kiwi", "banana", "Gorila", "apple", "pizza"}

	// declare channel
	// - write the data type which channle will contain
	c := make(chan string)
	for _, person := range people {
		// go sexyCount(person)
		go isSexy(person, c)
	}

	// We can make the control-flow be blocked just as much as we need
	// because we can expect the numbers of messages will be sent
	lenPeople := len(people)
	for i := 0; i < lenPeople; i++ {
		fmt.Println("Waiting for message #", i)
		// wait for 1 message from channel (blocking operation)
		fmt.Println("Received: ", <-c)
	}
}
