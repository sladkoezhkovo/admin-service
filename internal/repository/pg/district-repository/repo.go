package districtrepository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
	"github.com/sladkoezhkovo/admin-service/internal/repository/pg"
	districtservice "github.com/sladkoezhkovo/admin-service/internal/service/district-service"
)

var _ districtservice.DistrictRepository = (*districtRepository)(nil)

type districtRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *districtRepository {
	return &districtRepository{
		db: db,
	}
}

func (c *districtRepository) Create(district *entity.District) error {
	return c.db.Get(
		district,
		fmt.Sprintf(`INSERT INTO %s (name, city_id) VALUES ($1, $2) RETURNING id, name, city_id as "city.id"`, pg.DistrictTable),
		district.Name, district.City.Id)
}

func (c *districtRepository) FindById(id int64) (*entity.District, error) {
	var district entity.District
	if err := c.db.Get(
		&district,
		fmt.Sprintf(
			`SELECT 
						d.id, d.name, c.id as "city.id", c.name as "city.name" 
					FROM %s d
					INNER JOIN %s c ON d.city_id = c.id  	
					WHERE d.id = $1`, pg.DistrictTable, pg.CityTable),
		id,
	); err != nil {
		return nil, err
	}
	return &district, nil
}

func (c *districtRepository) FindByName(name string) (*entity.District, error) {
	var district entity.District
	if err := c.db.Get(
		&district,
		fmt.Sprintf(
			`SELECT 
						d.id, d.name, c.id as "city.id", c.name as "city.name" 
					FROM %s d
					INNER JOIN %s c ON d.city_id = c.id  	
					WHERE d.name = $1`, pg.DistrictTable, pg.CityTable),
		name,
	); err != nil {
		return nil, err
	}
	return &district, nil
}

func (c *districtRepository) List(limit, offset int) ([]*entity.District, error) {
	var cities []*entity.District
	if err := c.db.Select(
		&cities,
		fmt.Sprintf(
			`SELECT 
						d.id, d.name, c.name as "city.name"  
					FROM %s d 
					INNER JOIN %s c ON d.city_id = c.id	
					ORDER BY d.id 
					LIMIT $1 OFFSET $2`, pg.DistrictTable, pg.CityTable),
		limit,
		offset,
	); err != nil {
		return nil, err
	}
	return cities, nil
}

func (c *districtRepository) Update(district *entity.District) error {
	return c.db.Get(district,
		fmt.Sprintf(`UPDATE %s
				SET
					name = $1,
					city_id = $2
				WHERE id = $3 RETURNING *`, pg.DistrictTable),
		district.Name,
		district.City.Id,
		district.Id,
	)
}

func (c *districtRepository) Delete(id int64) error {
	if _, err := c.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = $1", pg.DistrictTable), id); err != nil {
		return err
	}
	return nil
}
