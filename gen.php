#!/pkg/main/dev-lang.php.core.cli/bin/php
<?php

$fp = fopen('languages.csv', 'r');
$cols = fgetcsv($fp);

$locale_fields = [
	'ABDAY' => 'AbDay', 'DAY' => 'Day',
	'ABMON' => 'AbMon', 'MON' => 'Mon',
	'AM_STR' => 'AmStr', 'PM_STR' => 'PmStr',
	'D_T_FMT' => 'DTFmt',
	'D_FMT' => 'DFmt',
	'T_FMT' => 'TFmt',
	'T_FMT_AMPM' => 'TFmtAmPm',
	'ERA' => 'Era',
	'ERA_D_T_FMT' => 'EraDTFmt',
	'ERA_D_FMT' => 'EraDFmt',
	'ERA_T_FMT' => 'EraTFmt',
];

$output = ['table.go' => []];

while(!feof($fp)) {
	$tmp = fgetcsv($fp);
	if (!$tmp) continue;

	$lin = [];
	foreach($cols as $k => $v) $lin[$v] = $tmp[$k];

	setlocale(LC_TIME, str_replace('-', '_', $lin['token']).'.utf8');

	$locale = [];
	foreach($locale_fields as $k => $v) {
		if (defined($k)) {
			$locale[$v] = nl_langinfo(constant($k));
			continue;
		}
		if (defined($k.'_1')) {
			$n = 1;
			$x = [];
			while(defined($k.'_'.$n)) {
				$x[] = nl_langinfo(constant($k.'_'.$n));
				$n += 1;
			}
			$locale[$v] = $x;
		}
	}

	$output['table.go'][$lin['variable']] = [$lin, $locale];
}

foreach($output as $file => $lngs) {
	$fp = fopen($file, 'w');
	fwrite($fp, "package lngdb\n\nimport \"golang.org/x/text/language\"\n\n");

	foreach($lngs as $var => $info) {
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
}
