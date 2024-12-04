package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
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

func doOrDont(locVal int, instMap map[int]string, sortedKeys []int) bool {
    lastVal := 0
    for _, key := range(sortedKeys){
        if key > locVal && lastVal != 0{
            if instMap[lastVal] == "do" {
                return true
            } else {
                return false
            }
        } else if key < locVal {
            lastVal = key
        } else if key > locVal && lastVal == 0 {
            return true
        }
    }
    return false
}

func sortMapKeys(inMap map[int]string) []int{
    var returnSlice []int
    for key := range(inMap) {
        returnSlice = append(returnSlice, key)
    }
    slices.Sort(returnSlice)
    return returnSlice
}

func processMuls(mulSlice [][]int, dataSlice [][]byte, instMap map[int]string) int {
    totalValue := 0
    sortedMapKeys := sortMapKeys(instMap)
    for key, val := range(instMap) {
        fmt.Println(key, " - ", val)
    }
    for index, value := range(mulSlice) {
        fmt.Println(value[0])
        if doOrDont(value[0], instMap, sortedMapKeys) == true {
            totalValue += calcMul(dataSlice[index])
        }
    }
    return totalValue
}

func main() {
	data, err := os.ReadFile("input.txt")
        if err != nil {
		log.Fatal("FAILURE TO OPEN INPUT")
	}
        mulMatch := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
        foundSlice := mulMatch.FindAll(data, -1)
        indexSlice := mulMatch.FindAllIndex(data,-1)
        var instructionMap = make(map[int]string)
        doMatch := regexp.MustCompile(`do\(\)`)
        doSlice := doMatch.FindAllIndex(data, -1)
        dontMatch := regexp.MustCompile(`don\'t\(\)`)
        dontSlice := dontMatch.FindAllIndex(data, -1)
        for _, item := range(doSlice) {
            instructionMap[item[0]] = "do"
        }
        for _, item := range(dontSlice) {
            instructionMap[item[0]] = "dont"
        }
        totalValue := processMuls(indexSlice, foundSlice, instructionMap)
        fmt.Println(totalValue)
}
