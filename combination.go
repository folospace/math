package main

import "fmt"

func main() {
    take := 2
    pool := []int{1,2,3,4,5}
    fmt.Println(fmt.Sprintf("from %v take %d numbers, list all combination:", pool, take))
    for _, v := range ListCombination(pool, take, nil) {
        fmt.Println(v)
    }
}


func ListCombination(pool []int, takeCount int, initialVals []int) [][]int {
    ret := make([][]int, 0)
    if takeCount == 0 {
        ret = append(ret, initialVals)
        return ret
    }
    for k, v := range pool {
        tempVals := make([]int, len(initialVals))
        copy(tempVals, initialVals)
        tempVals = append(tempVals, v)
        tempTry := takeCount - 1
        if tempTry == 0 {
            ret = append(ret, tempVals)
            continue
        }
        if len(pool[k+1:]) < tempTry {
            break
        } else if len(pool[k+1:]) == tempTry {
            tempVals = append(tempVals, pool[k+1:]...)
            ret = append(ret, tempVals)
            break
        } else {
            tempRes := make([]int, len(pool[k+1:]))
            copy(tempRes, pool[k+1:])
            ret = append(ret, ListCombination(tempRes, tempTry, tempVals)...)
        }
    }
    return ret
}