package lngdb

import "golang.org/x/text/language"

var SimplifiedChinese = &Lng{
	Tag:       language.SimplifiedChinese,
	LocalName: "汉语",
	Locale: &LngLocale{
		AbDay:    []string{"日", "一", "二", "三", "四", "五", "六"},
		Day:      []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
		AbMon:    []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		Mon:      []string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
		AmStr:    "上午",
		PmStr:    "下午",
		DTFmt:    "%Y年%m月%d日 %A %H时%M分%S秒",
		DFmt:     "%Y年%m月%d日",
		TFmt:     "%H时%M分%S秒",
		TFmtAmPm: "%p %I时%M分%S秒",
	},
}

var HongKongChinese = &Lng{
	Tag:       language.TraditionalChinese,
	LocalName: "漢語 (HK)",
	Locale: &LngLocale{
		AbDay:    []string{"日", "一", "二", "三", "四", "五", "六"},
		Day:      []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
		AbMon:    []string{"1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		Mon:      []string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
		AmStr:    "上午",
		PmStr:    "下午",
		DTFmt:    "%Y年%m月%d日 %A %H:%M:%S",
		DFmt:     "%Y年%m月%d日 %A",
		TFmt:     "%p %I時%M分%S秒 %Z",
		TFmtAmPm: "%p %I:%M:%S",
	},
}

var TaiwanChinese = &Lng{
	Tag:       language.TraditionalChinese,
	LocalName: "漢語 (TW)",
	Locale: &LngLocale{
		AbDay:    []string{"日", "一", "二", "三", "四", "五", "六"},
		Day:      []string{"週日", "週一", "週二", "週三", "週四", "週五", "週六"},
		AbMon:    []string{" 1月", " 2月", " 3月", " 4月", " 5月", " 6月", " 7月", " 8月", " 9月", "10月", "11月", "12月"},
		Mon:      []string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
		AmStr:    "上午",
		PmStr:    "下午",
		DTFmt:    "%Y年%m月%d日 (%A) %H時%M分%S秒",
		DFmt:     "%Y年%m月%d日",
		TFmt:     "%H時%M分%S秒",
		TFmtAmPm: "%p %I時%M分%S秒",
		Era:      "+:2:1913/01/01:+*:民國:%EC%Ey年",
	},
}
