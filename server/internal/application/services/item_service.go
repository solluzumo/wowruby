package item

import item "servre/internal/domain/repository"

type ItemService interface {
}

type itemServiceImpl struct {
	repo item.ItemRepository
}
