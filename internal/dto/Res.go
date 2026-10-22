package dto

type Res struct{
	Status bool `json:"status"`
	Message string `json:"message"`
	Data any `json:"data,omitempty"`
}