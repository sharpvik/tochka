package tochka_test

import (
	"testing"
	"time"

	"github.com/sharpvik/tochka/dto"
	"github.com/stretchr/testify/require"
)

const (
	sandboxAccountCode = "12345810901234567890"
	sandboxBankCode    = "044525104"
)

func createPaymentForSign(t *testing.T) dto.CreatePaymentForSignResult {
	moscow, err := time.LoadLocation("Europe/Moscow")
	require.NoError(t, err)

	params := dto.CreatePaymentForSignData{
		AccountCode:               sandboxAccountCode,
		BankCode:                  sandboxBankCode,
		CounterpartyBankBic:       sandboxBankCode,
		CounterpartyAccountNumber: "40817810802000000008",
		CounterpartyName:          "Тестовый Партнёр",
		CounterpartyINN:           "660000000000",
		PaymentAmount:             dto.KopeksFromRub(420),
		PaymentDate:               dto.Date(time.Now().In(moscow)),
		PaymentNumber:             dto.Natural(1),
		PaymentPriority:           "5",
		PaymentPurpose:            "Вознаграждение за привлечение пользователей. НДС не облагается.",
	}

	result, err := sandbox.CreatePaymentForSign(params)
	require.NoError(t, err)
	require.NotEmpty(t, result.Data.RequestID)
	require.NotEmpty(t, result.Data.RedirectURL)

	return result
}

func TestCreatePaymentForSign(t *testing.T) {
	createPaymentForSign(t)
}

func TestGetPaymentStatus(t *testing.T) {
	payment := createPaymentForSign(t)

	status, err := sandbox.GetPaymentStatus(payment.Data.RequestID)
	require.NoError(t, err)
	require.Equal(t, payment.Data.RequestID, status.Data.RequestID)
	require.Equal(t, dto.PaymentForSignWaitingForCreate, status.Data.Status)
	require.Empty(t, status.Data.Errors)
}

func TestGetPaymentStatusNotFound(t *testing.T) {
	_, err := sandbox.GetPaymentStatus("nonexistent-request-id")
	require.Error(t, err)
}
