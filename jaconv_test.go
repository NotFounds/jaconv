package jaconv

import "testing"

func TestKana2roman_blank(t *testing.T) {
	text := ""
	expected := ""
	roman, err := Kana2roman(text)
	if err != nil {
		t.Errorf("failed, error %s", err)
	}
	if roman != expected {
		t.Errorf("failed, got: %v\nexpected:%v", roman, expected)
	}
}
