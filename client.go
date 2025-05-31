package tochka

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/sharpvik/tochka/dto"
)

type Client struct {
	resty *resty.Client
}

func Live(token, clientID string) *Client {
	return &Client{
		resty: rest(ProdURL, token).SetPathParam("clientId", clientID),
	}
}

func Sandbox() *Client {
	return &Client{
		resty: rest(SandboxURL, SandboxToken),
	}
}

func (c *Client) Modified(modify func(*resty.Client) *resty.Client) *Client {
	c.resty = modify(c.resty)
	return c
}

// UTIL

func rest(url, token string) *resty.Client {
	client := resty.New()
	client.BaseURL = url
	client.SetAuthToken(token)
	client.SetPathParam("apiVersion", Version)
	client.OnAfterResponse(onAfterResponse)
	return client
}

func onAfterResponse(c *resty.Client, r *resty.Response) error {
	if r.StatusCode() >= http.StatusBadRequest {
		return new(dto.ErrorResult).From(r.Body())
	}

	return nil
}
