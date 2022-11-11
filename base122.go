package base122

// Encode the bytes using base122 scheme
func Encode(data []byte) ([]byte, error) {
	return NewBasicEncoder(data).Encode()
}

// Encode the bytes using base122 scheme
// and transform the result to string
func EncodeToString(data []byte) (string, error) {
	encoded, err := Encode(data)
	if err != nil {
		return "", err
	}
	return bytesToString(encoded), nil
}

// Decode the bytes using base122 scheme
// If the input is not a valid base122 string
// returns ErrInputInvalid error
func Decode(data []byte) ([]byte, error) {
	return NewBasicDecoder(data).Decode()
}

func DecodeFromString(data string) ([]byte, error) {
	return Decode(stringToBytes(data))
}
