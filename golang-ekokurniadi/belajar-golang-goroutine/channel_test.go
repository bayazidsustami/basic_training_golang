package belajargolanggoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	//channel <- "bay" // mengirim data ke channel channel <- data yang akan dikirim

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "bays"
		fmt.Println("selesai mengirim data")
	}()

	data := <-channel // menerima data dari channel data := < channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeRespose(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "bays"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeRespose(channel)

	data := <-channel // menerima data dari channel data := < channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "bays"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}
