package tochka_test

import (
	"os"
	"testing"

	"github.com/sharpvik/tochka"
	"github.com/stretchr/testify/require"
)

func TestGetClosingDocumentPDF(t *testing.T) {
	config := tochka.Config{
		Token:        os.Getenv("JWT_TOKEN"),
		ClientID:     os.Getenv("CLIENT_ID"),
		CustomerCode: os.Getenv("CUSTOMER_CODE"),
		AccountID:    os.Getenv("ACCOUNT_ID"),
	}

	pdf, err := tochka.Live(config).
		GetClosingDocumentPDF("DOCUMENT_ID") // Replace with a valid document ID
	require.NoError(t, err)

	file, err := os.Create("closing_document.pdf")
	require.NoError(t, err)
	n, err := file.Write(pdf)
	require.NoError(t, err)
	require.Equal(t, len(pdf), n, "written bytes count should match PDF length")
}
