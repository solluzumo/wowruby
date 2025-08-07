package domainRepository

import (
	"context"
	item "servre/internal/domain/model"
)

type ItemTypeRepository interface {
	Create(ctx context.Context, item *item.ItemTypeObject) error
	Update(ctx context.Context, item *item.ItemTypeObject) error
	FindById(ctx context.Context, itemID string) (*item.ItemTypeObject, error)
}
