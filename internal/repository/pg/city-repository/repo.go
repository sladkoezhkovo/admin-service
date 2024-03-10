package cityrepository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
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
	return c.db.Get(city, "INSERT INTO city (name) VALUES ($1) RETURNING *", city.Name)
}

func (c *cityRepository) FindById(id int64) (*entity.City, error) {
	var city entity.City
	if err := c.db.Get(&city, "SELECT * FROM city WHERE id = $1", id); err != nil {
		return nil, err
	}
	return &city, nil
}

func (c *cityRepository) FindByName(name string) (*entity.City, error) {
	var city entity.City
	if err := c.db.Get(&city, "SELECT * FROM city WHERE name = $1", name); err != nil {
		return nil, err
	}
	return &city, nil
}

func (c *cityRepository) List(limit, offset int) ([]*entity.City, error) {
	var cities []*entity.City
	if err := c.db.Select(&cities, "SELECT * FROM city ORDER BY id LIMIT $1 OFFSET $2", limit, offset); err != nil {
		return nil, err
	}
	return cities, nil
}

func (c *cityRepository) Update(city *entity.City) error {
	return c.db.Get(city,
		`UPDATE city
				SET
					name = $1
				WHERE id = $2 RETURNING *`, city.Name, city.Id)
}

func (c *cityRepository) Delete(id int64) error {
	if _, err := c.db.Exec("DELETE FROM city WHERE id = $1", id); err != nil {
		return err
	}
	return nil
}
