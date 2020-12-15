package bing

import (
	"github.com/prillc/go-tools/src/https"
	"regexp"
)

var IIDPattern, _ = regexp.Compile("data-iid=\"(.*?)\"")
var IGPattern, _ = regexp.Compile("{.*?IG:\"(.*?)\"")

type params struct {
	// 获取iig和ig的路由地址
	url     string
	session *https.Session

	ig  string
	iid string

	isInit bool
}

func (p *params) lazyInit() {
	// 初始化参数
	if p.isInit {
		return
	}
	_ = p.downloadParams()
}

func (p *params) downloadParams() error {
	text, err := p.session.Get(p.url, nil).Send().GetText()
	if err != nil {
		return err
	}

	iidMatches := IIDPattern.FindStringSubmatch(text)
	igMatches := IGPattern.FindStringSubmatch(text)

	if len(iidMatches) > 0 {
		p.iid = iidMatches[1]
	}

	if len(igMatches) > 0 {
		p.ig = igMatches[1]
	}

	return nil
}

func (p *params) SetUrl(url string) {
	p.url = url
}

func (p *params) requestParams() map[string]string {
	return map[string]string{
		"IID": p.iid,
		"IG":  p.ig,
	}
}
