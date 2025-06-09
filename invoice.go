package tochka

import (
	"github.com/sharpvik/tochka/dto"
)

func (c *Client) CreateInvoice(data dto.CreateInvoiceData) (
	result dto.CreateInvoiceResult,
	err error,
) {
	if data.AccountID == "" {
		data.AccountID = c.config.AccountID
	}

	if data.CustomerCode == "" {
		data.CustomerCode = c.config.CustomerCode
	}

	params := dto.CreateInvoiceParams{Data: data}

	_, err = c.resty.R().
		SetBody(&params).
		SetResult(&result).
		Post("/invoice/{apiVersion}/bills")

	return result, err
}

func (c *Client) GetInvoicePDF(documentID string) (
	pdf []byte,
	err error,
) {
	resp, err := c.resty.R().
		SetPathParam("documentId", documentID).
		Get("/invoice/{apiVersion}/bills/{customerCode}/{documentId}/file")

	return resp.Body(), err
}

func (c *Client) GetInvoicePaymentStatus(documentID string) (
	result dto.GetInvoicePaymentStatusResult,
	err error,
) {
	_, err = c.resty.R().
		SetPathParam("documentId", documentID).
		SetResult(&result).
		Get("/invoice/{apiVersion}/bills/{customerCode}/{documentId}/payment-status")

	return result, err
}

func (c *Client) DeleteInvoice(documentID string) (err error) {
	_, err = c.resty.R().
		SetPathParam("documentId", documentID).
		Get("/invoice/{apiVersion}/bills/{customerCode}/{documentId}")

	return err
}
