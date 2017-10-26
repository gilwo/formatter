package main

import (
        "fmt"
        line1 "github.com/gilwo/formatter/line"
       )



func main() {

    f := line1.LineFormatter("<-- %s | %s | %s : %s -->\n",
        []line1.LineFormatterField{
            {20, line1.AlignDirCenter},
            {20, line1.AlignDirCenter},
            {20, line1.AlignDirCenter},
            {20, line1.AlignDirCenter},
        }...)

    fmt.Printf(f("a", "b", "c", "d"))
    fmt.Printf(f("mama", "mia", "let", "me"))

}
