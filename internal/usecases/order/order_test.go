package orderusecases

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	entitydto "github.com/willjrcom/sales-backend-go/internal/infra/dto/entity"
	orderdto "github.com/willjrcom/sales-backend-go/internal/infra/dto/order"
	orderrepositorylocal "github.com/willjrcom/sales-backend-go/internal/infra/repository/local/order"
)

var (
	service *Service
	ctx     context.Context
)

func TestMain(m *testing.M) {
	ctx = context.Background()
	repo := orderrepositorylocal.NewOrderRepositoryLocal()
	//address := addressentity.NewAddressRepositoryLocal()
	service = NewService(repo, nil)

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestRegisterOrder(t *testing.T) {
	idOrder, err := service.CreateDefaultOrder(ctx)
	assert.Nil(t, err)

	dtoId := entitydto.NewIdRequest(idOrder)
	Order, err := service.GetOrderById(ctx, dtoId)

	assert.Nil(t, err)
	assert.NotContains(t, idOrder, uuid.Nil)
	assert.Equal(t, Order.Name, "Test Order")
	assert.Equal(t, Order.ID, idOrder)
}

func TestUpdateOrder(t *testing.T) {

	Orders, err := service.GetAllOrders(ctx)
	assert.Nil(t, err)

	assert.Equal(t, len(Orders), 1)
	idOrder := Orders[0].ID

	dto := &orderdto.UpdateObservationOrder{Observation: "New observation"}
	dtoId := entitydto.NewIdRequest(idOrder)

	// Test 1 - New observation
	err = service.UpdateOrderObservation(ctx, dtoId, dto)
	assert.Nil(t, err)
}

func TestGetAll(t *testing.T) {
	Orders, err := service.GetAllOrders(ctx)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(Orders))
}

func TestGetOrderById(t *testing.T) {
	Orders, err := service.GetAllOrders(ctx)
	assert.Nil(t, err)
	assert.Equal(t, len(Orders), 1)
	idOrder := Orders[0].ID

	dtoId := entitydto.NewIdRequest(idOrder)
	Order, err := service.GetOrderById(ctx, dtoId)

	assert.Nil(t, err)
	assert.Equal(t, "Test Order", Order.Name)
}
