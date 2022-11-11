package base122

import "unsafe"

// Zero-copy conversion from byte slice to string
func bytesToString(val []byte) string {
	return *(*string)(unsafe.Pointer(&val))
}

// Zero-copy conversion from string to byte slice
func stringToBytes(val string) []byte {
	strHeader := (*[2]uintptr)(unsafe.Pointer(&val))
	slcHeader := [3]uintptr{strHeader[0], strHeader[1], strHeader[1]}
	return *(*[]byte)(unsafe.Pointer(&slcHeader))
}
