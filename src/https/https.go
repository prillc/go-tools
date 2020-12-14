package https

import (
	"io"
	"net/http"
	"net/http/cookiejar"
)

type params map[string]string

type session struct {
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

func DefaultSession() *session {

	s := &session{
		client:   defaultClient(),
		maxRetry: DefaultMaxRetry,
	}
	return s
}

func (s *session) common(method, uri string, ps params, body io.Reader) {
	s.resetLast()
	// 当前请求的状态设置状态
	url := GenUrl(uri, params{})
	req, err := http.NewRequest(method, url, body)

	s.request = req
	s.err = err
}

func (s *session) resetLast() {
	s.request = nil

	s.response = nil
	s.close = true

	s.retry = 0
}

func (s *session) Get(uri string, ps params) *session {
	s.common(http.MethodGet, uri, ps, nil)
	return s
}

func (s *session) Post(uri string, ps params, body io.Reader) {
	s.common(http.MethodPost, uri, ps, body)
}

func (s *session) GetClient() *http.Client {
	if s.client == nil {
		s.client = defaultClient()
	}
	return s.client
}

func (s *session) SetClient(client *http.Client) {
	s.client = client
}

func (s *session) Response() *http.Response {
	return s.response
}

func (s *session) setResponse(response *http.Response) {
	s.response = response
}
