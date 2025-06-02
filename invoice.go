package tochka

import (
	"github.com/sharpvik/tochka/dto"
)

func (c *Client) CreateInvoice(params dto.CreateInvoiceParams) (
	result dto.CreateInvoiceResult,
	err error,
) {
	_, err = c.resty.R().
		SetBody(&params).
		SetResult(&result).
		Post("/invoice/{apiVersion}/bills")

	return result, err
}

func (c *Client) GetInvoice(customerCode, documentID string) (
	pdf []byte,
	err error,
) {
	resp, err := c.resty.R().
		SetPathParams(map[string]string{
			"customerCode": customerCode,
			"documentId":   documentID,
		}).
		Get("/invoice/{apiVersion}/bills/{customerCode}/{documentId}/file")

	return resp.Body(), err
}

func (c *Client) GetInvoicePaymentStatus(customerCode, documentID string) (
	result dto.GetInvoicePaymentStatusResult,
	err error,
) {
	_, err = c.resty.R().
		SetPathParams(map[string]string{
			"customerCode": customerCode,
			"documentId":   documentID,
		}).
		SetResult(&result).
		Get("invoice/{apiVersion}/bills/{customerCode}/{documentId}/payment-status")

	return result, err
}
