package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("FAILURE TO OPEN INPUT")
	}
	perLine := strings.Split(string(data), "\n")
	goodCount := 0
	for _, line := range(perLine){
		increasing := false
		good := true
		items := strings.Split(line, " ")
		for i := 0; i < len(items) - 1; i++ {
			item1, err := strconv.Atoi(items[i])
			if err != nil{
				log.Fatal("FAILURE TO CONVERT ITEM1")
			}
			item2, err := strconv.Atoi(items[i+1])
			if err != nil{
				log.Fatal("FAILURE TO CONVERT ITEM1")
			}
			value := item1 - item2
			if i == 0 {
				if value > 0 {
					increasing = true
				}
			}
			if math.Abs(float64(value)) == 0 {
				good = false
				break
			}
			if increasing == true {
				if value < 1 || value > 3 {
					good = false
				}
			}
			if increasing == false {
				if value > -1 || value < -3 {
					good = false
				}
			}
		}
		if good == true && line != ""{
			fmt.Println(line)
			goodCount++
		}
	}
	fmt.Println(goodCount)
}
