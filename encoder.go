package base122

import (
	"errors"
)

const (
	shortened = 0b00000111 // Used when the last two bytes are used to encode only less than 7 bits
)

var (
	errEndOfBytes = errors.New("end of bytes")
	illegalBytes  = map[byte]int{
		0:  0, // null
		10: 1, // newline
		13: 2, // carriage return
		34: 3, // double quote
		38: 4, // ampersand
		92: 5, // backslash
	}
)

type BasicEncoder struct {
	data    []byte
	curByte int
	curBit  int
}

func NewBasicEncoder(data []byte) *BasicEncoder {
	return &BasicEncoder{data, 0, 0}
}

func (enc *BasicEncoder) Encode() ([]byte, error) {
	// set the cap of the result byte slice to the length of source data bytes times 2
	result := make([]byte, 0, 2*len(enc.data))
	for {
		// get the next 7-bits byte
		firstSevenBitsByte, err := enc.next7Bits()
		if err != nil {
			// end of the data, break the loop
			if err == errEndOfBytes {
				break
			}
			// encouter other error, return the error
			return nil, err
		}
		// check if the byte hits the illegal bytes,
		// and get the index of first byte (which is the illegal byte)
		// inside the illegal byte list
		// e.g new line \r (10) --> 1
		var illegalByteIndex int
		if index, ok := illegalBytes[firstSevenBitsByte]; !ok {
			// if no hit, just save the single 7-bits byte into the result
			result = append(result, firstSevenBitsByte)
			continue
		} else {
			illegalByteIndex = index
		}

		firstEncodedByte := byte(0b11000010)
		secondEncodedByte := byte(0b10000000)

		// if the byte hits the illegal bytes, need to encode it
		// into 2-bytes format together with the next 7-bits byte
		hasNextByte := true
		secondSevenBitsByte, err := enc.next7Bits()
		if err != nil {
			if err == errEndOfBytes {
				hasNextByte = false
			} else {
				// encouter other error, return the error
				return nil, err
			}
		}

		if hasNextByte {
			firstEncodedByte |= (0b00000111 & byte(illegalByteIndex)) << 2
		} else {
			firstEncodedByte |= shortened << 2
			secondSevenBitsByte = firstSevenBitsByte // Encode the first 7-bits byte into the last byte in the shortened case
		}

		// put the first bit into the first byte to encode
		firstEncodedByte |= (0b01000000 & secondSevenBitsByte) >> 6
		// put the rest 6 bits into the second byte to encode
		secondEncodedByte |= (0b00111111 & secondSevenBitsByte)

		result = append(result, firstEncodedByte)
		result = append(result, secondEncodedByte)
	}
	return result, nil
}

func (enc *BasicEncoder) next7Bits() (byte, error) {
	if enc.curByte >= len(enc.data) {
		return 0, errEndOfBytes
	}
	firstByte := enc.data[enc.curByte]
	// extract the hightest (7 - enc.curBit) bits of the first byte, and shift 1 bit to the right
	// e.g. if enc.curBit == 6, 0b11111111 ---> 0b01100000
	firstEncodedByte := ((0b11111110 >> enc.curBit) & firstByte) << enc.curBit >> 1
	// check if we need to encode the second byte
	enc.curBit += 7
	// no need to encode the second byte, return the first byte
	if enc.curBit < 8 {
		return firstEncodedByte, nil
	}
	// encode the next byte
	enc.curBit -= 8
	enc.curByte++
	// already go to the end, no next byte
	// just return the first encoded byte
	if enc.curByte >= len(enc.data) {
		return firstEncodedByte, nil
	}
	secondByte := enc.data[enc.curByte]
	// extract the highest enc.curBit bits of the second byte, and align to the right
	// e.g. if enc.curBit == 2, bitsToMove == 6, 0b10101010 ---> 0b00101010
	bitsToMove := 8 - enc.curBit
	secondEncodedByte := (0b11111111 >> bitsToMove << bitsToMove) & secondByte >> bitsToMove
	return firstEncodedByte | secondEncodedByte, nil
}
