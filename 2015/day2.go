package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "sort"
)

var p = fmt.Println
var pf = fmt.Printf

func main() {
    f, err := os.Open("input.in")
    if err != nil {
        fmt.Println("Error open input.in")
        os.Exit(1)
    }
    scanner := bufio.NewScanner(f)
    var l,w,h,res int
    for scanner.Scan() {
        list := strings.Split(scanner.Text(), "x")
        l,_ = strconv.Atoi(list[0])
        w,_ = strconv.Atoi(list[1])
        h,_ = strconv.Atoi(list[2])
        temp := []int{l,w,h}
        sort.Ints(temp)
        peri := 2*(temp[0] + temp[1])
        vol := l*w*h
        res += peri + vol
    }
    p(res)
}
