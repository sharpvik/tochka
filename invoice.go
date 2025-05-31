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

	return
}
