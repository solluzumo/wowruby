package postgresql

import (
	"context"
	domainModel "servre/internal/domain/model"
	domainRepository "servre/internal/domain/repository"

	"github.com/jackc/pgx/v5"
)

type ItemTypeRepo struct {
	db *pgx.Conn
}

func NewItemTypeRepo(db *pgx.Conn) domainRepository.ItemTypeRepository {
	return &ItemTypeRepo{db: db}
}

func (r *ItemTypeRepo) Create(ctx context.Context, item *domainModel.ItemTypeObject) error {

}

func (r *ItemTypeRepo) Update(ctx context.Context, item *domainModel.ItemTypeObject) error {

}

func (r *ItemTypeRepo) FindById(ctx context.Context, id string) (*domainModel.ItemTypeObject, error) {
	query := `SELECT name, description FROM item_types WHERE name = $1`
	row := r.db.QueryRow(ctx, query, id)

	var itemType domainModel.ItemTypeObject
	err := row.Scan(&itemType.Name, &itemType.Description)
	if err != nil {
		return nil, err
	}

	return &itemType, nil
}
