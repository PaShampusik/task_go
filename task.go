package main

import (
    "fmt"
    "sort"
)

func main() {
    var n int64
    fmt.Scanln(&n)

    a := make([][]int64, n)
    for i := int64(0); i < n; i++ {
        a[i] = make([]int64, n)
        for j := int64(0); j < n; j++ {
            fmt.Scan(&a[i][j])
        }
    }

    cntStr := make([]int64, n)
    cntStlb := make([]int64, n)

    for i := int64(0); i < n; i++ {
        for j := int64(0); j < n; j++ {
            cntStr[i] += a[i][j]
        }
    }

    for j := int64(0); j < n; j++ {
        for i := int64(0); i < n; i++ {
            cntStlb[j] += a[i][j]
        }
    }

    sort.Slice(cntStr, func(i, j int) bool {
        return cntStr[i] < cntStr[j]
    })

    sort.Slice(cntStlb, func(i, j int) bool {
        return cntStlb[i] < cntStlb[j]
    })

    if slicesEqual(cntStr, cntStlb) {
        fmt.Println("yes")
    } else {
        fmt.Println("no")
    }
}

func slicesEqual(a, b []int64) bool {
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}