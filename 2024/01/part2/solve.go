package main

import (
	"fmt"
)

func buildIndMap (inSlice []int) map[int]int {
    var returnMap = make(map[int]int)
    for _, index := range(inSlice){
        returnMap[index]++
    }
    return returnMap
}

func main(){
    map2 := buildIndMap(list2)
    similarity := 0

    for _, value := range(list1){
        similarity += (value * map2[value])
    }
    fmt.Println(similarity)
}
