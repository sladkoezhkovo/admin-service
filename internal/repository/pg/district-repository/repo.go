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
						d.id, d.name, c.name as "city.name" 
					FROM %s d
					INNER JOIN %s c ON d.city_id = c.id  	
					WHERE d.id = $1`, pg.DistrictTable, pg.CityTable),
		id,
	); err != nil {
		return nil, err
	}
	return &district, nil
}

func (c *districtRepository) ListByCityId(cityId int64) ([]*entity.District, error) {
	var dd []*entity.District
	if err := c.db.Select(
		&dd,
		fmt.Sprintf(
			`SELECT 
						d.id, d.name, c.name as "city.name"  
					FROM %s d 
					INNER JOIN %s c ON d.city_id = c.id	
					WHERE c.id = $1 ORDER BY d.id`, pg.DistrictTable, pg.CityTable),
		cityId,
	); err != nil {
		return nil, err
	}
	return dd, nil
}

func (c *districtRepository) List(limit, offset int) ([]*entity.District, int64, error) {
	var dd []*entity.District
	if err := c.db.Select(
		&dd,
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
		return nil, 0, err
	}
	var count int64
	if err := c.db.Get(&count, fmt.Sprintf("SELECT COUNT(id) FROM %s", pg.DistrictTable)); err != nil {
		return nil, 0, err
	}
	return dd, count, nil
}

func (c *districtRepository) Update(district *entity.District) error {
	fmt.Printf("%s: district to update: %v\n", pg.DistrictTable, district)
	return c.db.Get(district,
		fmt.Sprintf(`UPDATE %s
				SET
					name = $1,
					city_id = $2
				WHERE id = $3 RETURNING id, name, city_id as "city.id"`, pg.DistrictTable),
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
