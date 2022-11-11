package base122

// Encode the bytes using base122 scheme
func Encode(src []byte) ([]byte, error) {
	return NewBasicEncoder(src).Encode()
}

// Encode the bytes using base122 scheme
// and transform the result to string
func EncodeToString(src []byte) (string, error) {
	dst, err := Encode(src)
	if err != nil {
		return "", err
	}
	return string(dst), nil
}
