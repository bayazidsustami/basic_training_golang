package main

import (
	"fmt"
	"time"
)

func main() {
	var time1 = time.Now()
	fmt.Printf("time %v\n", time1)

	var time2 = time.Date(2021, 11, 23, 11, 0, 0, 0, time.UTC)
	fmt.Printf("time %v\n", time2)

	fmt.Println("Year :", time1.Year(), "month:", time1.Month(), "year day:", time.Now().YearDay())

	parsingStringToTime()
	predefinedLayout()
	timeToString()
	errorHandlingTimeParse()
	timeSleep()
}

func parsingStringToTime() {
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"

	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())
}

func predefinedLayout() {
	var date, _ = time.Parse(time.RFC822, "02 Sep 20 08:00 WIB")
	fmt.Println("time : \t", date.String())
}

func timeToString() {
	var date, _ = time.Parse(time.RFC822, "02 Sep 20 08:00 WIB")
	var dates1 = date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println("dates1", dates1)

	var dates2 = date.Format(time.RFC3339)
	fmt.Println("dates2", dates2)
}

func errorHandlingTimeParse() {
	var date, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")
	if err != nil {
		fmt.Println("Error parse", err.Error())
		return
	}
	fmt.Println(date)
}

//blocking all program operation until sleep finish
func timeSleep() {
	fmt.Println("start")
	time.Sleep(time.Second * 4)
	fmt.Println("after 4 second")

	//scheduler()
	timeNewTimer()
	timerAfterFunc()
	schedulerWithTicker()
}

//for schedule print hello every one second
/*func scheduler() {
	for true {
		fmt.Println("Hello !!")
		time.Sleep(1 * time.Second)
	}
}*/

func timeNewTimer() {
	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("Start")
	<-timer.C
	fmt.Println("Finish")
}

func timerAfterFunc() {
	var ch = make(chan bool)
	time.AfterFunc(4*time.Second, func() {
		fmt.Println("Expired")
		ch <- true
	})

	fmt.Println("Start timer after func")
	<-ch
	fmt.Println("finish after func")

	<-time.After(4 * time.Second)
	fmt.Println("expired")
}

func schedulerWithTicker() {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("Hello !!", t)
		}
	}

}
