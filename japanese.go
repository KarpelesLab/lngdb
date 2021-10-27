package lngdb

import "golang.org/x/text/language"

var Japanese = &Lng{
	Tag:       language.MustParse("ja-JP"),
	LocalName: "日本語",
	Locale: &LngLocale{
		AbDay:    []string{"日", "月", "火", "水", "木", "金", "土"},
		Day:      []string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"},
		AbMon:    []string{" 1月", " 2月", " 3月", " 4月", " 5月", " 6月", " 7月", " 8月", " 9月", "10月", "11月", "12月"},
		Mon:      []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		AmStr:    "午前",
		PmStr:    "午後",
		DTFmt:    "%Y年%m月%d日 %H時%M分%S秒",
		DFmt:     "%Y年%m月%d日",
		TFmt:     "%H時%M分%S秒",
		TFmtAmPm: "%p%I時%M分%S秒",
		Era:      "+:2:2020/01/01:+*:令和:%EC%Ey年",
		EraDTFmt: "%EY%m月%d日 %H時%M分%S秒",
		EraDFmt:  "%EY%m月%d日",
		EraTFmt:  "",
	},
}
