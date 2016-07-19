package evcard

import (
	"net/http"
)

type Session struct {
	EVcardClient *http.Client
}

func NewSession(ro *RequestOptions) *Session {
	if ro == nil {
		ro = &RequestOptions{}
	}

	ro.UseCookieJar = true
	return &Session{EVcardClient: GenerateHTTPClient(*ro)}
}

func (s *Session) Get(url string, ro *RequestOptions) (*Response, error) {
}
