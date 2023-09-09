package me

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	httpclient "github.com/DiLRandI/circle-ci/internal/httpClient"
	"github.com/DiLRandI/circle-ci/internal/service/common"
)

const (
	path = "me"
)

type Service interface {
	GetMe(context.Context) (*Me, error)
}

type service struct {
	client httpclient.Client
}

func New(client httpclient.Client) Service {
	return &service{
		client: client,
	}
}

func (s *service) GetMe(ctx context.Context) (*Me, error) {
	url := common.Endpoint + path
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create the request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do the request: %w", err)
	}
	defer resp.Body.Close()

	var meDto Me

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the response body: %w", err)
	}

	if err := json.Unmarshal(body, &meDto); err != nil {
		return nil, fmt.Errorf("failed to unmarshal the response body: %w", err)
	}

	return &meDto, nil
}
