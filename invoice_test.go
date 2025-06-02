package tochka_test

import (
	"os"
	"testing"

	"github.com/sharpvik/tochka"
	"github.com/sharpvik/tochka/dto"
	"github.com/stretchr/testify/require"
)

const (
	customerCode     = "customerCode"
	accountID        = "accountID"
	oooYandexTaxCode = "7736207543"
	oooYandexKPP     = "770401001"
)

func TestCreateAndGetInvoice(t *testing.T) {
	params := dto.CreateInvoiceParams{
		Data: dto.CreateInvoiceData{
			CustomerCode: customerCode,
			AccountID:    accountID,
			SecondSide: dto.CreateInvoiceSecondSide{
				TaxCode: oooYandexTaxCode,
				KPP:     oooYandexKPP,
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

	sandbox := tochka.Sandbox()

	result, err := sandbox.CreateInvoice(params)
	require.NoError(t, err)

	documentID := result.Data.DocumentID
	t.Log("DocumentID:", documentID)

	pdf, err := sandbox.GetInvoice(customerCode, documentID)
	require.NoError(t, err)
	file, err := os.Create("invoice.pdf")
	require.NoError(t, err)
	_, err = file.Write(pdf)
	require.NoError(t, err)
}
