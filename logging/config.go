package logging

import "os"

type Config struct {
	File     *os.File
	Level    Level
	Colorize bool
}
