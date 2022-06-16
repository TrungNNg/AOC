package main

import (
    "fmt"
    "os"
)
var p = fmt.Println
func main() {
    b, err := os.ReadFile("input.in")
    if err != nil {
        fmt.Println("Error reading input.in")
        os.Exit(1)
    }
    s := string(b)

    res := 0
    for i, r := range s {
        if r == '(' {
            res++
        }
        if r == ')'{
            res--
        }
        if res == -1 {
            fmt.Println(i+1)
            os.Exit(0)
        }
    }
    fmt.Println(res)
}
