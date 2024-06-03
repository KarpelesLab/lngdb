#!/pkg/main/dev-lang.php.core.cli/bin/php
<?php

$fp = fopen('languages.csv', 'r');
$cols = fgetcsv($fp);

$locale_fields = [
	'LC_TIME.abday' => 'AbDay*',
	'LC_TIME.day' => 'Day*',
	'LC_TIME.abmon' => 'AbMon*',
	'LC_TIME.mon' => 'Mon*',
	'LC_TIME.am_pm' => 'AmStr;PmStr*',
	'LC_TIME.d_t_fmt' => 'DTFmt',
	'LC_TIME.d_fmt' => 'DFmt',
	'LC_TIME.t_fmt' => 'TFmt',
	'LC_TIME.t_fmt_ampm' => 'TFmtAmPm',
	'LC_TIME.era' => 'Era*',
	'LC_TIME.era_d_t_fmt' => 'EraDTFmt',
	'LC_TIME.era_d_fmt' => 'EraDFmt',
	'LC_TIME.era_t_fmt' => 'EraTFmt',
	'LC_MESSAGES.yesstr' => 'YesStr',
	'LC_MESSAGES.nostr' => 'NoStr',
	'LC_NUMERIC.decimal_point' => 'DecimalPoint',
	'LC_NUMERIC.thousands_sep' => 'ThousandsSep',
	'LC_ADDRESS.lang_name' => 'LangName',
	'LC_ADDRESS.country_name' => 'CountryName',
];

$pull = ['LC_TIME', 'LC_ADDRESS', 'LC_MESSAGES', 'LC_NUMERIC', 'LC_IDENTIFICATION'];

$output = [];

while(!feof($fp)) {
	$tmp = fgetcsv($fp);
	if (!$tmp) continue;

	$lin = [];
	foreach($cols as $k => $v) $lin[$v] = $tmp[$k];

	$locale_name = str_replace('-', '_', $lin['token']).'.utf8';
	$locale_data = [];
	if ($locale_name == 'en_EU.utf8') $locale_name = 'en_GB.utf8';

	echo "Fetching $locale_name\n";

	foreach($pull as $cat)
		$locale_data[$cat] = fetch_locale_info($locale_name, $cat);

	$locale = [];
	foreach($locale_fields as $k => $v) {
		$k = explode('.', $k);
		if (!isset($locale_data[$k[0]][$k[1]])) continue;
		$val = $locale_data[$k[0]][$k[1]];
		if (substr($v, -1) == '*') {
			// as array
			if (!is_array($val)) {
				if ($val === '') {
					$val = [];
				} else {
					$val = explode(';', $val);
				}
			}
			$v = substr($v, 0, -1);
		}
		$av = explode(';', $v);
		if (count($av) > 1) {
			foreach($av as $n => $sv) {
				$locale[$sv] = $val[$n];
			}
			continue;
		}
		$locale[$v] = $val;
	}

	$output[$lin['variable']] = [$lin, $locale];
}

$fp = fopen('table.go', 'w');
fwrite($fp, "package lngdb\n\nimport \"golang.org/x/text/language\"\n\n");

fwrite($fp, "var Languages = map[string]*Lng{\n");
foreach($output as $var => $info) {
	fwrite($fp, "\t\"".$info[0]['token']."\": $var,\n");
}
fwrite($fp, "}\n\n");

foreach($output as $var => $info) {
	$lin = $info[0];
	$locale = $info[1];
	fwrite($fp, 'var '.$var." = &Lng{\n");
	fwrite($fp, "\tTag:       language.MustParse(\"".$lin['token']."\"),\n");
	fwrite($fp, "\tLocalName: \"".$lin["local_name"]."\",\n");
	fwrite($fp, "\tLocale: &LngLocale{\n");
	foreach($locale as $key => $val) {
		if (!$val) continue;
		if (is_array($val)) {
			$val = '[]string{"'.implode('", "', $val).'"}';
		} else {
			$val = '"'.$val.'"';
		}
		fwrite($fp, "\t\t".$key.': '.$val.",\n");
	}
	fwrite($fp, "\t},\n");
	fwrite($fp, "}\n\n");
}

function fetch_locale_info($locale_name, $cat) {
	$res = shell_exec('LANG='.escapeshellarg($locale_name).' /pkg/main/sys-libs.glibc.core/bin/locale -k '.escapeshellarg($cat));
	$res = explode("\n", $res);
	$final = [];
	foreach($res as $lin) {
		$lin = trim($lin);
		if ($lin == '') continue;

		// a=
		// a=""
		// a="..."
		// a="...;...;..."
		// a="...";"...";"..."
		$pos = strpos($lin, '=');
		if ($pos === false) continue; // should not happen
		$var = substr($lin, 0, $pos);
		$lin = trim(substr($lin, $pos+1));
		if ($lin == '') {
			$final[$var] = $lin;
			continue;
		}
		if ($lin[0] != '"') {
			$final[$var] = $lin;
			continue;
		}

		$sub = [];
		while(strlen($lin) > 0) {
			$lin = substr($lin, 1);
			$pos = strpos($lin, '"');
			if ($pos === false) throw new \Exception('badly formatted var');
			$sub[] = substr($lin, 0, $pos);
			$lin = substr($lin, $pos+1);
			if ($lin === '') break;
			if ($lin[0] != ';') throw new \Exception('var to var missing ";"');
			$lin = substr($lin, 1);
		}
		if (count($sub) == 1) {
			$final[$var] = $sub[0];
		} else {
			$final[$var] = $sub;
		}
	}
	return $final;
}
