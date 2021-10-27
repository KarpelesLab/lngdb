package lngdb

import "golang.org/x/text/language"

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
