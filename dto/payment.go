package dto

type CreatePaymentForSignParams DataParams[CreatePaymentForSignData]

type CreatePaymentForSignData struct {
	AccountCode               string  `json:"accountCode"`               // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Номер счёта плательщика
	BankCode                  string  `json:"bankCode"`                  // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] БИК банка плательщика
	CounterpartyBankBic       string  `json:"counterpartyBankBic"`       // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] БИК банка получателя
	CounterpartyAccountNumber string  `json:"counterpartyAccountNumber"` // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Номер счёта получателя
	CounterpartyName          string  `json:"counterpartyName"`          // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Наименование получателя
	CounterpartyINN           string  `json:"counterpartyINN,omitempty"` // ИНН получателя
	CounterpartyKPP           string  `json:"counterpartyKPP,omitempty"` // КПП получателя
	PaymentAmount             Kopeks  `json:"paymentAmount"`             // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Сумма платежа
	PaymentDate               Date    `json:"paymentDate"`               // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Дата платёжного документа: по Москве и не позже сегодня
	PaymentNumber             Natural `json:"paymentNumber,omitempty"`   // Номер платёжного документа
	PaymentPriority           string  `json:"paymentPriority,omitempty"` // Очерёдность платежа
	PaymentPurpose            string  `json:"paymentPurpose"`            // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Назначение платежа
	CodePurpose               string  `json:"codePurpose,omitempty"`     // Код вида дохода (для выплат физлицам на счета 40817)
}

type CreatePaymentForSignResult struct {
	Data struct {
		RequestID   string `json:"requestId"`
		RedirectURL string `json:"redirectURL"`
	} `json:"Data"`
	Links struct {
		Self string `json:"self"`
	} `json:"Links"`
	Meta struct {
		TotalPages int `json:"totalPages"`
	} `json:"Meta"`
}

type GetPaymentStatusResult struct {
	Data struct {
		RequestID string               `json:"requestId"`
		Status    PaymentForSignStatus `json:"status"`
		Errors    []string             `json:"errors"`
	} `json:"Data"`
	Links struct {
		Self string `json:"self"`
	} `json:"Links"`
	Meta struct {
		TotalPages int `json:"totalPages"`
	} `json:"Meta"`
}

// PaymentForSignStatus is deliberately a plain string: Tochka is known to
// return undocumented statuses, and a client used for monitoring must not
// fail to parse a response because of an unfamiliar value.
type PaymentForSignStatus string

const (
	PaymentForSignInitiated        PaymentForSignStatus = "Initiated"
	PaymentForSignWaitingForCreate PaymentForSignStatus = "WaitingForCreate"
	PaymentForSignCreated          PaymentForSignStatus = "Created"
	PaymentForSignPaid             PaymentForSignStatus = "Paid"
	PaymentForSignRejected         PaymentForSignStatus = "Rejected"
	PaymentForSignCanceled         PaymentForSignStatus = "Canceled"
)
