package test

import (
	"fmt"
	"github.com/prillc/go-tools/src/translate/bing"
	"testing"
)

func TestTranslate(test *testing.T) {
	b := bing.Default()
	r := b.Translate("这是啥，今天周二，天气晴", "auto-detect", "en")
	r1 := b.Translate("hello", "auto-detect", "zh-Hans")
	fmt.Println(r.Result, r1.Result)
}
