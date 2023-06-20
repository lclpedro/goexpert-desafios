package orders

import (
	"context"

	"github.com/lclpedro/scaffold-golang-fiber/internal/application/repositories/orders"
	"github.com/lclpedro/scaffold-golang-fiber/pkg/mysql"
)

type Service interface {
	GetAllOrders(ctx context.Context) ([]Output, error)
}

type ordersService struct {
	uow mysql.UnitOfWorkInterface
}

func NewOrdersService(uow mysql.UnitOfWorkInterface) Service {
	return &ordersService{
		uow: uow,
	}
}

func (h *ordersService) getOrdersRepository(ctx context.Context) (orders.Repository, error) {
	repo, err := h.uow.GetRepository(ctx, "OrdersRepository")
	if err != nil {
		return nil, err
	}
	return repo.(orders.Repository), nil
}

type Output struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (o *ordersService) GetAllOrders(ctx context.Context) ([]Output, error) {
	ordersRepo, err := o.getOrdersRepository(ctx)

	if err != nil {
		return []Output{}, err
	}
	_orders, err := ordersRepo.GetAll()
	if err != nil {
		return []Output{}, err
	}
	var output []Output
	for _, order := range _orders {
		output = append(output, Output{
			ID:    order.ID,
			Name:  order.Name,
			Price: order.Price,
		})
	}
	return output, nil
}