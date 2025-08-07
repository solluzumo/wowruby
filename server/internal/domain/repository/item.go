package domainRepository

import (
	"context"
	item "servre/internal/domain/model"
)

type ItemRepository interface {
	Create(ctx context.Context, item *item.ItemObject) error
	Delete(ctx context.Context, itemID string) error
	Update(ctx context.Context, item *item.ItemObject) error
	FindById(ctx context.Context, itemID string) (*item.ItemObject, error)
	FindListById(ctx context.Context, itemIDS []string) ([]*item.ItemObject, error)
}
