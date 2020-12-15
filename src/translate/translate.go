package translate

import "net/http"

type translate interface {
	Translate(text, src, dest string) *Translated
}

type Translated struct {
	// 输入和输出的字符串
	Text   string
	Result string

	// 输入的语言和输出的语言
	Src  string
	Dest string

	Err error

	Response *http.Response
}
