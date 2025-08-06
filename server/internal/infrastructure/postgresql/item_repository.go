package postgresql

import (
	"context"
	domainModel "servre/internal/domain/model"
	domainRepository "servre/internal/domain/repository"

	"github.com/jackc/pgx/v5"
)

type ItemRepo struct {
	db *pgx.Conn
}

func NewItemRepo(db *pgx.Conn) domainRepository.ItemRepository {
	return &ItemRepo{db: db}
}

func (r *ItemRepo) Create(ctx context.Context, item *domainModel.ItemObject) error {
	query := `
        INSERT INTO items (
            item_id, name, price, required_level, 
            max_stack, rarity, item_type_id
        ) VALUES ($1, $2, $3, $4, $5, $6, $7)
        ON CONFLICT (item_id) DO UPDATE
        SET name = $2, price = $3, required_level = $4,
            max_stack = $5, rarity = $6, item_type_id = $7
    `

	_, err := r.db.Exec(
		ctx,
		query,
		item.Id,
		item.Name,
		item.Price,
		item.ReqLevel,
		item.MaxStack,
		string(item.Rarity),
		item.ItemType.Name,
	)

	return err
}

func (r *ItemRepo) Delete(ctx context.Context, itemID string) error {

}

func (r *ItemRepo) Update(ctx context.Context, item *domainModel.ItemObject) error {

}

func (r *ItemRepo) FindById(ctx context.Context, itemID string) (*domainModel.ItemObject, error) {

}

func (r *ItemRepo) FindListById(ctx context.Context, itemIDS []string) ([]*domainModel.ItemObject, error) {

}
