package tochka

import (
	"strings"

	"github.com/sharpvik/tochka/dto"
)

func (c *Client) CreatePaymentForSign(data dto.CreatePaymentForSignData) (
	result dto.CreatePaymentForSignResult,
	err error,
) {
	if data.AccountCode == "" && data.BankCode == "" {
		data.AccountCode, data.BankCode, _ = strings.Cut(c.config.AccountID, "/")
	}

	params := dto.CreatePaymentForSignParams{Data: data}

	_, err = c.resty.R().
		SetBody(&params).
		SetResult(&result).
		Post("/payment/{apiVersion}/for-sign")

	return result, err
}

func (c *Client) GetPaymentStatus(requestID string) (
	result dto.GetPaymentStatusResult,
	err error,
) {
	_, err = c.resty.R().
		SetPathParam("requestId", requestID).
		SetResult(&result).
		Get("/payment/{apiVersion}/status/{requestId}")

	return result, err
}
