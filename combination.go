package main

import "fmt"

func main() {
    take := 2
    pool := []int{1,2,3,4,5}
    fmt.Println(fmt.Sprintf("from %v take %d numbers, find all combinations:", pool, take))
    for _, v := range FindCombination(pool, take, nil) {
        fmt.Println(v)
    }
}


func FindCombination(pool []int, takeCount int, initialVals []int) [][]int {
    //存储结果集
    ret := make([][]int, 0)

    for k, v := range pool {
        //递归时,继承上层的组合结果
        combination := make([]int, len(initialVals))
        copy(combination, initialVals)
        combination = append(combination, v)

        //取走一个
        leftTryTimes := takeCount - 1

        if leftTryTimes <= 0 {
            //这次组合取的数量满足了, 轮到下一个组合
            ret = append(ret, combination)
            continue
        } else if len(pool[k+1:]) < leftTryTimes {
            //池子剩下的数量不够了,结束循环
            break
        } else if len(pool[k+1:]) == leftTryTimes {
            //池子剩下的数量刚好等于需要的数量, 直接合并数量, 结束循环
            combination = append(combination, pool[k+1:]...)
            ret = append(ret, combination)
            break
        } else {
            //取了一个之后,池子还剩下很多, 递归上述流程
            //将递归结果集存储
            childPool := make([]int, len(pool[k+1:]))
            copy(childPool, pool[k+1:])
            ret = append(ret, FindCombination(childPool, leftTryTimes, combination)...)
        }
    }
    return ret
}