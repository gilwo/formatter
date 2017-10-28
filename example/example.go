package main

import (
        "fmt"
        "github.com/gilwo/formatter"
       )



func main() {

    f := formatter.LineFormatter("<-- %s | %s | %s : %s -->\n",
        []formatter.LineFormatterField{
            {20, formatter.AlignDirCenter},
            {20, formatter.AlignDirCenter},
            {20, formatter.AlignDirCenter},
            {20, formatter.AlignDirCenter},
        }...)

    fmt.Printf(f("a", "b", "c", "d"))
    fmt.Printf(f("mama", 444, "let", "me"))
    fmt.Printf(f("mama", struct{x int; y string}{545345, "ab hfhf hf kjdshkjdh ksjdkjd dkfjh bb"}, "let", "me"))

    fmt.Println()
    f2 := formatter.LineFormatter("<-- %s | %s | %s : %s :: '%s' -->\n",
        []formatter.LineFormatterField{
            {4, formatter.AlignDirCenter},
            {20, formatter.AlignDirCenter},
            {20, formatter.AlignDirCenter},
            {20, formatter.AlignDirNone},
            {0, formatter.AlignDirNone},
        }...)

    fmt.Printf(f2("title a", "title b", "title c", "title d", "title e"))
    fmt.Printf(f2("mama", 444, "let", "me", "ending value for first line"))
    fmt.Printf(f2("mama", "let", "me", struct{x int; y string}{545322145, "a very long string to represnet string with length that is not limimted by some size and using align none show everythong ... "}, "bit"))


    fields := []formatter.LineFormatterField{
                {15, formatter.AlignDirCenter},
                {7,  formatter.AlignDirLeft},
                {10, formatter.AlignDirRight},
                {9,  formatter.AlignDirNone},   // <- no alignment, just limit the field length
                {0,  formatter.AlignDirCenter}, // <- field of len 0 override any alignment
            }
    lineFunc := formatter.LineFormatter("< %s | %s | %s | %s | %s -->\n", fields...)

    fmt.Printf(lineFunc("A", "B", "C", "D", "D" ))
    fmt.Printf(lineFunc("Coloumn A", "Column B", "Column C", "Column D", "Column E" ))
    fmt.Printf(lineFunc("mama", 444, "let", "me and you go to where we want", 99191))
    fmt.Printf(lineFunc("mama", struct{x string; y int}{"why so serious", 777}, "let", "me", 111))
    fmt.Printf(lineFunc("mama", 1234567890, "let", "fantastic mr. fox", map[string]struct{x string; y int}{"key": {"val", 5}}))
}
