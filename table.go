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
		AbDay: []string{"So", "Mo", "Di", "Mi", "Do", "Fr", "Sa"},
		Day:   []string{"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"},
		AbMon: []string{"Jan", "Feb", "Mär", "Apr", "Mai", "Jun", "Jul", "Aug", "Sep", "Okt", "Nov", "Dez"},
		Mon:   []string{"Januar", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"},
		DTFmt: "%a %d %b %Y %T %Z",
		DFmt:  "%d.%m.%Y",
		TFmt:  "%T",
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
		DTFmt:    "%a %d %b %Y %r",
		DFmt:     "%Y-%m-%d",
		TFmt:     "%r",
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
		DTFmt:    "%a %d %b %Y %r",
		DFmt:     "%Y-%m-%d",
		TFmt:     "%r",
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
		AmStr:    "am",
		PmStr:    "pm",
		DTFmt:    "%a %d %b %Y %T %Z",
		DFmt:     "%d/%m/%y",
		TFmt:     "%T",
		TFmtAmPm: "%l:%M:%S %P %Z",
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
		AbDay: []string{"dom", "lun", "mar", "mié", "jue", "vie", "sáb"},
		Day:   []string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"},
		AbMon: []string{"ene", "feb", "mar", "abr", "may", "jun", "jul", "ago", "sep", "oct", "nov", "dic"},
		Mon:   []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
		DTFmt: "%a %d %b %Y %T",
		DFmt:  "%d/%m/%y",
		TFmt:  "%T",
	},
}

var MexicanSpanish = &Lng{
	Tag:       language.MustParse("es-MX"),
	LocalName: "Español (MX)",
	Locale: &LngLocale{
		AbDay: []string{"dom", "lun", "mar", "mié", "jue", "vie", "sáb"},
		Day:   []string{"domingo", "lunes", "martes", "miércoles", "jueves", "viernes", "sábado"},
		AbMon: []string{"ene", "feb", "mar", "abr", "may", "jun", "jul", "ago", "sep", "oct", "nov", "dic"},
		Mon:   []string{"enero", "febrero", "marzo", "abril", "mayo", "junio", "julio", "agosto", "septiembre", "octubre", "noviembre", "diciembre"},
		DTFmt: "%a %d %b %Y %T",
		DFmt:  "%d/%m/%y",
		TFmt:  "%T",
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
		AbDay:    []string{"रवि", "सोम", "मंगल", "बुध", "गुरु", "शुक्र", "शनि"},
		Day:      []string{"रविवार", "सोमवार", "मंगलवार", "बुधवार", "गुरुवार", "शुक्रवार", "शनिवार"},
		AbMon:    []string{"जन॰", "फ़र॰", "मार्च", "अप्रैल", "मई", "जून", "जुल॰", "अग॰", "सित॰", "अक्तू॰", "नव॰", "दिस॰"},
		Mon:      []string{"जनवरी", "फ़रवरी", "मार्च", "अप्रैल", "मई", "जून", "जुलाई", "अगस्त", "सितंबर", "अक्तूबर", "नवंबर", "दिसंबर"},
		AmStr:    "पूर्वाह्न",
		PmStr:    "अपराह्न",
		DTFmt:    "%A %d %b %Y %I:%M:%S %p",
		DFmt:     "%-d/%-m/%y",
		TFmt:     "%I:%M:%S %p %Z",
		TFmtAmPm: "%I:%M:%S %p %Z",
	},
}

