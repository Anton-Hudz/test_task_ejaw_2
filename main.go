package main

import (
	"bytes"
	"strings"
)

func concat(str []string) string {
	result := ""
	for _, v := range str {
		result += v
	}
	return result
}

func concatStringBuilder(str []string) string {
	var builder strings.Builder
	for _, v := range str {
		builder.WriteString(v)
	}
	return builder.String()
}

func concatWithJoin(str []string) string {
	return strings.Join(str, "")
}

func concatWithBuffer(str []string) string {
	var buffer bytes.Buffer
	for _, v := range str {
		buffer.WriteString(v)
	}
	return buffer.String()
}
