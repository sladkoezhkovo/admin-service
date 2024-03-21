package packagingrepository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sladkoezhkovo/admin-service/internal/entity"
	"github.com/sladkoezhkovo/admin-service/internal/repository/pg"
	packagingservice "github.com/sladkoezhkovo/admin-service/internal/service/packaging-service"
)

var _ packagingservice.PackagingRepository = (*packagingRepository)(nil)

type packagingRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *packagingRepository {
	return &packagingRepository{
		db: db,
	}
}

func (c *packagingRepository) Create(packaging *entity.Packaging) error {
	return c.db.Get(
		packaging,
		fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING *", pg.PackagingTable),
		packaging.Name)
}

func (c *packagingRepository) FindById(id int64) (*entity.Packaging, error) {
	var packaging entity.Packaging
	if err := c.db.Get(
		&packaging,
		fmt.Sprintf("SELECT * FROM %s WHERE id = $1", pg.PackagingTable),
		id,
	); err != nil {
		return nil, err
	}
	return &packaging, nil
}

func (c *packagingRepository) List(limit, offset int) ([]*entity.Packaging, int64, error) {
	var pp []*entity.Packaging
	if err := c.db.Select(
		&pp,
		fmt.Sprintf("SELECT * FROM %s ORDER BY id LIMIT $1 OFFSET $2", pg.PackagingTable),
		limit,
		offset,
	); err != nil {
		return nil, 0, err
	}
	var count int64
	if err := c.db.Get(&count, fmt.Sprintf("SELECT COUNT(id) FROM %s", pg.PackagingTable)); err != nil {
		return nil, 0, err
	}
	return pp, count, nil
}

func (c *packagingRepository) Update(packaging *entity.Packaging) error {
	return c.db.Get(packaging,
		fmt.Sprintf(`UPDATE %s
				SET
					name = $1
				WHERE id = $2 RETURNING *`, pg.PackagingTable),
		packaging.Name,
		packaging.Id,
	)
}

func (c *packagingRepository) Delete(id int64) error {
	if _, err := c.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id = $1", pg.PackagingTable), id); err != nil {
		return err
	}
	return nil
}
