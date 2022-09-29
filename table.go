package lngdb

import "golang.org/x/text/language"

var Languages = map[string]*Lng{
	"de-DE": German,
	"en-CA": AustralianEnglish,
	"en-EU": EuropeanEnglish,
	"en-GB": BritishEnglish,
	"en-US": AmericanEnglish,
	"es-ES": Spanish,
	"es-MX": MexicanSpanish,
	"fr-FR": French,
	"hi-IN": Hindi,
	"it-IT": Italian,
	"ja-JP": Japanese,
	"ko-KR": Korean,
	"nl-NL": Dutch,
	"pl-PL": Polish,
	"pt-BR": BrazilianPortuguese,
	"pt-PT": Portuguese,
	"ru-RU": Russian,
	"sv-SE": Swedish,
	"th-TH": Thai,
	"zh-CN": SimplifiedChinese,
	"zh-HK": HongKongChinese,
	"zh-TW": TaiwanChinese,
}

var German = &Lng{
	Tag:       language.MustParse("de-DE"),
	LocalName: "Deutsche",
	Locale: &LngLocale{
		AbDay:    []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		Day:      []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		AbMon:    []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		Mon:      []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
		AmStr:    "AM",
		PmStr:    "PM",
		DTFmt:    "%a %b %e %H:%M:%S %Y",
		DFmt:     "%m/%d/%y",
		TFmt:     "%H:%M:%S",
		TFmtAmPm: "%I:%M:%S %p",
	},
}

var AustralianEnglish = &Lng{
	Tag:       language.MustParse("en-CA"),
	LocalName: "English (Canada)",
	Locale: &LngLocale{
		AbDay:    []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		Day:      []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		AbMon:    []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		Mon:      []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
		AmStr:    "AM",
		PmStr:    "PM",
		DTFmt:    "%a %b %e %H:%M:%S %Y",
		DFmt:     "%m/%d/%y",
		TFmt:     "%H:%M:%S",
		TFmtAmPm: "%I:%M:%S %p",
	},
}

var EuropeanEnglish = &Lng{
	Tag:       language.MustParse("en-EU"),
	LocalName: "English (EU)",
	Locale: &LngLocale{
		AbDay:    []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		Day:      []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		AbMon:    []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		Mon:      []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
		AmStr:    "AM",
		PmStr:    "PM",
		DTFmt:    "%a %b %e %H:%M:%S %Y",
		DFmt:     "%m/%d/%y",
		TFmt:     "%H:%M:%S",
		TFmtAmPm: "%I:%M:%S %p",
	},
}

var BritishEnglish = &Lng{
	Tag:       language.MustParse("en-GB"),
	LocalName: "English (UK)",
	Locale: &LngLocale{
		AbDay:    []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		Day:      []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		AbMon:    []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		Mon:      []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
		AmStr:    "AM",
		PmStr:    "PM",
		DTFmt:    "%a %b %e %H:%M:%S %Y",
		DFmt:     "%m/%d/%y",
		TFmt:     "%H:%M:%S",
		TFmtAmPm: "%I:%M:%S %p",
	},
}

var AmericanEnglish = &Lng{
	Tag:       language.MustParse("en-US"),
	LocalName: "English (US)",
	Locale: &LngLocale{
		AbDay:    []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		Day:      []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		AbMon:    []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		Mon:      []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
		AmStr:    "AM",
		PmStr:    "PM",
		DTFmt:    "%a %d %b %Y %r %Z",
		DFmt:     "%m/%d/%Y",
		TFmt:     "%r",
		TFmtAmPm: "%I:%M:%S %p",
	},
}

var Spanish = &Lng{
	Tag:       language.MustParse("es-ES"),
	LocalName: "Español",
	Locale: &LngLocale{
		AbDay:    []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		Day:      []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		AbMon:    []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		Mon:      []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
		AmStr:    "AM",
		PmStr:    "PM",
		DTFmt:    "%a %d %b %Y %r %Z",
		DFmt:     "%m/%d/%Y",
		TFmt:     "%r",
		TFmtAmPm: "%I:%M:%S %p",
	},
}