var Italian = &Lng{
	Tag:       language.MustParse("it-IT"),
	LocalName: "Italiano",
	Locale: &LngLocale{
		AbDay: []string{"dom", "lun", "mar", "mer", "gio", "ven", "sab"},
		Day:   []string{"domenica", "lunedì", "martedì", "mercoledì", "giovedì", "venerdì", "sabato"},
		AbMon: []string{"gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"},
		Mon:   []string{"gennaio", "febbraio", "marzo", "aprile", "maggio", "giugno", "luglio", "agosto", "settembre", "ottobre", "novembre", "dicembre"},
		DTFmt: "%a %-d %b %Y, %T",
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

var Dutch = &Lng{
	Tag:       language.MustParse("nl-NL"),
	LocalName: "Nederlands",
	Locale: &LngLocale{
		AbDay: []string{"zo", "ma", "di", "wo", "do", "vr", "za"},
		Day:   []string{"zondag", "maandag", "dinsdag", "woensdag", "donderdag", "vrijdag", "zaterdag"},
		AbMon: []string{"jan", "feb", "mrt", "apr", "mei", "jun", "jul", "aug", "sep", "okt", "nov", "dec"},
		Mon:   []string{"januari", "februari", "maart", "april", "mei", "juni", "juli", "augustus", "september", "oktober", "november", "december"},
		DTFmt: "%a %d %b %Y %T %Z",
		DFmt:  "%d-%m-%y",
		TFmt:  "%T",
	},
}

var Polish = &Lng{
	Tag:       language.MustParse("pl-PL"),
	LocalName: "język",
	Locale: &LngLocale{
		AbDay: []string{"nie", "pon", "wto", "śro", "czw", "pią", "sob"},
		Day:   []string{"niedziela", "poniedziałek", "wtorek", "środa", "czwartek", "piątek", "sobota"},
		AbMon: []string{"sty", "lut", "mar", "kwi", "maj", "cze", "lip", "sie", "wrz", "paź", "lis", "gru"},
		Mon:   []string{"stycznia", "lutego", "marca", "kwietnia", "maja", "czerwca", "lipca", "sierpnia", "września", "października", "listopada", "grudnia"},
		DTFmt: "%a, %-d %b %Y, %T",
		DFmt:  "%d.%m.%Y",
		TFmt:  "%T",
	},
}

var BrazilianPortuguese = &Lng{
	Tag:       language.MustParse("pt-BR"),
	LocalName: "Portuguesa (BR)",
	Locale: &LngLocale{
		AbDay: []string{"dom", "seg", "ter", "qua", "qui", "sex", "sáb"},
		Day:   []string{"domingo", "segunda", "terça", "quarta", "quinta", "sexta", "sábado"},
		AbMon: []string{"jan", "fev", "mar", "abr", "mai", "jun", "jul", "ago", "set", "out", "nov", "dez"},
		Mon:   []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"},
		DTFmt: "%a %d %b %Y %T",
		DFmt:  "%d/%m/%Y",
		TFmt:  "%T",
	},
}

var Portuguese = &Lng{
	Tag:       language.MustParse("pt-PT"),
	LocalName: "Portuguesa",
	Locale: &LngLocale{
		AbDay: []string{"dom", "seg", "ter", "qua", "qui", "sex", "sáb"},
		Day:   []string{"domingo", "segunda", "terça", "quarta", "quinta", "sexta", "sábado"},
		AbMon: []string{"jan", "fev", "mar", "abr", "mai", "jun", "jul", "ago", "set", "out", "nov", "dez"},
		Mon:   []string{"janeiro", "fevereiro", "março", "abril", "maio", "junho", "julho", "agosto", "setembro", "outubro", "novembro", "dezembro"},
		DTFmt: "%a %d %b %Y %T",
		DFmt:  "%d/%m/%Y",
		TFmt:  "%T",
	},
}

var Russian = &Lng{
	Tag:       language.MustParse("ru-RU"),
	LocalName: "ру́сский",
	Locale: &LngLocale{
		AbDay: []string{"Вс", "Пн", "Вт", "Ср", "Чт", "Пт", "Сб"},
		Day:   []string{"Воскресенье", "Понедельник", "Вторник", "Среда", "Четверг", "Пятница", "Суббота"},
		AbMon: []string{"янв", "фев", "мар", "апр", "мая", "июн", "июл", "авг", "сен", "окт", "ноя", "дек"},
		Mon:   []string{"января", "февраля", "марта", "апреля", "мая", "июня", "июля", "августа", "сентября", "октября", "ноября", "декабря"},
		DTFmt: "%a %d %b %Y %T",
		DFmt:  "%d.%m.%Y",
		TFmt:  "%T",
	},
}

var Swedish = &Lng{
	Tag:       language.MustParse("sv-SE"),
	LocalName: "Svenska",
	Locale: &LngLocale{
		AbDay: []string{"sön", "mån", "tis", "ons", "tor", "fre", "lör"},
		Day:   []string{"söndag", "måndag", "tisdag", "onsdag", "torsdag", "fredag", "lördag"},
		AbMon: []string{"jan", "feb", "mar", "apr", "maj", "jun", "jul", "aug", "sep", "okt", "nov", "dec"},
		Mon:   []string{"januari", "februari", "mars", "april", "maj", "juni", "juli", "augusti", "september", "oktober", "november", "december"},
		DTFmt: "%a %e %b %Y %H:%M:%S",
		DFmt:  "%Y-%m-%d",
		TFmt:  "%H:%M:%S",
	},
}

var Thai = &Lng{
	Tag:       language.MustParse("th-TH"),
	LocalName: "ภาษาไทย ",
	Locale: &LngLocale{
		AbDay:    []string{"อา.", "จ.", "อ.", "พ.", "พฤ.", "ศ.", "ส."},
		Day:      []string{"อาทิตย์", "จันทร์", "อังคาร", "พุธ", "พฤหัสบดี", "ศุกร์", "เสาร์"},
		AbMon:    []string{"ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."},
		Mon:      []string{"มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"},
		AmStr:    "AM",
		PmStr:    "PM",
		DTFmt:    "%a %e %b %Ey, %H:%M:%S",
		DFmt:     "%d/%m/%Ey",
		TFmt:     "%H:%M:%S",
		TFmtAmPm: "%I:%M:%S %p",
		Era:      "+:1:-543/01/01:+*:พ.ศ.:%EC %Ey",
		EraDTFmt: "วัน%Aที่ %e %B %EC %Ey, %H.%M.%S น.",
		EraDFmt:  "%e %b %Ey",
		EraTFmt:  "%H.%M.%S น.",
	},
}

var SimplifiedChinese = &Lng{
	Tag:       language.MustParse("zh-CN"),
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
	Tag:       language.MustParse("zh-HK"),
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
	Tag:       language.MustParse("zh-TW"),
	LocalName: "漢語 (TW)",
	Locale: &LngLocale{
		AbDay:    []string{"日", "一", "二", "三", "四", "五", "六"},
		Day:      []string{"週日", "週一", "週二", "週三", "週四", "週五", "週六"},
		AbMon:    []string{" 1月", " 2月", " 3月", " 4月", " 5月", " 6月", " 7月", " 8月", " 9月", "10月", "11月", "12月"},
		Mon:      []string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"},
		AmStr:    "上午",
		PmStr:    "下午",
		DTFmt:    "西元%Y年%m月%d日 (%A) %H時%M分%S秒",
		DFmt:     "西元%Y年%m月%d日",
		TFmt:     "%H時%M分%S秒",
		TFmtAmPm: "%p %I時%M分%S秒",
		Era:      "+:2:1913/01/01:+*:民國:%EC%Ey年",
	},
}
