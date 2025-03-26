package types

type Language string

const (
	// Major world languages
	Language_en = "en" // English
	Language_es = "es" // Spanish
	Language_fr = "fr" // French
	Language_de = "de" // German
	Language_it = "it" // Italian
	Language_pt = "pt" // Portuguese
	Language_ru = "ru" // Russian
	Language_zh = "zh" // Chinese
	Language_ja = "ja" // Japanese
	Language_ko = "ko" // Korean
	Language_ar = "ar" // Arabic
	Language_hi = "hi" // Hindi

	// European languages
	Language_nl = "nl" // Dutch
	Language_sv = "sv" // Swedish
	Language_fi = "fi" // Finnish
	Language_da = "da" // Danish
	Language_no = "no" // Norwegian
	Language_pl = "pl" // Polish
	Language_cs = "cs" // Czech
	Language_sk = "sk" // Slovak
	Language_hu = "hu" // Hungarian
	Language_ro = "ro" // Romanian
	Language_el = "el" // Greek
	Language_tr = "tr" // Turkish
	Language_bg = "bg" // Bulgarian

	// Asian languages
	Language_th = "th" // Thai
	Language_vi = "vi" // Vietnamese
	Language_id = "id" // Indonesian
	Language_ms = "ms" // Malay
	Language_he = "he" // Hebrew
	Language_fa = "fa" // Persian

	// Other languages
	Language_af = "af" // Afrikaans
	Language_sq = "sq" // Albanian
	Language_eu = "eu" // Basque
	Language_ca = "ca" // Catalan
	Language_hr = "hr" // Croatian
	Language_et = "et" // Estonian
	Language_ga = "ga" // Irish
	Language_lv = "lv" // Latvian
	Language_lt = "lt" // Lithuanian
	Language_mt = "mt" // Maltese
	Language_sl = "sl" // Slovenian
	Language_sr = "sr" // Serbian
	Language_tl = "tl" // Tagalog

	// Language variants (ISO 639-1 + ISO 3166-1)
	Language_zh_cn = "zh-cn" // Chinese (Simplified)
	Language_zh_tw = "zh-tw" // Chinese (Traditional)
	Language_pt_br = "pt-br" // Portuguese (Brazil)
	Language_es_es = "es-es" // Spanish (Spain)
	Language_es_mx = "es-mx" // Spanish (Mexico)
	Language_fr_ca = "fr-ca" // French (Canada)
	Language_fr_fr = "fr-fr" // French (France)
	Language_de_de = "de-de" // German (Germany)
	Language_de_at = "de-at" // German (Austria)
	Language_de_ch = "de-ch" // German (Switzerland)
)
