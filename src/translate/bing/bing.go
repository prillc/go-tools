package bing

import (
	"fmt"
	"github.com/prillc/go-tools/src/https"
	"github.com/prillc/go-tools/src/translate"
	"io"
	"strings"
)

type bing struct {
	lang    string
	session *https.Session

	params params
}

func Default() *bing {
	session := https.DefaultSession()
	ps := params{
		session: session,
	}
	return &bing{
		session: session,
		params:  ps,
		lang:    CNLang,
	}
}

func (b *bing) Translate(text, src, dest string) *translate.Translated {
	// 初始化请求的参数
	b.initParams()

	translateUrl := b.translateUrl()
	formData := b.buildForm(text, src, dest)
	resultText, err := b.session.
		Post(translateUrl, b.params.requestParams(), formData).
		SetHeader("Content-Type", https.ContentTypeFormUrlEncode).
		SetHeader("User-Agent", https.DefaultUserAgent).
		Send().
		GetText()

	if err != nil {
		resultText = ""
	}

	translatedText, err := parseResult(resultText)

	t := &translate.Translated{
		Text:     text,
		Result:   translatedText,
		Src:      src,
		Dest:     dest,
		Err:      err,
		Response: b.session.Response(),
	}

	return t
}

func (b *bing) baseUrl() string {
	return BaseUrl(b.lang)
}

func (b *bing) translateUrl() string {
	return fmt.Sprintf("%s/%s", b.baseUrl(), "ttranslatev3")
}

func (b *bing) paramsUrl() string {
	return fmt.Sprintf("%s/%s", b.baseUrl(), "Translator")
}

func (b *bing) initParams() {
	b.params.SetUrl(b.paramsUrl())
	b.params.lazyInit()
}

func (b *bing) buildForm(text, src, dest string) io.Reader {
	s := fmt.Sprintf("fromLang=%s&text=%s&to=%s", src, text, dest)
	payload := strings.NewReader(s)
	return payload
}
