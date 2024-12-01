package main

import (
	"fmt"
	"math"
	"sort"
)

func main(){
    sort.Ints(list1)
    sort.Ints(list2)
    totaldistance := 0
    for ind, value := range(list1){
        distance := value - list2[ind]
        floatDistance := float64(distance)
        absDistance := int(math.Abs(floatDistance))
        totaldistance += absDistance
    }
    fmt.Println(totaldistance)
}
