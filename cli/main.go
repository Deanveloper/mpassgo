package main

import (
    "github.com/deanveloper/mpassgo"
    "fmt"
)

func main() {
    name := []byte("dean bassett")
    site := []byte("mtu.edu")
    masterPass := []byte("pass")
    pass := mpassgo.GetPassword(name, site, masterPass, 1, mpassgo.Long)

    fmt.Println(string(pass))
}