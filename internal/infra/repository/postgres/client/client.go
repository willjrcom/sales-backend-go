package clientrepositorybun

import (
	"sync"

	"github.com/uptrace/bun"
	addressentity "github.com/willjrcom/sales-backend-go/internal/domain/address"
	cliententity "github.com/willjrcom/sales-backend-go/internal/domain/client"
	personentity "github.com/willjrcom/sales-backend-go/internal/domain/person"

	"golang.org/x/net/context"
)

type ClientRepositoryBun struct {
	mu sync.Mutex
	db *bun.DB
}

func NewClientRepositoryBun(db *bun.DB) *ClientRepositoryBun {
	return &ClientRepositoryBun{db: db}
}

func (r *ClientRepositoryBun) RegisterClient(ctx context.Context, c *cliententity.Client) error {
	r.mu.Lock()
	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	// Register client
	if _, err := tx.NewInsert().Model(c).Exec(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	// Register contacts
	for _, contact := range c.Contacts {
		if _, err := tx.NewInsert().Model(&contact).Exec(ctx); err != nil {
			return rollback(&tx, err)
		}
	}

	// Register addresses
	for _, address := range c.Addresses {
		if _, err := tx.NewInsert().Model(&address).Exec(ctx); err != nil {
			return rollback(&tx, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *ClientRepositoryBun) UpdateClient(ctx context.Context, c *cliententity.Client) error {
	r.mu.Lock()
	_, err := r.db.NewUpdate().Model(c).Where("id = ?", c.ID).Exec(ctx)
	r.mu.Unlock()

	if err != nil {
		return err
	}

	return nil
}

func (r *ClientRepositoryBun) DeleteClient(ctx context.Context, id string) error {
	r.mu.Lock()
	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	// Delete client
	if _, err = tx.NewDelete().Model(&cliententity.Client{}).Where("id = ?", id).Exec(ctx); err != nil {
		return rollback(&tx, err)
	}

	// Delete contacts
	if _, err = tx.NewDelete().Model(&personentity.Contact{}).Where("person_id = ?", id).Exec(ctx); err != nil {
		return rollback(&tx, err)
	}

	// Delete addresses
	if _, err = tx.NewDelete().Model(&addressentity.Address{}).Where("person_id = ?", id).Exec(ctx); err != nil {
		return rollback(&tx, err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *ClientRepositoryBun) GetClientById(ctx context.Context, id string) (*cliententity.Client, error) {
	client := &cliententity.Client{}

	r.mu.Lock()
	err := r.db.NewSelect().Model(client).Where("client.id = ?", id).Relation("Addresses").Relation("Contacts").Scan(ctx)
	r.mu.Unlock()

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (r *ClientRepositoryBun) GetClientsBy(ctx context.Context, c *cliententity.Client) ([]cliententity.Client, error) {
	clients := []cliententity.Client{}

	r.mu.Lock()
	query := r.db.NewSelect().Model(&cliententity.Client{})

	if c.Name != "" {
		query.Where("client.name LIKE ?", "%"+c.Name+"%")
	}

	if c.Cpf != "" {
		query.Where("client.cpf LIKE ?", "%"+c.Cpf+"%")
	}

	err := query.Relation("Addresses").Relation("Contacts").Scan(ctx, &clients)
	r.mu.Unlock()

	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (r *ClientRepositoryBun) GetAllClients(ctx context.Context) ([]cliententity.Client, error) {
	clients := []cliententity.Client{}
	r.mu.Lock()
	err := r.db.NewSelect().Model(&clients).Relation("Addresses").Relation("Contacts").Scan(ctx)

	r.mu.Unlock()

	if err != nil {
		return nil, err
	}

	return clients, nil
}

func rollback(tx *bun.Tx, err error) error {
	if err := tx.Rollback(); err != nil {
		return err
	}

	return err
}
