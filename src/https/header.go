package https

func (s *Session) SetHeader(key, value string) *Session {
	s.request.Header.Set(key, value)
	return s
}

func (s *Session) SetHeaders(headers map[string]string) *Session {
	for key, value := range headers {
		s.SetHeader(key, value)
	}
	return s
}
