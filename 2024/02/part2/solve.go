package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkDistances(initems []int) bool {
	var distanceSlice []int
	for i := 0; i < len(initems) - 1; i++{
		value := initems[i] - initems[i+1]
		distanceSlice = append(distanceSlice, value)
	}
	good := true
	increasing := true
	for ind, value := range(distanceSlice){
		if ind == 0 {
			if value < 0 {
				increasing = false
			} else if value > 0 {
				increasing = true
			} else {
				good = false
				break
			}
		}
		if value == 0 {
			good = false
		}
		if increasing == true {
			if value > 3 || value < 1 {
				good = false
			}
		}
		if increasing == false {
			if value < -3 || value > -1 {
				good = false
			}
		}
	}
	return good
}

func checkLine(initems []int) bool{
	if checkDistances(initems) == true{
		return true
	}
	for i := 0; i<len(initems); i++ {
		var permutationSlice []int
		for ind, item := range(initems){
			if ind != i {
				permutationSlice = append(permutationSlice, item)
			}
		}
		if checkDistances(permutationSlice) == true {
			return true
		}
	}
	return false
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("FAILURE TO OPEN INPUT")
	}
	perLine := strings.Split(string(data), "\n")
	goodCount := 0
	for _, line := range(perLine){
		if line == ""{
			break
		}
		items := strings.Split(line, " ")
		var intItems []int
		for _, item := range(items){
			intItem, err := strconv.Atoi(item)
			if err != nil {
				log.Fatal("CANT CONVERT ITEM,", item)
			}
			intItems = append(intItems, intItem)
		}
		if checkLine(intItems) == true {
			goodCount++
		}
	}
	fmt.Println(goodCount)
}
