package ansicolor

const (
	Black   = 0x000000
	Red     = 0xCC0403
	Green   = 0x19CB00
	Yellow  = 0xCECB00
	Blue    = 0x0D73CC
	Magenta = 0xCB1ED1
	Cyan    = 0x0DCDCD
	White   = 0xDDDDDD
)

// RGB is a 24-bit color in 0xRRGGBB form.
type RGB uint32

/*
// Color returns the Color for c. Bits above 24 are ignored.
func (c RGB) Color() Color {
	return Color{uint8(c >> 16), uint8(c >> 8), uint8(c)}
}
*/
