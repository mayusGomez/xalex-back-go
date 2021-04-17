package shared

func MoneyToInt(money float32) int {
	value := int(money * MoneyFactor)
	return value
}

func IntToMoney(money int) float32 {
	value := float32(money) / MoneyFactor
	return value
}
