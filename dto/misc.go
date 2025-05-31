package dto

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type CompanyType string

const (
	CompanyTypeIP      CompanyType = "ip"
	CompanyTypeCompany CompanyType = "company"
)

type UnitCode string

const (
	Pieces            UnitCode = "шт."
	ThousandsOfPieces UnitCode = "тыс.шт."
	Set               UnitCode = "компл."
	Couples           UnitCode = "пар."
	Units             UnitCode = "усл.ед."
	Packages          UnitCode = "упак."
	Services          UnitCode = "услуга."
	Packs             UnitCode = "пач."
	Minutes           UnitCode = "мин."
	Hours             UnitCode = "ч."
	Days              UnitCode = "сут."
	Grams             UnitCode = "г."
	Kilograms         UnitCode = "кг."
	Litres            UnitCode = "л."
	Meters            UnitCode = "м."
	SquareMeters      UnitCode = "м2."
	CubeMeters        UnitCode = "м3."
	Kilometers        UnitCode = "км."
	Hectares          UnitCode = "га."
	Kilowatts         UnitCode = "кВт."
	KilowattHours     UnitCode = "кВт.ч."
)

type NDSKind string

const (
	NDS0       NDSKind = "nds_0"
	NDS5       NDSKind = "nds_5"
	NDS7       NDSKind = "nds_7"
	NDS10      NDSKind = "nds_10"
	NDS20      NDSKind = "nds_20"
	WithoutNDS NDSKind = "without_nds"
)

type Date time.Time

func (date Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(date).Format(time.DateOnly))
}

type Natural uint

func (natural Natural) MarshalJSON() ([]byte, error) {
	return json.Marshal(strconv.FormatUint(uint64(natural), 10))
}

type Quantity float64

func (quantity Quantity) MarshalJSON() ([]byte, error) {
	if quantity < 0 {
		return nil, fmt.Errorf("Quantity(%f) < 0", quantity)
	}

	return json.Marshal(float64(quantity))
}
