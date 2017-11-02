# formatter

formatter of text when printing consecutive data which need to be at the same format

usage example
```go
import "github.com/gilwo/formatter"
```

need to set fields
```go
    fields := []formatter.LineFormatterField{
                {15, formatter.AlignDirCenter},
                {7,  formatter.AlignDirLeft},
                {10, formatter.AlignDirRight},
                {9,  formatter.AlignDirNone},   // <- no alignment, just limit the field length
                {0,  formatter.AlignDirCenter}, // <- field of len 0 override any alignment
            }
```

define format string
```go
    formatString := "< %s | %s | %s | %s | %s -->\n"
```

now get the closure function 
```go
    lineFunc := formatter.LineFormatter(formatString, fields...)
```

now just print the values
```go
    fmt.Printf(lineFunc("A", "B", "C", "D", "E" ))
    fmt.Printf(lineFunc("Coloumn A", "Column B", "Column C", "Column D", "Column E" ))
    fmt.Printf(lineFunc("mama", 444, "let", "me and you go to where we want", 99191))
    fmt.Printf(lineFunc("mama", struct{x string; y int}{"why so serious", 777}, "let", "me", 111))
    fmt.Printf(lineFunc("mama", 1234567890, "let", "fantastic mr. fox", map[string]struct{x string; y int}{"key": {"val", 5}}))
```

the outout will look like this
```bash
<        A        | B       |          C | D | E -->
<    Coloumn A    | Colu... |   Column C | Column D | Column E -->
<       mama      | 444     |        let | me and... | 99191 -->
<       mama      | {why... |        let | me | 111 -->
<       mama      | 1234... |        let | fantas... | map[key:{val 5}] -->
```

##### * Note about formatter panic: 
a panic will be called in the case when the variable `fields` variadic count passed to `LineFormatter` method is not the same count as the `values` variadic variable passed to the closure function created by calling `LineFormatter`

