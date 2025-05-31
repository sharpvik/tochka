package tochka_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/sharpvik/tochka"
	"github.com/sharpvik/tochka/dto"
	"github.com/stretchr/testify/assert"
)

func TestCreateInvoice(t *testing.T) {
	params := dto.CreateInvoiceParams{
		Data: dto.CreateInvoiceData{
			CustomerCode: "CustomerCode",
			AccountID:    "AccountID",
			SecondSide: dto.CreateInvoiceSecondSide{
				TaxCode: "TaxCode",
				Type:    dto.CompanyTypeCompany,
			},
			Content: dto.CreateInvoiceContent{
				Invoice: dto.CreateInvoiceInvoice{
					Number:      dto.Natural(1),
					TotalAmount: dto.KopeksFromRub(420),
					Positions: []dto.CreateInvoicePosition{
						{
							PositionName: "PositionName",
							UnitCode:     dto.Units,
							NDSKind:      dto.WithoutNDS,
							Price:        dto.KopeksFromRub(1),
							Quantity:     dto.Quantity(420.00),
							TotalAmount:  dto.KopeksFromRub(420),
						},
					},
				},
			},
		},
	}

	_, err := tochka.Sandbox().CreateInvoice(params)
	assert.Error(t, err)
	t.Log(err)

	result, ok := err.(*dto.ErrorResult)
	assert.True(t, ok)
	assert.Equal(t, strconv.Itoa(http.StatusBadRequest), result.Code)
}
