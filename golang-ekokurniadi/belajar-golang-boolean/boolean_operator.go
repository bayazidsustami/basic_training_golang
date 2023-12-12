package main

import "fmt"

func main() {
	var finalScore = 90
	var attendance = 80

	var isPassedFinalScore = finalScore > 80
	var isPassedAttendace = attendance > 80

	var isPassed = isPassedFinalScore && isPassedAttendace

	fmt.Println(isPassed)
}
