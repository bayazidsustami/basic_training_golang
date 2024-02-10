package main

import (
	"fmt"
	"strconv"
)

//input = ["5", "2", "C", "D", "+"]

func main() {
	fmt.Println(callPoints([]string{"5", "2", "C", "D", "+"}))
}

func callPoints(ops []string) int {
	var res = 0
	var newPoints = []int{}
	for i := 0; i < len(ops); i++ {
		intVal, err := strconv.Atoi(ops[i])
		if err != nil {
			if ops[i] == "C" {
				if len(newPoints) > 0 {
					res -= newPoints[len(newPoints)-1]
					newPoints = newPoints[:len(newPoints)-1]
				}
			} else if ops[i] == "D" {
				if len(newPoints) > 0 {
					doubleValue := newPoints[len(newPoints)-1] * 2
					res += doubleValue
					newPoints = append(newPoints, doubleValue)
				}
			} else if ops[i] == "+" {
				if len(newPoints) > 1 {
					sum := newPoints[len(newPoints)-1] + newPoints[len(newPoints)-2]
					res += sum
					newPoints = append(newPoints, sum)
				}
			}
		} else {
			res += intVal
			newPoints = append(newPoints, intVal)
		}
	}
	return res
}

func removeElementByIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}
