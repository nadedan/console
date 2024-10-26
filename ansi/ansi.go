package ansi

import "strconv"

type direction string

const (
	esc = "\x1b"
	osb = "["
	csi = esc + osb

	Up    direction = "A"
	Down  direction = "B"
	Right direction = "C"
	Left  direction = "D"

	position = "H"
	clear    = "J"
)

func Cursor(d direction, n uint8) string {
	if n == 0 {
		return ""
	}
	return csi + strconv.Itoa(int(n)) + string(d)
}

func Position(row uint8, col uint8) string {
	return csi + strconv.Itoa(int(row)) + ";" + strconv.Itoa(int(col)) + position
}

func ClearDown() string {
	return csi + "0" + clear
}

func ClearUp() string {
	return csi + "1" + clear
}

func Clear() string {
	return csi + "2" + clear
}
