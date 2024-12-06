package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	x int
	y int
}

func rotateDir(in Coord) Coord {
	outx := in.x
	outy := in.y

	outx++
	outy++

	if outx == 3 {
		outx = -1
	}
	if outy == 3 {
		outy = -1
	}

	var outCoord Coord
	outCoord.x = outx
	outCoord.y = outy
	return outCoord
}

func translateVec(in Coord) Coord {
	checkX := in.x
	checkY := in.y
	if checkX == 2 {
		checkX = 0
	}
	if checkY == 2 {
		checkY = 0
	}
	return Coord{x:checkX, y:checkY}
}

func findGuard(theGrid [][]string) Coord {
	var foundCoord Coord
	for ycoord, yval := range(theGrid){
		for xcoord, xval := range(yval){
			if xval == "^"{
				foundCoord.x = xcoord
				foundCoord.y = ycoord
			}
		}
	}
	return foundCoord
}

func checkCoordSlice(x int, y int, coordSlice []Coord) bool{
	for _, coord := range(coordSlice){
		if coord.x == x && coord.y == y {
			return true
		}
	}
	return false
}

func calculateGuardPath(guard Coord, vec Coord, theGrid [][]string) int {
	returnCount := 0
	holdVec := vec
	dirVec := translateVec(vec)
	nextX := guard.x
	nextY := guard.y
	var spotSlice []Coord
	for nextX > -1 && nextX < len(theGrid[0]) && nextY > -1 && nextY < len(theGrid) {
		fmt.Printf("%d - %d\n", nextX, nextY)
		if theGrid[nextY][nextX] == "#" {
			holdVec = rotateDir(holdVec)
			dirVec = translateVec(holdVec)
			nextX = guard.x + dirVec.x
			nextY = guard.y + dirVec.y
		} else {
			guard.x = nextX
			guard.y = nextY
			nextX += dirVec.x
			nextY += dirVec.y
		}
		if checkCoordSlice(guard.x, guard.y, spotSlice) == false {
			spotSlice = append(spotSlice, Coord{x:guard.x, y:guard.y})
			returnCount++
		}
	}
	return returnCount
}

func main(){
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("aww man")
	}
	dataString := string(data)
	dataLines := strings.Split(dataString, "\n")
	var theGrid [][]string
	for _, line := range(dataLines){
		splitLine := strings.Split(line, "")
		theGrid = append(theGrid, splitLine)
	}
	theGuard := findGuard(theGrid)
	var guardDir = Coord{x: 0, y:-1}
	result := calculateGuardPath(theGuard, guardDir, theGrid)
	fmt.Println(result)
}
