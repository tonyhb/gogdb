package gui

import (
	"mime"
	"strings"
)

func getMime(filename string) string {
	period := strings.LastIndex(filename, ".")
	if period < 0 {
		return ""
	}
	return mime.TypeByExtension(filename[period:])
}
