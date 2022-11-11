package base122

import "testing"

func TestDecode(t *testing.T) {
	if decoded, err := NewBasicDecoderFromString(":.7!He1[lF+H").Decode(); err != nil || bytesToString(decoded) != "test decoder" {
		t.Log("decode failed 1")
		t.Fail()
	}
	if decoded, err := NewBasicDecoderFromString(";.'Ile9$3dyLWd@v2׏YJr<Hf+Ir ;.'Ile9$3dyLWd@v2׏YJr<Hf+Ir ;.'Ile9$c=קWCPB!@").Decode(); err != nil || bytesToString(decoded) != "very very very very very very very very very very very very very very very very long text!!!" {
		t.Log("decode failed 2")
		t.Fail()
	}
}
