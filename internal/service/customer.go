package service

import (
	"context"
	"database/sql"
	"github.com/ZumaAkbarID/go-library-fiber/domain"
	"github.com/ZumaAkbarID/go-library-fiber/internal/dto"
	"github.com/google/uuid"
	"time"
)

type CustomerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
	}
}

func (c CustomerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var customerData []dto.CustomerData
	for _, r := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID:   r.ID,
			Code: r.Code,
			Name: r.Name,
		})
	}

	return customerData, nil
}

func (c CustomerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customers := domain.Customer{
		ID:        uuid.NewString(),
		Code:      req.Code,
		Name:      req.Name,
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	return c.customerRepository.Save(ctx, &customers)
}
