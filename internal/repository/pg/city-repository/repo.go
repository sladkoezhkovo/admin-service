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
	return c.db.QueryRowx(
		"INSERT INTO city(name) VALUES ($1) RETURNING id",
		city.Name,
	).Scan(&city.Id)
}

func (c *cityRepository) FindById(id int64) (*entity.City, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cityRepository) FindByName(name string) (*entity.City, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cityRepository) List(limit, offset int) ([]*entity.City, error) {
	//TODO implement me
	panic("implement me")
}

func (c *cityRepository) Update(city *entity.City) error {
	//TODO implement me
	panic("implement me")
}

func (c *cityRepository) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}
