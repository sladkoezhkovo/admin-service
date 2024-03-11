package cityrepository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
	"github.com/sladkoezhkovo/admin-service/internal/repository/pg"
	cityservice "github.com/sladkoezhkovo/admin-service/internal/service/city-service"
)

var _ cityservice.CityRepository = (*cityRepository)(nil)

type cityRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *cityRepository {
	return &cityRepository{
		db: db,
	}
}

func (c *cityRepository) Create(city *entity.City) error {
	return c.db.Get(
		city,
		fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING *", pg.CityTable),
		city.Name)
}

func (c *cityRepository) FindById(id int64) (*entity.City, error) {
	var city entity.City
	if err := c.db.Get(
		&city,
		fmt.Sprintf("SELECT * FROM %s WHERE id = $1", pg.CityTable),
		id,
	); err != nil {
		return nil, err
	}
	return &city, nil
}

func (c *cityRepository) ListByName(name string, limit, offset int32) ([]*entity.City, error) {
	var cc []*entity.City
	if err := c.db.Select(
		&cc,
		fmt.Sprintf(`SELECT * FROM %s WHERE name ILIKE  $1  LIMIT $2 OFFSET $3`, pg.CityTable),
		"%"+name+"%", limit, offset,
	); err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *cityRepository) List(limit, offset int) ([]*entity.City, error) {
	var cities []*entity.City
	if err := c.db.Select(
		&cities,
		fmt.Sprintf("SELECT * FROM %s ORDER BY id LIMIT $1 OFFSET $2", pg.CityTable),
		limit,
		offset,
	); err != nil {
		return nil, err
	}
	return cities, nil
}

func (c *cityRepository) Update(city *entity.City) error {
	return c.db.Get(city,
		fmt.Sprintf(`UPDATE %s
				SET
					name = $1
				WHERE id = $2 RETURNING *`, pg.CityTable),
		city.Name,
		city.Id,
	)
}

func (c *cityRepository) Delete(id int64) error {
	if _, err := c.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = $1", pg.CityTable), id); err != nil {
		return err
	}
	return nil
}
