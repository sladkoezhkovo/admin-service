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

func (c *ptRepository) ListByName(name string, limit, offset int32) ([]*entity.PropertyType, error) {
	var cc []*entity.PropertyType
	if err := c.db.Select(
		&cc,
		fmt.Sprintf(`SELECT * FROM %s WHERE name ILIKE  $1  LIMIT $2 OFFSET $3`, pg.PropertyTypeTable),
		"%"+name+"%", limit, offset,
	); err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *ptRepository) List(limit, offset int) ([]*entity.PropertyType, error) {
	var cities []*entity.PropertyType
	if err := c.db.Select(
		&cities,
		fmt.Sprintf("SELECT * FROM %s ORDER BY id LIMIT $1 OFFSET $2", pg.PropertyTypeTable),
		limit,
		offset,
	); err != nil {
		return nil, err
	}
	return cities, nil
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
