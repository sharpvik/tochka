package dto

import (
	"encoding/json"
	"fmt"
)

type CreateInvoiceParams DataParams[CreateInvoiceData]

type CreateInvoiceData struct {
	CustomerCode string                  `json:"customerCode"` // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Уникальный код клиента
	AccountID    string                  `json:"accountId"`    // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Уникальный и неизменный идентификатор счёта
	SecondSide   CreateInvoiceSecondSide `json:"SecondSide"`   // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ]
	Content      CreateInvoiceContent    `json:"Content"`      // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ]
}

type CreateInvoiceContent struct {
	Invoice CreateInvoiceInvoice `json:"Invoice"` // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ]
}

type CreateInvoiceInvoice struct {
	Positions         []CreateInvoicePosition `json:"Positions"`         // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ]
	Number            Natural                 `json:"number"`            // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Номер выставляемого счёта
	TotalAmount       Kopeks                  `json:"totalAmount"`       // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Сумма всех позиций с НДС
	BasedOn           string                  `json:"basedOn"`           // Документ, на основании которого выставляется счёт
	Comment           string                  `json:"comment"`           // Комментарий
	PaymentExpiryDate Date                    `json:"paymentExpiryDate"` // Срок оплаты в виде даты, приведенной к часовому поясу Москвы
	Date              Date                    `json:"date"`              // Дата выставления счета, приведенная к часовому поясу Москвы. Если не передана, то текущая дата
	TotalNDS          Kopeks                  `json:"totalNds"`          // Сумма НДС
}

type CreateInvoicePosition struct {
	PositionName string   `json:"positionName"` // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Название товара или услуги
	UnitCode     UnitCode `json:"unitCode"`     // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Единица измерения
	NDSKind      NDSKind  `json:"ndsKind"`      // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Тип НДС
	Price        Kopeks   `json:"price"`        // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Цена единицы с НДС
	Quantity     Quantity `json:"quantity"`     // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Количество
	TotalAmount  Kopeks   `json:"totalAmount"`  // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Сумма позиции с НДС
	TotalNDS     Kopeks   `json:"totalNds"`     // Сумма НДС
}

type CreateInvoiceSecondSide struct {
	TaxCode         string      `json:"taxCode"`         // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] ИНН покупателя или заказчика
	Type            CompanyType `json:"type"`            // [ОБЯЗАТЕЛЬНОЕ ПОЛЕ] Тип юр. лица
	AccountID       string      `json:"accountId"`       // Уникальный и неизменный идентификатор счёта
	LegalAddress    string      `json:"legalAddress"`    // Юридический адрес
	KPP             string      `json:"kpp"`             // КПП
	BankName        string      `json:"bankName"`        // Название банка
	BankCorrAccount string      `json:"bankCorrAccount"` // Корреспондентский счет банка
	Name            string      `json:"secondSideName"`  // Наименование покупателя или заказчика
}

type CreateInvoiceResult struct {
	Data struct {
		DocumentID string `json:"documentId"`
	} `json:"Data"`
	Links struct {
		Self string `json:"self"`
	} `json:"Links"`
	Meta struct {
		TotalPages int `json:"totalPages"`
	} `json:"Meta"`
}

type GetInvoicePaymentStatusResult struct {
	Data struct {
		PaymentStatus PaymentStatus `json:"paymentStatus"`
	} `json:"Data"`
	Links struct {
		Self string `json:"self"`
	} `json:"Links"`
	Meta struct {
		TotalPages int `json:"totalPages"`
	} `json:"Meta"`
}

type PaymentStatus string

const (
	PaymentWaiting PaymentStatus = "payment_waiting"
	PaymentExpired PaymentStatus = "payment_expired"
	PaymentPaid    PaymentStatus = "payment_paid"
)

var stringToPaymentStatusMap = map[string]PaymentStatus{
	"payment_waiting": PaymentWaiting,
	"payment_expired": PaymentExpired,
	"payment_paid":    PaymentPaid,
}

func (status *PaymentStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("failed to unmarshal PaymentStatus: %v", err)
	}

	mapped, found := stringToPaymentStatusMap[s]
	if !found {
		return fmt.Errorf("unknown string when unmarshaling PaymentStatus: %s", s)
	}

	*status = mapped

	return nil
}
