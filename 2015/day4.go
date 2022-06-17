package main

import (
    "fmt"
    "crypto/md5"
    "strconv"
)

var p = fmt.Println
var pf = fmt.Printf

func main() {
    input := "bgvyzdsv"
    i := 1
    for {
        hash := fmt.Sprintf("%x", md5.Sum( []byte(input + strconv.Itoa(i)) ))
        if hash[:6] == "000000" {
            p(hash)
            p(input,i)
            break
        }
        i++
    }
}
