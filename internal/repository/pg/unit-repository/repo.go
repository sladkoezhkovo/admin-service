package unitrepository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
	"github.com/sladkoezhkovo/admin-service/internal/repository/pg"
	unitservice "github.com/sladkoezhkovo/admin-service/internal/service/unit-service"
)

var _ unitservice.UnitRepository = (*unitRepository)(nil)

type unitRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *unitRepository {
	return &unitRepository{
		db: db,
	}
}

func (c *unitRepository) Create(unit *entity.Unit) error {
	return c.db.Get(
		unit,
		fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING *", pg.UnitTable),
		unit.Name)
}

func (c *unitRepository) FindById(id int64) (*entity.Unit, error) {
	var unit entity.Unit
	if err := c.db.Get(
		&unit,
		fmt.Sprintf("SELECT * FROM %s WHERE id = $1", pg.UnitTable),
		id,
	); err != nil {
		return nil, err
	}
	return &unit, nil
}

func (c *unitRepository) FindByName(name string) (*entity.Unit, error) {
	var unit entity.Unit
	if err := c.db.Get(
		&unit,
		fmt.Sprintf("SELECT * FROM %s WHERE name = $1", pg.UnitTable),
		name,
	); err != nil {
		return nil, err
	}
	return &unit, nil
}

func (c *unitRepository) List(limit, offset int) ([]*entity.Unit, error) {
	var cities []*entity.Unit
	if err := c.db.Select(
		&cities,
		fmt.Sprintf("SELECT * FROM %s ORDER BY id LIMIT $1 OFFSET $2", pg.UnitTable),
		limit,
		offset,
	); err != nil {
		return nil, err
	}
	return cities, nil
}

func (c *unitRepository) Update(unit *entity.Unit) error {
	return c.db.Get(unit,
		fmt.Sprintf(`UPDATE %s
				SET
					name = $1
				WHERE id = $2 RETURNING *`, pg.UnitTable),
		unit.Name,
		unit.Id,
	)
}

func (c *unitRepository) Delete(id int64) error {
	if _, err := c.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = $1", pg.UnitTable), id); err != nil {
		return err
	}
	return nil
}
