package base122

const (
	shortened = byte(0b00000111) // Used when the last two bytes are used to encode only less than 7 bits
)

var (
	illegalBytes    = []byte{0, 10, 13, 34, 38, 92}
	illegalBytesMap = map[byte]int{
		0:  0, // null
		10: 1, // newline
		13: 2, // carriage return
		34: 3, // double quote
		38: 4, // ampersand
		92: 5, // backslash
	}
)
