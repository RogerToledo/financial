package dto

import "github.com/me/finance/pkg/entity"

type InstallmentResponse struct{
	Response []entity.Installment `json:"response"`
	Paid  float64 `json:"paid"`
	ToPay float64 `json:"to_pay"`
	Total float64 `json:"total"`
}