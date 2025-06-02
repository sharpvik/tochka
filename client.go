package tochka

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/sharpvik/tochka/dto"
)

type Client struct {
	config *Config
	resty  *resty.Client
}

type Config struct {
	Token        string
	ClientID     string
	CustomerCode string
	AccountID    string
}

func Live(config Config) *Client {
	return New(config, ProdURL)
}

func Sandbox(config Config) *Client {
	config.Token = SandboxToken
	return New(config, SandboxURL)
}

func New(config Config, url string) *Client {
	return &Client{
		config: &config,
		resty: rest(url, config.Token).
			SetPathParams(map[string]string{
				"clientId":     config.ClientID,
				"customerCode": config.CustomerCode,
			}),
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
