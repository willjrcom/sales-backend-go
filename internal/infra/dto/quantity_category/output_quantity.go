package quantitydto

import (
	"github.com/google/uuid"
	productentity "github.com/willjrcom/sales-backend-go/internal/domain/product"
)

type QuantityOutput struct {
	ID       uuid.UUID `json:"id"`
	Quantity float64   `json:"quantity"`
}

func (s *QuantityOutput) FromModel(model *productentity.Quantity) {
	s.ID = model.ID
	s.Quantity = model.Quantity
}
