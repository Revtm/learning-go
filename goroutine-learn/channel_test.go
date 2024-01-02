package goroutine_learn

import (
	"fmt"
	"testing"
	"time"
)

func getTestWording(channel chan string) {
	time.Sleep(1 * time.Second)
	channel <- "Test Wording"
}

func writeToChannel(channel chan<- string, str string) {
	time.Sleep(1 * time.Second)
	channel <- str
}

func readFromChannel(channel <-chan string) {
	data := <-channel
	fmt.Println("OUT", data)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	//go func() {
	//	time.Sleep(1 * time.Second)
	//	channel <- "Test!"
	//	fmt.Println("Finished 1")
	//}()

	go getTestWording(channel)

	data := <-channel
	fmt.Println(data)
	fmt.Println("Finished 2")
	time.Sleep(3 * time.Second)

	go writeToChannel(channel, "Test IN")
	go readFromChannel(channel)

	time.Sleep(3 * time.Second)
}
