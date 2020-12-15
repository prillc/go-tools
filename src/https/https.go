package https

import (
	"io"
	"net/http"
	"net/http/cookiejar"
)

type params map[string]string

type Session struct {
	client *http.Client

	// 最近一次的请求数据
	request *http.Request

	response *http.Response
	close    bool

	// 重试次数
	retry    uint
	maxRetry uint

	err error
}

func defaultClient() *http.Client {
	jar, _ := cookiejar.New(nil)
	return &http.Client{
		Jar: jar,
	}
}

func DefaultSession() *Session {

	s := &Session{
		client:   defaultClient(),
		maxRetry: DefaultMaxRetry,
	}
	return s
}

func (s *Session) common(method, uri string, ps params, body io.Reader) {
	s.resetLast()
	// 当前请求的状态设置状态
	url := GenUrl(uri, params{})
	req, err := http.NewRequest(method, url, body)

	s.request = req
	s.err = err
}

func (s *Session) resetLast() {
	s.request = nil

	s.response = nil
	s.close = true

	s.retry = 0
}

func (s *Session) Get(uri string, ps params) *Session {
	s.common(http.MethodGet, uri, ps, nil)
	return s
}

func (s *Session) Post(uri string, ps params, body io.Reader) *Session {
	s.common(http.MethodPost, uri, ps, body)
	return s
}

func (s *Session) GetClient() *http.Client {
	if s.client == nil {
		s.client = defaultClient()
	}
	return s.client
}

func (s *Session) SetClient(client *http.Client) {
	s.client = client
}

func (s *Session) Response() *http.Response {
	return s.response
}

func (s *Session) setResponse(response *http.Response) {
	s.response = response
}
