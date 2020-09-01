package qr

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	apiMethod = "create-qr-code"
)

type Service struct {
	url     string
	version string
	ctx     context.Context
	client  *http.Client
}

func NewService(
	url string,
	version string,
	ctx context.Context,
	client *http.Client,
) *Service {
	return &Service{
		url: url, version: version, ctx: ctx, client: client,
	}
}

func (s *Service) Encode(line string) (data []byte, err error) {
	values := make(url.Values)
	reqURL := fmt.Sprintf("%s/%s/%s", s.url, s.version, apiMethod)
	values.Set("data", line)
	values.Set("size", "100x100")

	req, err := http.NewRequestWithContext(
		s.ctx,
		http.MethodGet,
		fmt.Sprintf("%s?%s", reqURL, values.Encode()),
		nil,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			if err == nil {
				log.Println(err)
				err = cerr
			}
		}
	}()

	return respBody, nil
}
