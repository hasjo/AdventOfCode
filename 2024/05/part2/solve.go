package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
    number int
    before []int
    after []int
}

func conditionalAppend(before int, after int, dataMap map[int]Rule){
    modifyBefore := dataMap[before]
    modifyAfter := dataMap[after]
    if modifyBefore.number == 0 {
        modifyBefore.number = before
    }
    if modifyAfter.number == 0 {
        modifyAfter.number = after
    }
    if !slices.Contains(modifyBefore.before, after){
        modifyBefore.before = append(modifyBefore.before, after)
    }
    if !slices.Contains(modifyAfter.after, before){
        modifyAfter.after = append(modifyAfter.after, before)
    }
    dataMap[before] = modifyBefore
    dataMap[after] = modifyAfter
}

func parseRules(inData []string) map[int]Rule {
    var returnMap = make(map[int]Rule)
    for _, line := range(inData) {
        splitSlice := strings.Split(line, "|")
        before := splitSlice[0]
        beforeInt, err := strconv.Atoi(before)
        if err != nil {
            log.Fatal("cant convert to int:", before)
        }
        after := splitSlice[1]
        afterInt, err := strconv.Atoi(after)
        if err != nil {
            log.Fatal("cant convert to int:", before)
        }
        conditionalAppend(beforeInt, afterInt, returnMap)
    }
    return returnMap
}

func validateUpdate(update []int, ruleMap map[int]Rule) int {
    returnVal := 1
    for thisInd, value := range(update){
        for checkInd, checkValue := range(update){
            if checkInd == thisInd {
                continue
            } else if checkInd > thisInd {
                if slices.Contains(ruleMap[value].after, checkValue) == true {
                    return 0
                }
            } else if checkInd < thisInd {
                if slices.Contains(ruleMap[value].before, checkValue) == true {
                    return 0
                }
            }
        }

    }
    return returnVal
}

func intifyStringSlice(inSlice []string) [][]int{
    var outSlice [][]int
    for _, line := range(inSlice){
        splitSlice := strings.Split(line, ",")
        var intSlice []int
        for _, indVal := range(splitSlice){
            thisInt, err := strconv.Atoi(indVal)
            if err != nil {
                log.Fatal("bad news, boss")
            }
            intSlice = append(intSlice, thisInt)
        }
        outSlice = append(outSlice, intSlice)
    }
    return outSlice
}

func getMiddleValue(inSlice []int) int {
    lenNum := len(inSlice)
    halfLen := lenNum/2
    return inSlice[halfLen]
}

func fixUpdateandGetMiddle(inUpdate []int, rulesMap map[int]Rule) int {
    returnval := 0
    return returnval
}

func main(){
    data, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal("oh geez, this goofed")
    }
    var ruleSlice []string
    var updateSlice []string
    stringData := string(data)
    dataLines := strings.Split(stringData, "\n")
    rules := true
    for _, line := range(dataLines){
        if line == ""{
            rules = false
            continue
        }
        if rules == true {
            ruleSlice = append(ruleSlice, line)
        } else {
            updateSlice = append(updateSlice, line)
        }
    }
    updateIntSlice := intifyStringSlice(updateSlice)
    ruleMap := parseRules(ruleSlice)
    var badUpdates [][]int
    for _, update := range(updateIntSlice){
        if validateUpdate(update, ruleMap) == 0 {
            badUpdates = append(badUpdates, update)
        }
    }
    var resultCount int
    for _, update := range(badUpdates){
        resultCount += fixUpdateandGetMiddle(update, ruleMap)
    }
    fmt.Println(resultCount)
}
