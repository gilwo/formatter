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

func align(val interface{}, fieldLen int, direction AlignDirection) string {

	actVal := fmt.Sprintf("%v", val)
	lenActVal := len(actVal)
	spaces := fieldLen - lenActVal

	if fieldLen <= 0 {
		return fmt.Sprintf("%s", actVal)
	}

	if spaces < 0 {
		if fieldLen > 3 {
			return fmt.Sprintf("%.*s...", fieldLen-3, actVal)
		} else {
			return fmt.Sprintf("%.*s", fieldLen, actVal)
		}
	}

	padRight := 0
	padLeft := 0

	switch direction {
	case AlignDirLeft:
		padLeft = fieldLen
	case AlignDirCenter:
		padLeft = spaces/2 + lenActVal
		padRight = spaces/2 + spaces%2 + padLeft
	case AlignDirRight:
		padRight = fieldLen
	}

	padLeftFmt := fmt.Sprintf("%%-%ds", padLeft)
	padRightFmt := fmt.Sprintf("%%%ds", padRight)

	return fmt.Sprintf(padRightFmt, fmt.Sprintf(padLeftFmt, actVal))
}

func AlignLeft(val string, fieldLen int) string   { return align(val, fieldLen, AlignDirLeft) }
func AlignRight(val string, fieldLen int) string  { return align(val, fieldLen, AlignDirRight) }
func AlignCenter(val string, fieldLen int) string { return align(val, fieldLen, AlignDirCenter) }

type LineFormatterField struct {
	FieldLen   int
	FieldAlign AlignDirection
}

func LineFormatter(format string, fields ...LineFormatterField) func(...interface{}) string {

	return func(values ...interface{}) string {
		if len(values) != len(fields) {
			pmsg := fmt.Sprintf("fields count '%d' (%v) mismatch values count '%d' (%v)",
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
