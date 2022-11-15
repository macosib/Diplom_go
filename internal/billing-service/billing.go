package billing_service

import (
	"Diplom_Makarov/internal/utils"
)

type BillingData struct {
	CreateCustomer bool
	Purchase       bool
	Payout         bool
	Recurring      bool
	FraudControl   bool
	CheckoutPage   bool
}

// StartBillingService - Функция запускает сервис для получения данных о о системе Billing из файла формата CSV.
// Результат выполениния - *BillingData.
func StartBillingService() *BillingData {
	res := utils.ReadCsvFile(utils.ConfigData.BillingDataPath)[0][0]

	var newBillingData BillingData

	newBillingData.CreateCustomer = utils.ConvertToBool(res[0])
	newBillingData.Purchase = utils.ConvertToBool(res[1])
	newBillingData.Payout = utils.ConvertToBool(res[2])
	newBillingData.Recurring = utils.ConvertToBool(res[3])
	newBillingData.FraudControl = utils.ConvertToBool(res[4])
	newBillingData.CheckoutPage = utils.ConvertToBool(res[5])

	return &newBillingData
}
