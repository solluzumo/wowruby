package applicationService

import (
	"context"
	itemDto "servre/internal/application/dto"
	itemRepository "servre/internal/domain/repository"
)

type ItemService interface {
	CreateItem(ctx context.Context, cmd itemDto.ItemCreationDto) (*itemDto.ItemResponseDto, error)
	DeleteItem(ctx context.Context, id string) (*itemDto.ItemResponseDto, error)
	UpdateItem(ctx context.Context) (*itemDto.ItemResponseDto, error)
	GetItemById(ctx context.Context, id string) (*itemDto.ItemDetailsDto, error)
	ListItems(ctx context.Context, itemIDS []string) ([]*itemDto.ItemDetailsDto, error)
}

type itemServiceImpl struct {
	repo itemRepository.ItemRepository
}

func NewItemService(repo itemRepository.ItemRepository) ItemService {
	return &itemServiceImpl{repo: repo}
}

func (s *itemServiceImpl) CreateItem(ctx context.Context, cmd itemDto.ItemCreationDto) (*itemDto.ItemResponseDto, error) {

}
func (s *itemServiceImpl) DeleteItem(ctx context.Context, id string) (*itemDto.ItemResponseDto, error) {

}
func (s *itemServiceImpl) UpdateItem(ctx context.Context) (*itemDto.ItemResponseDto, error) {

}
func (s *itemServiceImpl) GetItemById(ctx context.Context, id string) (*itemDto.ItemDetailsDto, error) {

}
func (s *itemServiceImpl) ListItems(ctx context.Context, itemIDS []string) ([]*itemDto.ItemDetailsDto, error) {

}
