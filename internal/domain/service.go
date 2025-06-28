package domain

import (
	"strconv"
	"time"
)

type Service struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Interval string `json:"interval"`
	Strict   bool   `json:"strict"`
	Timeout  int    `json:"timeout"`
}

func (s Service) ParseInterval() (time.Duration, error) {
	if intervalNumber, err := strconv.Atoi(s.Interval); err == nil {
		return time.Duration(intervalNumber) * time.Millisecond, nil
	}

	return time.ParseDuration(s.Interval)
}

type ServiceError struct {
	Name       string `json:"name"`
	StatusCode int    `json:"statusCode"`
	Body       []byte `json:"body"`
	Error      error  `json:"error"`
}
