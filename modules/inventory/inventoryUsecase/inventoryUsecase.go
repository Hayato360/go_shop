package inventoryusecase

import inventoryrepository "github.com/Hayato360/go_shop/modules/inventory/inventoryRepository"

type (
	InventoryUsecaseService interface{}

	inventoryUsecase struct {
		inventoryRepository inventoryrepository.InventoryRepositoryService
	}
)

func NewInventoryUsecase(inventoryRepository inventoryrepository.InventoryRepositoryService) InventoryUsecaseService {
	return &inventoryUsecase{
		inventoryRepository: inventoryRepository,
	}
}