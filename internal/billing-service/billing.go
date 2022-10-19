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

func StartBillingService() *BillingData {
	res := utils.ReadCsvFile("../simulator/skillbox-diploma/billing.data")[0][0]
	//var resultNumber uint8
	//for idx, value := range res {
	//	if value == 49 {
	//		resultNumber += uint8(math.Pow(2, float64(len(res[0][0])-idx-1)))
	//		fmt.Println(resultNumber)
	//	}
	//}
	var newBillingData BillingData

	newBillingData.CreateCustomer = utils.ConvertToBool(res[0])
	newBillingData.Purchase = utils.ConvertToBool(res[1])
	newBillingData.Payout = utils.ConvertToBool(res[2])
	newBillingData.Recurring = utils.ConvertToBool(res[3])
	newBillingData.FraudControl = utils.ConvertToBool(res[4])
	newBillingData.CheckoutPage = utils.ConvertToBool(res[5])

	return &newBillingData
}
