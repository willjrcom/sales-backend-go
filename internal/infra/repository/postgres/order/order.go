package orderrepositorybun

import (
	"context"
	"database/sql"
	"sync"

	"github.com/uptrace/bun"
	orderentity "github.com/willjrcom/sales-backend-go/internal/domain/order"
)

type OrderRepositoryBun struct {
	mu sync.Mutex
	db *bun.DB
}

func NewOrderRepositoryBun(db *bun.DB) *OrderRepositoryBun {
	return &OrderRepositoryBun{db: db}
}

func (r *OrderRepositoryBun) CreateOrder(ctx context.Context, order *orderentity.Order) error {
	r.mu.Lock()
	_, err := r.db.NewInsert().Model(order).Exec(ctx)
	r.mu.Unlock()

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepositoryBun) PendingOrder(ctx context.Context, p *orderentity.Order) error {
	tx, err := r.db.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		return err
	}

	r.mu.Lock()
	_, err = tx.NewUpdate().Model(p).Where("id = ?", p.ID).Exec(ctx)
	r.mu.Unlock()

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, group := range p.Groups {
		_, err = tx.NewUpdate().Model(&group).WherePK().Exec(ctx)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *OrderRepositoryBun) UpdateOrder(ctx context.Context, order *orderentity.Order) error {
	r.mu.Lock()
	_, err := r.db.NewUpdate().Model(order).Where("id = ?", order.ID).Exec(ctx)
	r.mu.Unlock()

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepositoryBun) DeleteOrder(ctx context.Context, id string) error {
	r.mu.Lock()
	_, err := r.db.NewDelete().Model(&orderentity.Order{}).Where("id = ?", id).Exec(ctx)
	r.mu.Unlock()

	if err != nil {
		return err
	}

	return nil
}

func (r *OrderRepositoryBun) GetOrderById(ctx context.Context, id string) (*orderentity.Order, error) {
	order := &orderentity.Order{}

	r.mu.Lock()
	err := r.db.NewSelect().Model(order).Where("id = ?", id).Relation("Groups.Items").Scan(ctx)
	r.mu.Unlock()
	// .Relation("Delivery").Relation("Table").Relation("Groups").Relation("Attendant")

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepositoryBun) GetAllOrders(ctx context.Context) ([]orderentity.Order, error) {
	orders := []orderentity.Order{}
	r.mu.Lock()
	err := r.db.NewSelect().Model(&orders).Relation("Groups.Items").Scan(ctx)

	r.mu.Unlock()

	if err != nil {
		return nil, err
	}

	return orders, nil
}
