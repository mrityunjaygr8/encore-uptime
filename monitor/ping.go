package monitor

import (
	"context"
	"net/http"
	"strings"
)

type PingResponse struct {
	Up bool `json:"up"`
}

//encore:api public path=/ping/*url
func Ping(ctx context.Context, url string) (*PingResponse, error) {
	if !strings.HasPrefix(url, "http:") && !strings.HasPrefix(url, "https:") {
		url = "https://" + url
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &PingResponse{Up: false}, err
	}

	resp.Body.Close()

	up := resp.StatusCode < 400
	return &PingResponse{Up: up}, nil
}
