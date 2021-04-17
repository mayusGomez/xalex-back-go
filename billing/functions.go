package billing

import (
	"github.com/mayusGomez/xalex/shared"
)

func (service *Service) SetMoneyToInt() {
	service.PriceInt = shared.MoneyToInt(service.Price)
	service.CostInt = shared.MoneyToInt(service.Cost)
}

func (service *Service) SetMoneyToFloat() {
	service.Price = shared.IntToMoney(service.PriceInt)
	service.Cost = shared.IntToMoney(service.CostInt)
}

func (service *Service) InactiveService() {
	service.Status = InactiveServStatus
}
