package ptrepository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
	"github.com/sladkoezhkovo/admin-service/internal/repository/pg"
	ptservice "github.com/sladkoezhkovo/admin-service/internal/service/pt-service"
)

var _ ptservice.PropertyTypeRepository = (*ptRepository)(nil)

type ptRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *ptRepository {
	return &ptRepository{
		db: db,
	}
}

func (c *ptRepository) Create(pt *entity.PropertyType) error {
	return c.db.Get(
		pt,
		fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING *", pg.PropertyTypeTable),
		pt.Name)
}

func (c *ptRepository) FindById(id int64) (*entity.PropertyType, error) {
	var pt entity.PropertyType
	if err := c.db.Get(
		&pt,
		fmt.Sprintf("SELECT * FROM %s WHERE id = $1", pg.PropertyTypeTable),
		id,
	); err != nil {
		return nil, err
	}
	return &pt, nil
}

func (c *ptRepository) List(limit, offset int) ([]*entity.PropertyType, int64, error) {
	var ptpt []*entity.PropertyType
	if err := c.db.Select(
		&ptpt,
		fmt.Sprintf("SELECT * FROM %s ORDER BY id LIMIT $1 OFFSET $2", pg.PropertyTypeTable),
		limit,
		offset,
	); err != nil {
		return nil, 0, err
	}
	var count int64
	if err := c.db.Get(&count, fmt.Sprintf("SELECT COUNT(id) FROM %s", pg.PropertyTypeTable)); err != nil {
		return nil, 0, err
	}
	return ptpt, count, nil
}

func (c *ptRepository) Update(pt *entity.PropertyType) error {
	return c.db.Get(pt,
		fmt.Sprintf(`UPDATE %s
				SET
					name = $1
				WHERE id = $2 RETURNING *`, pg.PropertyTypeTable),
		pt.Name,
		pt.Id,
	)
}

func (c *ptRepository) Delete(id int64) error {
	if _, err := c.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = $1", pg.PropertyTypeTable), id); err != nil {
		return err
	}
	return nil
}
