package model

type Parameter struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	FieldType string `json:"field_type"`
	Value     string `json:"value"`
}
