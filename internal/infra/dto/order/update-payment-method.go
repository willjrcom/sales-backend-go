package orderdto

import orderentity "github.com/willjrcom/sales-backend-go/internal/domain/order"

type UpdatePaymentMethod struct {
	IsPaid    *bool                  `json:"is_paid"`
	TotalPaid *float64               `json:"total_paid"`
	Change    *float64               `json:"change"`
	Method    *orderentity.PayMethod `json:"method"`
}

func (u *UpdatePaymentMethod) UpdateModel(order *orderentity.Order) {
	if u.IsPaid != nil {
		order.IsPaid = *u.IsPaid
	}
	if u.TotalPaid != nil {
		*order.TotalPaid = *u.TotalPaid
	}
	if u.Change != nil {
		*order.Change = *u.Change
	}
	if u.Method != nil {
		*order.Method = *u.Method
	}
}
