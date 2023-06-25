package main

import (
    "fmt"
    "sort"
    "time"
)

type Bucket struct {
    Cap  int
    Len  int
    Want int
}

func main() {
    a := Bucket{Cap: 12, Len: 12, Want: 6}
    b := Bucket{Cap: 9, Len: 0, Want: 6}
    c := Bucket{Cap: 5, Len: 0, Want: 0}

    var success = make(chan Bucket, 100)

    go SwapBucket(a, b, c, cap(success), success, []Bucket{a, b, c})

    i := 0
LOOP:
    for {
        select {
        case v := <-success:
            if v.Cap == 0 {
                break LOOP
            }
            fmt.Print(v.Len, " ")
            i++
            if i%3 == 0 {
                fmt.Println()
            }
        }
    }

}

func SwapBucket(a, b, c Bucket, depth int, success chan Bucket, caches []Bucket) {
    if depth <= 0 || len(success) > 0 {
        return
    }
    time.Sleep(time.Millisecond * 100)

    delta := b.Cap - b.Len
    if a.Len < delta {
        delta = a.Len
    }

    if delta > 0 {
        a.Len -= delta
        b.Len += delta
        temp := []Bucket{a, b, c}
        sort.SliceStable(temp, func(i, j int) bool {
            return temp[i].Cap > temp[j].Cap
        })
        caches = append(caches, temp...)
        if a.Len == a.Want && b.Len == b.Want && c.Len == c.Want && len(success) == 0 {
            for k := range caches {
                success <- caches[k]
            }
            success <- Bucket{}
            return
        }
    }

    caches = append([]Bucket{}, caches...)

    depth -= 1
    go SwapBucket(a, c, b, depth, success, caches)
    go SwapBucket(b, c, a, depth, success, caches)
    go SwapBucket(c, a, b, depth, success, caches)
    go SwapBucket(c, b, a, depth, success, caches)
}
