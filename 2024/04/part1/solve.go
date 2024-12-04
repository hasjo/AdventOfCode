package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func directionalLook(xdir int, ydir int, xstart int, ystart int, inGrid [][]string) int{
    goodCount := 0
    var wantedSlice = []string{"X", "M", "A", "S"}
    for ind, char := range(wantedSlice){
        nextY := ystart + (ydir * ind)
        nextX := xstart + (xdir * ind)
        fmt.Println(nextX, nextY)
        if nextX >= 0 && nextX < len(inGrid[0]) && nextY >= 0 && nextY < len(inGrid) {
            if inGrid[nextY][nextX] == char {
                if ind == len(wantedSlice) - 1{
                    goodCount++
                }
            } else {
                break
            }
        }
    }
    return goodCount
}

func lookAroundYou(xval int, yval int, inGrid [][]string) int {
    foundItems := 0
    directions := [][]int{
        {-1, -1},
        {0, -1},
        {1, -1},
        {-1,0},
        {1, 0},
        {-1, 1},
        {0, 1},
        {1, 1},
    }
    for _, direction := range(directions){
        xdir := direction[0]
        ydir := direction[1]
        foundItems += directionalLook(xdir, ydir, xval, yval, inGrid)
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
            if xitem == "X"{
                newFinds += lookAroundYou(xin, yin, secondDimensionLines)
            }
        }
    }
    fmt.Println(newFinds)
}
