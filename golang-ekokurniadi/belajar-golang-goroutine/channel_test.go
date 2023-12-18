package belajargolanggoroutine

import (
	"fmt"
	"strconv"
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

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "bay"
		channel <- "bayazid"
		channel <- "yazid"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima", data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeRespose(channel1)
	go GiveMeRespose(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
