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
	res := utils.ReadCsvFile("../simulator/skillbox-diploma/billing.data")
	//var resultNumber uint8
	//for idx, value := range res[0][0] {
	//	if value == 49 {
	//		resultNumber += uint8(math.Pow(2, float64(len(res[0][0])-idx-1)))
	//		fmt.Println(resultNumber)
	//	}
	//}
	var newBillingData BillingData
	newBillingData.CreateCustomer = utils.ConvertToBool(res[0][0][0])
	newBillingData.Purchase = utils.ConvertToBool(res[0][0][1])
	newBillingData.Payout = utils.ConvertToBool(res[0][0][2])
	newBillingData.Recurring = utils.ConvertToBool(res[0][0][3])
	newBillingData.FraudControl = utils.ConvertToBool(res[0][0][4])
	newBillingData.CheckoutPage = utils.ConvertToBool(res[0][0][5])
	return &newBillingData
}
