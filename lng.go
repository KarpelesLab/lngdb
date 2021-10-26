package lngdb

import "golang.org/x/text/language"

type Lng struct {
	language.Tag
	LocalName string `json:"Local_Name"` // name of language in the language itself
	Locale    *LngLocale
}

type LngLocale struct {
	AbDay    []string `json:"ABDAY"` // abbreviated days
	Day      []string `json:"DAY"`   // days
	AbMon    []string `json:"ABMON"` // abbreviated months
	Mon      []string `json:"MON"`   // months
	AmStr    string   `json:"AM_STR"`
	PmStr    string   `json:"PM_STR"`
	DTFmt    string   `json:"D_T_FMT"`
	DFmt     string   `json:"D_FMT"`
	TFmt     string   `json:"T_FMT"`
	TFmtAmPm string   `json:"T_FMT_AMPM"`
	Era      string   `json:"ERA,omitempty"`
	EraDTFmt string   `json:"ERA_D_T_FMT,omitempty"`
	EraDFmt  string   `json:"ERA_D_FMT,omitempty"`
	EraTFmt  string   `json:"ERA_T_FMT,omitempty"`
}
