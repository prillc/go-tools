package test

import (
	"fmt"
	"github.com/prillc/go-tools/src/https"
	"testing"
)

func TestHttps(test *testing.T) {
	s := https.DefaultSession().Get("http://www.baidu.com", nil)
	fmt.Println(s)
	s.Send()

	t, _ := s.GetText()
	fmt.Println("text", t)
}
