package itemusecase

import "github.com/Hayato360/go_shop/modules/item/itemRepository"

type (
	ItemUsecaseService interface{}

	itemUsecase struct {
		itemRepository itemrepository.ItemRepositoryService
	}
)

func NewItemUsecase(itemRepository itemrepository.ItemRepositoryService) ItemUsecaseService {
	return &itemUsecase{itemRepository: itemRepository}
}