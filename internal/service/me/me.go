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
	logger, token, err := common.GetLoggerAndToken(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get logger and token: %w", err)
	}

	url := common.Endpoint + path
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create the request: %w", err)
	}

	req.Header.Add("Circle-Token", token)

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do the request: %w", err)
	}
	defer resp.Body.Close()

	logger.Debug("Response received", "status", resp.Status, "headers", resp.Header)

	var meDto Me

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read the response body: %w", err)
	}

	logger.Debug("Response body", "body", string(body))

	if err := json.Unmarshal(body, &meDto); err != nil {
		return nil, fmt.Errorf("failed to unmarshal the response body: %w", err)
	}

	return &meDto, nil
}
