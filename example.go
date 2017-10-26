package main

import (
        "fmt"
        line1 "github.com/gilwo/formatter/line"
       )



func main() {

	x := line1.LineFormatterField{"A", 20, line1.AlignDirCenter}
	y := line1.LineFormatterField{"B", 20, line1.AlignDirCenter}

	out := line1.LineFormatter("%s <---> %s", x, y)
	fmt.Println(out)

	v := []line1.LineFormatterField{
		{"Z", 5, line1.AlignDirCenter},
		{"X", 5, line1.AlignDirCenter},
	}
	out = line1.LineFormatter("%s <---> %s", v...)
	fmt.Println(out)

	out = line1.LineFormatter2("%s || %s", struct {
		FStr   string
		FLen   int
		FAlign line1.AlignDirection
	}{"A", 20, line1.AlignDirCenter}, struct {
		FStr   string
		FLen   int
		FAlign line1.AlignDirection
	}{"B", 20, line1.AlignDirCenter})
	fmt.Println(out)
}
