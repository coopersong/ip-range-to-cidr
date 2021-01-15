package main

import (
    "flag"
    "fmt"
)

import (
    "github.com/coopersong/irtc"
)

var help = flag.Bool("h", false, "help")
var begin = flag.String("begin", "192.168.1.0", "ip range begin")
var end = flag.String("end", "192.168.1.255", "ip range end")

func main() {
    flag.Parse()

    if *help {
        flag.PrintDefaults()
        return
    }

    cidrs, err := irtc.ConvertIpRangeToCidrs(*begin, *end)
    if err != nil {
        fmt.Printf("[%s, %s]: %s", err.Error())
        return
    }

    fmt.Println(cidrs)
}
