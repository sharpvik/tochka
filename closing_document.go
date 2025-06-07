package tochka

func (c *Client) GetClosingDocumentPDF(documentID string) (
	pdf []byte,
	err error,
) {
	resp, err := c.resty.R().
		SetPathParam("documentId", documentID).
		Get("/invoice/{apiVersion}/closing-documents/{customerCode}/{documentId}/file")

	return resp.Body(), err
}
