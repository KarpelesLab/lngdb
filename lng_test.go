package lngdb

import "testing"

func TestLng(t *testing.T) {
	// some tests
	l := Languages["en-US"]

	if l.LocalName != "English (US)" {
		t.Errorf("invalid name for en-US: %s", l.LocalName)
	}

	l = Languages["ru-RU"]

	if l.Locale.Day[1] != "Понедельник" {
		t.Errorf("invalid value for monday in Russia: %s (should be Понедельник)", l.Locale.Day[1])
	}
}
