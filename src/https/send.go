package https

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
)

func (s *session) send() error {
	if s.err != nil {
		return s.err
	}
	resp, err := s.GetClient().Do(s.request)

	s.setResponse(resp)
	s.err = err

	if s.retry == 0 {
		return err
	}

	return err
}

func (s *session) Send() *session {
	err := s.send()
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func (s *session) Close() error {
	if s.err != nil || s.response == nil {
		return s.err
	}

	err := s.response.Body.Close()
	if err != nil {
		s.err = err
	}
	return err
}

func (s *session) GetReader() (io.Reader, error) {
	if s.response == nil || s.response.Body == nil {
		return nil, errors.New("empty response")
	}
	return s.response.Body, nil
}

func (s *session) GetBytes() ([]byte, error) {
	reader, err := s.GetReader()
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(reader)
	// 关闭response.Body
	err = s.Close()
	if err != nil {
		return nil, err
	}

	return data, err
}

func (s *session) GetText() (string, error) {

	bytesData, err := s.GetBytes()
	if err != nil {
		return "", err
	}
	return string(bytesData), nil
}
