package lngdb

import "golang.org/x/text/language"

var Korean = &Lng{
	Tag:       language.Korean,
	LocalName: "한국어",
	Locale: &LngLocale{
		AbDay:    []string{"일", "월", "화", "수", "목", "금", "토"},
		Day:      []string{"일요일", "월요일", "화요일", "수요일", "목요일", "금요일", "토요일"},
		AbMon:    []string{" 1월", " 2월", " 3월", " 4월", " 5월", " 6월", " 7월", " 8월", " 9월", "10월", "11월", "12월"},
		Mon:      []string{"1월", "2월", "3월", "4월", "5월", "6월", "7월", "8월", "9월", "10월", "11월", "12월"},
		AmStr:    "오전",
		PmStr:    "오후",
		DTFmt:    "%x (%a) %r",
		DFmt:     "%Y년 %m월 %d일",
		TFmt:     "%H시 %M분 %S초",
		TFmtAmPm: "%p %I시 %M분 %S초",
	},
}
