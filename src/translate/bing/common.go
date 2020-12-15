package bing

func BaseUrl(lang string) string {
	// 分为中文和国际
	host := Url
	if lang == CNLang {
		host = CNUrl
	}
	return host
}
