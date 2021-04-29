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

func (event *Event) SetMoneyToInt() {

	for i, detail := range event.Services {
		detail.PriceInt = shared.MoneyToInt(detail.Price)
		detail.CostInt = shared.MoneyToInt(detail.Cost)
		event.Services[i] = detail

	}

}

func (event *Event) SetMoneyToFloat() {

	for i, detail := range event.Services {
		detail.Price = shared.IntToMoney(detail.PriceInt)
		detail.Cost = shared.IntToMoney(detail.CostInt)
		event.Services[i] = detail
	}
}
