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
    fmt.Printf(f("mama", 444, "let", "me"))
    fmt.Printf(f("mama", struct{x int; y string}{545345, "ab hfhf hf kjdshkjdh ksjdkjd dkfjh bb"}, "let", "me"))

    fmt.Println()
    f2 := line1.LineFormatter("<-- %s | %s | %s : %s :: '%s' -->\n",
        []line1.LineFormatterField{
            {4, line1.AlignDirCenter},
            {20, line1.AlignDirCenter},
            {20, line1.AlignDirCenter},
            {20, line1.AlignDirNone},
            {0, line1.AlignDirNone},
        }...)

    fmt.Printf(f2("title a", "title b", "title c", "title d", "title e"))
    fmt.Printf(f2("mama", 444, "let", "me", "ending value for first line"))
    fmt.Printf(f2("mama", "let", "me", struct{x int; y string}{545322145, "a very long string to represnet string with length that is not limimted by some size and using align none show everythong ... "}, "bit"))
}
