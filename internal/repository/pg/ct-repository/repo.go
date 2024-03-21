package ctrepository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
	"github.com/sladkoezhkovo/admin-service/internal/repository/pg"
	ctservice "github.com/sladkoezhkovo/admin-service/internal/service/ct-service"
)

var _ ctservice.ConfectionaryTypeRepository = (*ctRepository)(nil)

type ctRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *ctRepository {
	return &ctRepository{
		db: db,
	}
}

func (c *ctRepository) Create(ct *entity.ConfectionaryType) error {
	return c.db.Get(
		ct,
		fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING *", pg.ConfectionaryTypeTable),
		ct.Name)
}

func (c *ctRepository) FindById(id int64) (*entity.ConfectionaryType, error) {
	var ct entity.ConfectionaryType
	if err := c.db.Get(
		&ct,
		fmt.Sprintf("SELECT * FROM %s WHERE id = $1", pg.ConfectionaryTypeTable),
		id,
	); err != nil {
		return nil, err
	}
	return &ct, nil
}

func (c *ctRepository) ListByName(name string, limit, offset int32) ([]*entity.ConfectionaryType, error) {
	var cc []*entity.ConfectionaryType
	if err := c.db.Select(
		&cc,
		fmt.Sprintf(`SELECT * FROM %s WHERE name ILIKE  $1  LIMIT $2 OFFSET $3`, pg.ConfectionaryTypeTable),
		"%"+name+"%", limit, offset,
	); err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *ctRepository) List(limit, offset int) ([]*entity.ConfectionaryType, int64, error) {
	var ctct []*entity.ConfectionaryType
	if err := c.db.Select(
		&ctct,
		fmt.Sprintf("SELECT * FROM %s ORDER BY id LIMIT $1 OFFSET $2", pg.ConfectionaryTypeTable),
		limit,
		offset,
	); err != nil {
		return nil, 0, err
	}
	var count int64
	if err := c.db.Get(&count, fmt.Sprintf("SELECT COUNT(id) FROM %s", pg.ConfectionaryTypeTable)); err != nil {
		return nil, 0, err
	}
	return ctct, count, nil
}

func (c *ctRepository) Update(ct *entity.ConfectionaryType) error {
	return c.db.Get(ct,
		fmt.Sprintf(`UPDATE %s
				SET
					name = $1
				WHERE id = $2 RETURNING *`, pg.ConfectionaryTypeTable),
		ct.Name,
		ct.Id,
	)
}

func (c *ctRepository) Delete(id int64) error {
	if _, err := c.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = $1", pg.ConfectionaryTypeTable), id); err != nil {
		return err
	}
	return nil
}
