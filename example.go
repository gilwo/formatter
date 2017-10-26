package main

import (
        "fmt"
        line1 "github.com/gilwo/formatter/line"
       )



func main() {

    f := line1.LineFormatter("<-- %v | %v | %v : %v -->\n",
        []line1.LineFormatterField{
            {20, line1.AlignDirCenter},
            {20, line1.AlignDirCenter},
            {20, line1.AlignDirCenter},
            {20, line1.AlignDirCenter},
        }...)

    fmt.Printf(f("a", "b", "c", "d"))
    fmt.Printf(f("mama", 444, "let", "me"))
    fmt.Printf(f("mama", struct{x int; y string}{545345, "abbb"}, "let", "me"))

}
