package tochka_test

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/sharpvik/tochka"
	"github.com/sharpvik/tochka/dto"
	"github.com/stretchr/testify/require"
)

const (
	oooYandexTaxCode = "7736207543"
	oooYandexKPP     = "770401001"
)

var (
	config  tochka.Config
	sandbox *tochka.Client
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	config = tochka.Config{
		CustomerCode: os.Getenv("CUSTOMER_CODE"),
		AccountID:    os.Getenv("ACCOUNT_ID"),
	}

	sandbox = tochka.Sandbox(config)
}

func TestCreateInvoice(t *testing.T) {
	params := dto.CreateInvoiceData{
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
	}

	result, err := sandbox.CreateInvoice(params)
	require.NoError(t, err)
	require.NotEmpty(t, result.Data.DocumentID)
}

func TestGetInvoice(t *testing.T) {
	documentID := uuid.NewString()
	pdf, err := sandbox.GetInvoice(documentID)
	require.NoError(t, err)
	require.NotEmpty(t, pdf)
}

func TestGetInvoicePaymentStatus(t *testing.T) {
	documentID := uuid.NewString()
	status, err := sandbox.GetInvoicePaymentStatus(documentID)
	require.NoError(t, err)
	require.Equal(t, dto.PaymentWaiting, status.Data.PaymentStatus)
}
