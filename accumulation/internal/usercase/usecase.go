package usecase

import "accumulation_consumer/internal/domain"

type HttpClientCase struct {
	HTTPClient domain.HTTPClient
}

func (u *HttpClientCase) Post(url string, body []byte) {
	u.HTTPClient.Post(url, body)
}
