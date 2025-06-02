package tochka_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sharpvik/tochka"
	"github.com/sharpvik/tochka/dto"
	"github.com/stretchr/testify/require"
)

var (
	customerCode     = "customerCode"
	accountID        = "accountID"
	oooYandexTaxCode = "7736207543"
	oooYandexKPP     = "770401001"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	customerCode = os.Getenv("CUSTOMER_CODE")
	accountID = os.Getenv("ACCOUNT_ID")
}

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

	pdf, err := sandbox.GetInvoice(customerCode, documentID)
	require.NoError(t, err)
	require.NotEmpty(t, pdf)

	// file, err := os.Create("invoice.pdf")
	// require.NoError(t, err)
	// _, err = file.Write(pdf)
	// require.NoError(t, err)
}

func TestCreateInvoiceAndGetItsPaymentStatus(t *testing.T) {
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

	status, err := sandbox.GetInvoicePaymentStatus(customerCode, documentID)
	require.NoError(t, err)
	require.Equal(t, dto.PaymentWaiting, status.Data.PaymentStatus)
}
