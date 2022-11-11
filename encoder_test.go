package base122

import (
	"bytes"
	"testing"
)

func TestEncoderNext7Bits(t *testing.T) {
	encoder := NewBasicEncoder([]byte{0b11111111, 0b00001111})
	if r, err := encoder.next7Bits(); err != nil || r != 0b01111111 {
		t.Log("next 7 bits failed 1")
		t.FailNow()
	}
	if r, err := encoder.next7Bits(); err != nil || r != 0b01000011 {
		t.Log("next 7 bits failed 2")
		t.FailNow()
	}
	if r, err := encoder.next7Bits(); err != nil || r != 0b01100000 {
		t.Log("next 7 bits failed 3")
		t.FailNow()
	}
	if _, err := encoder.next7Bits(); err != ErrEndOfBytes {
		t.Log("next 7 bits failed 4")
		t.FailNow()
	}
}

func TestEncode(t *testing.T) {
	if encoded, err := Encode([]byte("hello world")); err != nil || !bytes.Equal(encoded, []byte{52, 25, 45, 70,
		99, 60, 64, 119, 55, 215, 141, 70, 32}) {
		t.Log("encode failed 1")
		t.Fail()
	}
	if encoded, err := Encode([]byte("")); err != nil || !bytes.Equal(encoded, []byte{}) {
		t.Log("encode failed 2")
		t.Fail()
	}
	if encoded, err := Encode([]byte("very very very very very very very very very very very " +
		"very very very very very long text!!!")); err != nil || !bytes.Equal(encoded, []byte{59,
		25, 46, 39, 73, 1, 108, 101, 57, 30, 36, 7, 51, 21, 100, 121, 16, 29, 76, 87, 19, 100, 64,
		118, 50, 215, 143, 18, 3, 89, 74, 114, 60, 72, 14, 102, 43, 73, 114, 32, 59, 25, 46, 39, 73,
		1, 108, 101, 57, 30, 36, 7, 51, 21, 100, 121, 16, 29, 76, 87, 19, 100, 64, 118, 50, 215, 143,
		18, 3, 89, 74, 114, 60, 72, 14, 102, 43, 73, 114, 32, 59, 25, 46, 39, 73, 1, 108, 101, 57, 30,
		36, 6, 99, 61, 215, 167, 16, 29, 12, 87, 67, 80, 66, 33, 16, 64}) {
		t.Log("encode failed 3")
		t.Fail()
	}
}
