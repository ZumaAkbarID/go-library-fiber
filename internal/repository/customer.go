package repository

import (
	"context"
	"database/sql"
	"github.com/ZumaAkbarID/go-library-fiber/domain"
	"github.com/doug-martin/goqu/v9"
	"time"
)

type customerRepository struct {
	db *goqu.Database
}

func NewCustomer(con *sql.DB) domain.CustomerRepository {
	return &customerRepository{
		db: goqu.New("default", con),
	}
}

func (cr customerRepository) FindAll(ctx context.Context) (result []domain.Customer, err error) {
	dataset := cr.db.From("customers").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

func (cr customerRepository) FindByID(ctx context.Context, id string) (result domain.Customer, err error) {
	dataset := cr.db.From("customers").Where(goqu.C("deleted_at").IsNull(), goqu.C("id").Eq(id))
	_, err = dataset.ScanStructContext(ctx, &result)
	return
}

func (cr customerRepository) Save(ctx context.Context, c *domain.Customer) error {
	executor := cr.db.Insert("customers").Rows(c).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (cr customerRepository) Update(ctx context.Context, c *domain.Customer) error {
	executor := cr.db.Update("customers").Where(goqu.C("id").Eq(c.ID)).Set(c).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

func (cr customerRepository) Delete(ctx context.Context, id string) error {
	executor := cr.db.Update("customers").
		Where(goqu.C("id").Eq(id)).
		Set(goqu.Record{"deleted_at": sql.NullTime{Time: time.Now(), Valid: true}}).
		Executor()
	_, err := executor.ExecContext(ctx)
	return err
}
