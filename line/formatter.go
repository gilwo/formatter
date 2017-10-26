package formatter

import (
	"fmt"
)

type AlignDirection int

const (
	AlignDirLeft AlignDirection = iota
	AlignDirCenter
	AlignDirRight
)

func align(val string, fieldLen int, direction AlignDirection) (out string) {
	spaces := fieldLen - len(val)
	if spaces < 0 || fieldLen < 0 {
		return
	}

	padRight := 0
	padLeft := 0

	switch direction {
	case AlignDirLeft:
		padRight = fieldLen
	case AlignDirCenter:
		padLeft = spaces/2 + len(val)
		padRight = spaces/2 + spaces%2 + padLeft
	case AlignDirRight:
		padLeft = fieldLen
	}

	padLeftFmt := fmt.Sprintf("%%-%ds", padLeft)
	padRightFmt := fmt.Sprintf("%%%ds", padRight)

	out = fmt.Sprintf(padRightFmt, fmt.Sprintf(padLeftFmt, val))

	return
}

func AlignLeft(val string, fieldLen int) string   { return align(val, fieldLen, AlignDirLeft) }
func AlignRight(val string, fieldLen int) string  { return align(val, fieldLen, AlignDirRight) }
func AlignCenter(val string, fieldLen int) string { return align(val, fieldLen, AlignDirCenter) }

type LineFormatterField struct {
	FieldStr   string
	FieldLen   int
	FieldAlign AlignDirection
}

func LineFormatter(format string, fields ...LineFormatterField) (out string) {
	//func LineFormatter(format string, fields ...struct{fieldStr string; fieldLen int; align AlignDirection}) (out string) {

	//var vars []string
	var vars []interface{}
	for _, e := range fields {
		//vars = append(vars, align(e.fieldStr, e.fieldLen, e.align))
		vars = append(vars, align(e.FieldStr, e.FieldLen, e.FieldAlign))
	}

	//test := []interface{}(vars)
	//return fmt.Sprintf(format, test...)
	return fmt.Sprintf(format, vars...)
}

func LineFormatter2(format string, fields ...struct {
	FStr   string
	FLen   int
	FAlign AlignDirection
}) (out string) {
	var vars []interface{}
	for _, e := range fields {
		vars = append(vars, align(e.FStr, e.FLen, e.FAlign))
	}
	return fmt.Sprintf(format, vars...)
}

