package orderentity

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
	employeeentity "github.com/willjrcom/sales-backend-go/internal/domain/employee"
	"github.com/willjrcom/sales-backend-go/internal/domain/entity"
	itementity "github.com/willjrcom/sales-backend-go/internal/domain/item"
)

type Order struct {
	entity.Entity
	OrderType
	OrderDetail
	bun.BaseModel `bun:"table:orders"`
	Payment       *PaymentOrder
	OrderNumber   int                    `bun:"order_number,notnull"`
	Status        StatusOrder            `bun:"status,notnull"`
	Groups        []itementity.GroupItem `bun:"rel:has-many,join:id=order_id"`
}

type OrderDetail struct {
	Name        string                   `bun:"name"`
	Observation string                   `bun:"observation"`
	AttendantID *uuid.UUID               `bun:"column:attendant_id,type:uuid"`
	Attendant   *employeeentity.Employee `bun:"rel:belongs-to"`
	LaunchedAt  *time.Time               `bun:"launch"`
}

type OrderType struct {
	Delivery *DeliveryOrder `bun:"rel:has-one,join:id=order_id"`
	Table    *TableOrder    `bun:"rel:has-one,join:id=order_id"`
}

func NewDefaultOrder() *Order {
	return &Order{
		Entity: entity.NewEntity(),
		Status: OrderStatusStaging,
	}
}

func (o *Order) StagingOrder() {
	o.Status = OrderStatusStaging
}

func (o *Order) LaunchOrder() {
	o.Status = OrderStatusPending
	*o.LaunchedAt = time.Now()
}

func (o *Order) ReadyOrder() {
	o.Status = OrderStatusFinished
}

func (o *Order) FinishOrder() {
	o.Status = OrderStatusFinished

	if o.Delivery != nil {
		*(*o.Delivery).Delivered = time.Now()
	}
}

func (o *Order) CancelOrder() {
	o.Status = OrderStatusCanceled
}

func (o *Order) ArchiveOrder() {
	o.Status = OrderStatusArchived
}