var MexicanSpanish = &Lng{
	Tag:       language.MustParse("es-MX"),
	LocalName: "Español (MX)",
	Locale: &LngLocale{
		AbDay:    []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"},
		Day:      []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
		AbMon:    []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"},
		Mon:      []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"},
		AmStr:    "AM",
		PmStr:    "PM",
		DTFmt:    "%a %d %b %Y %r %Z",
		DFmt:     "%m/%d/%Y",
		TFmt:     "%r",
		TFmtAmPm: "%I:%M:%S %p",
	},
}

var French = &Lng{
	Tag:       language.MustParse("fr-FR"),
	LocalName: "Français",
	Locale: &LngLocale{
		AbDay: []string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."},
		Day:   []string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"},
		AbMon: []string{"janv.", "févr.", "mars", "avril", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		Mon:   []string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
		DTFmt: "%a %d %b %Y %T",
		DFmt:  "%d/%m/%Y",
		TFmt:  "%T",
	},
}

var Hindi = &Lng{
	Tag:       language.MustParse("hi-IN"),
	LocalName: "हिन्दी",
	Locale: &LngLocale{
		AbDay: []string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."},
		Day:   []string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"},
		AbMon: []string{"janv.", "févr.", "mars", "avril", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		Mon:   []string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
		DTFmt: "%a %d %b %Y %T",
		DFmt:  "%d/%m/%Y",
		TFmt:  "%T",
	},
}

var Italian = &Lng{
	Tag:       language.MustParse("it-IT"),
	LocalName: "Italiano",
	Locale: &LngLocale{
		AbDay: []string{"dim.", "lun.", "mar.", "mer.", "jeu.", "ven.", "sam."},
		Day:   []string{"dimanche", "lundi", "mardi", "mercredi", "jeudi", "vendredi", "samedi"},
		AbMon: []string{"janv.", "févr.", "mars", "avril", "mai", "juin", "juil.", "août", "sept.", "oct.", "nov.", "déc."},
		Mon:   []string{"janvier", "février", "mars", "avril", "mai", "juin", "juillet", "août", "septembre", "octobre", "novembre", "décembre"},
		DTFmt: "%a %d %b %Y %T",
		DFmt:  "%d/%m/%Y",
		TFmt:  "%T",
	},
}

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
	},
}

var Korean = &Lng{
	Tag:       language.MustParse("ko-KR"),
	LocalName: "한국어",
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
	},
}

var Dutch = &Lng{
	Tag:       language.MustParse("nl-NL"),
	LocalName: "Nederlands",
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
	},
}

var Polish = &Lng{
	Tag:       language.MustParse("pl-PL"),
	LocalName: "język",
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
	},
}

var BrazilianPortuguese = &Lng{
	Tag:       language.MustParse("pt-BR"),
	LocalName: "Portuguesa (BR)",
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
	},
}

var Portuguese = &Lng{
	Tag:       language.MustParse("pt-PT"),
	LocalName: "Portuguesa",
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
	},
}

var Russian = &Lng{
	Tag:       language.MustParse("ru-RU"),
	LocalName: "ру́сский",
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
	},
}

var Swedish = &Lng{
	Tag:       language.MustParse("sv-SE"),
	LocalName: "Svenska",
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
	},
}

var Thai = &Lng{
	Tag:       language.MustParse("th-TH"),
	LocalName: "ภาษาไทย ",
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
	},
}

var SimplifiedChinese = &Lng{
	Tag:       language.MustParse("zh-CN"),
	LocalName: "汉语",
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
	},
}

var HongKongChinese = &Lng{
	Tag:       language.MustParse("zh-HK"),
	LocalName: "漢語 (HK)",
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
	},
}

var TaiwanChinese = &Lng{
	Tag:       language.MustParse("zh-TW"),
	LocalName: "漢語 (TW)",
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
	},
}
