package formatter

import (
	"fmt"
)

type AlignDirection int

const (
	AlignDirNone AlignDirection = iota
	AlignDirLeft
	AlignDirCenter
	AlignDirRight
)

func align(val interface{}, fieldLen int, direction AlignDirection) (out string) {
    actVal := fmt.Sprintf("%v", val)
    lenVal := len(actVal)

	spaces := fieldLen - lenVal
    if (spaces < 0 || fieldLen < 0) && direction != AlignDirNone {
		return
	}

	padRight := 0
	padLeft := 0

	switch direction {
	case AlignDirLeft:
		padRight = fieldLen
	case AlignDirCenter:
		padLeft = spaces/2 + lenVal
		padRight = spaces/2 + spaces%2 + padLeft
	case AlignDirRight:
		padLeft = fieldLen
    default:
	}

	padLeftFmt := fmt.Sprintf("%%-%ds", padLeft)
	padRightFmt := fmt.Sprintf("%%%ds", padRight)

	out = fmt.Sprintf(padRightFmt, fmt.Sprintf(padLeftFmt, actVal))

	return
}

func AlignLeft(val string, fieldLen int) string   { return align(val, fieldLen, AlignDirLeft) }
func AlignRight(val string, fieldLen int) string  { return align(val, fieldLen, AlignDirRight) }
func AlignCenter(val string, fieldLen int) string { return align(val, fieldLen, AlignDirCenter) }

type LineFormatterField struct {
	FieldLen   int
	FieldAlign AlignDirection
}

func LineFormatter(format string, fields ...LineFormatterField) func (...interface{}) string {

    return func(values ...interface{}) string {
        if len(values) != len(fields) {
            pmsg := fmt.Sprintf("fields len '%d' (%v) mismatch values len '%d' (%v)",
                len(fields), fields, len(values), values)
            panic(pmsg)
        }
        var vars []interface{}
        for i, e := range fields {
            vars = append(vars, align(values[i], e.FieldLen, e.FieldAlign))
        }
        return fmt.Sprintf(format, vars...)
    }
}

