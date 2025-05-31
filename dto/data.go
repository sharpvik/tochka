package dto

type DataParams[T any] struct {
	Data T `json:"Data"`
}
