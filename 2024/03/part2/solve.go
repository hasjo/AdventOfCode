package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func calcMul(match []byte) int{
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
    return newVal
}

func main() {
	data, err := os.ReadFile("input.txt")
        if err != nil {
		log.Fatal("FAILURE TO OPEN INPUT")
	}
        mulMatch := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
        foundSlice := mulMatch.FindAll(data, -1)
        // indexSlice := mulMatch.FindAllIndex(data,-1)
        doMatch := regexp.MustCompile(`do\(\)`)
        doSlice := doMatch.FindAllIndex(data, -1)
        fmt.Println(doSlice)
        dontMatch := regexp.MustCompile(`don\'t\(\)`)
        dontSlice := dontMatch.FindAllIndex(data, -1)
        fmt.Println(dontSlice)
        totalMul := 0
        for _, match := range(foundSlice){
            newVal := calcMul(match)
            totalMul += newVal
        }
        fmt.Println(totalMul)
}
