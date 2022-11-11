package base122

type BasicDecoder struct {
	data      []byte
	curByte   byte
	curOffset byte
}

func NewBasicDecoder(data []byte) *BasicDecoder {
	return &BasicDecoder{data, 0, 0}
}

func NewBasicDecoderFromString(data string) *BasicDecoder {
	return &BasicDecoder{stringToBytes(data), 0, 0}
}

func (dec *BasicDecoder) Decode() ([]byte, error) {
	var resultBuffer []byte
	// loop the data as a UTF-8 string
	strData := bytesToString(dec.data)
	for _, nextRune := range strData {
		// UTF-8 codePoint should > 127
		if nextRune > 127 {
			// double bytes case
			// 0b110xxx1y 0b10yyyyyy
			// extract xxx from the first byte
			illegalIndex := byte(nextRune>>10) & 0b00000111
			if illegalIndex != shortened {
				dec.putNext7Bits(&resultBuffer, illegalBytes[illegalIndex])
			}
			dec.putNext7Bits(&resultBuffer, byte(nextRune)&127)
		} else {
			// single byte case, put into the buffer directly
			dec.putNext7Bits(&resultBuffer, byte(nextRune))
		}
	}
	return resultBuffer, nil
}

func (dec *BasicDecoder) putNext7Bits(resultBuffer *[]byte, nextByte byte) {
	nextByte <<= 1
	dec.curByte |= (nextByte >> dec.curOffset)
	dec.curOffset += 7
	if dec.curOffset >= 8 {
		// we have a full 8-bits byte now, put it into the buffer
		*resultBuffer = append(*resultBuffer, dec.curByte)
		// reduce the offset (same as taking the modular of 8)
		dec.curOffset -= 8
		// put the remaining bits into curByte
		dec.curByte = nextByte << (7 - dec.curOffset)
	}
}
