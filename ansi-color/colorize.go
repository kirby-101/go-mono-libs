// Package ansicolor colors strings for terminals using ANSI truecolor (24-bit).
package ansicolor

import "fmt"

const reset = "\x1b[0m"

func ColorizeString(color RGB, s string) string {
	if s == "" {
		return s
	}

	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s%s", uint8(color>>16), uint8(color>>8), uint8(color), s, reset)
}

/*
 * ColorizeColorString(string) string
 */

func ColorizeBlackString(s string) string {
	return ColorizeString(Black, s)
}

func ColorizeRedString(s string) string {
	return ColorizeString(Red, s)
}

func ColorizeGreenString(s string) string {
	return ColorizeString(Green, s)
}

func ColorizeYellowString(s string) string {
	return ColorizeString(Yellow, s)
}

func ColorizeBlueString(s string) string {
	return ColorizeString(Blue, s)
}

func ColorizeMagentaString(s string) string {
	return ColorizeString(Magenta, s)
}

func ColorizeCyanString(s string) string {
	return ColorizeString(Cyan, s)
}

func ColorizeWhiteString(s string) string {
	return ColorizeString(White, s)
}
