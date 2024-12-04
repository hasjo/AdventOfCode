package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func goodPair(inSlice [][]string, checkSlice []string) bool{
    for _, option := range(inSlice){
        good := true
        for ind, item := range(checkSlice){
            if option[ind] != item {
                good = false
            }
        }
        if good == true{
            return good
        }
    }
    return false
}

func lookAroundYou(xval int, yval int, inGrid [][]string) int {
    foundItems := 0
    directions := [][]int{
        {-1, -1},
        {1, -1},
        {-1, 1},
        {1, 1},
    }
    charCounts := make(map[string]int)
    var foundSlice []string
    for _, direction := range(directions){
        xdir := direction[0]
        ydir := direction[1]
        findX := xval + xdir
        findY := yval + ydir
        foundChar := inGrid[findY][findX]
        foundSlice = append(foundSlice, foundChar)
        charCounts[foundChar]++
    }
    if charCounts["M"] == 2 && charCounts["S"] == 2{
        var options = [][]string{{"M","S"}, {"S","M"}}
        matchSlice1 := []string{foundSlice[0],foundSlice[3]}
        matchSlice2 := []string{foundSlice[1],foundSlice[2]}
        match1 := goodPair(options, matchSlice1)
        match2 := goodPair(options,matchSlice2)
        if match1 == true && match2 == true {
            foundItems++
        }
        fmt.Println(charCounts)
    }
    
    return foundItems
}

func main(){
    data, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal("BAD NEWS BUDDY")
    }
    stringData := string(data)
    dataLines := strings.Split(stringData, "\n")
    var secondDimensionLines [][]string
    for _, line := range(dataLines){
        sliceLine := strings.Split(line, "")
        if len(sliceLine) > 0 {
            secondDimensionLines = append(secondDimensionLines,sliceLine)
        }
    }
    var newFinds int
    for yin, yline := range(secondDimensionLines){
        for xin, xitem := range(yline){
            if xitem == "A" && xin > 0 && yin > 0 && xin < len(yline) - 1 && yin < len(secondDimensionLines) - 1{
                newFinds += lookAroundYou(xin, yin, secondDimensionLines)
            }
        }
    }
    fmt.Println(newFinds)
}
