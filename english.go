package lngdb

import "golang.org/x/text/language"

var AmericanEnglish = &Lng{
	Tag:       language.AmericanEnglish,
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

var BritishEnglish = &Lng{
	Tag:       language.BritishEnglish,
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
