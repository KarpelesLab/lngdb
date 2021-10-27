package lngdb

import "golang.org/x/text/language"

var German = &Lng{
	Tag:       language.German,
	LocalName: "Deutsche",
	Locale: &LngLocale{
		AbDay: []string{"So", "Mo", "Di", "Mi", "Do", "Fr", "Sa"},
		Day:   []string{"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"},
		AbMon: []string{"Jan", "Feb", "Mär", "Apr", "Mai", "Jun", "Jul", "Aug", "Sep", "Okt", "Nov", "Dez"},
		Mon:   []string{"Januar", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"},
		DTFmt: "%a %d %b %Y %T %Z",
		DFmt:  "%d.%m.%Y",
		TFmt:  "%T",
	},
}
