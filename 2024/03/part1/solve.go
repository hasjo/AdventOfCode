package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
        if err != nil {
		log.Fatal("FAILURE TO OPEN INPUT")
	}
        mulMatch := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
        foundSlice := mulMatch.FindAll(data, -1)
        totalMul := 0
        for _, match := range(foundSlice){
            matchStr := string(match)
            workStr := strings.Replace(matchStr, "mul(", "", 1)
            workStr = strings.Replace(workStr, ")", "", 1)
            values := strings.SplitN(workStr, ",", 2)
            value1, err := strconv.Atoi(values[0])
            if err != nil {
                log.Fatal("HEY CANT CONVERT:", values[0])
            }
            value2, err := strconv.Atoi(values[1])
            if err != nil {
                log.Fatal("HEY CANT CONVERT:", values[1])
            }
            newVal := value1 * value2
            totalMul += newVal
        }
        fmt.Println(totalMul)
}
