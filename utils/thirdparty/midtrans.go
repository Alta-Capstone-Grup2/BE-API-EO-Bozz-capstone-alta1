package thirdparty

import (
	"capstone-alta1/utils/helper"
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

func Init() {
	if os.Getenv("MIDTRANS_SERVER") == "" {
		helper.LogDebug("MIDTRANS_SERVER value not found.")
	}
	if os.Getenv("MIDTRANS_CLIENT") == "" {
		helper.LogDebug("MIDTRANS_CLIENT value not found.")
	}
}

func OrderMidtrans(orderId string, price int64) *snap.Response {
	Init()
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT")
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New(os.Getenv("MIDTRANS_SERVER"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: price,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	}

	snapResp, _ := snap.CreateTransaction(req)
	return snapResp
}

func CheckMidtrans(orderId string) *coreapi.TransactionStatusResponse {
	Init()
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT")
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New(os.Getenv("MIDTRANS_SERVER"), midtrans.Sandbox)

	res, _ := c.CheckTransaction(orderId)
	return res
}

func OrderMidtransCore(orderId string, grossAmmount int64, paymentType midtrans.Bank) *coreapi.ChargeResponse {
	Init()
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT")
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New(os.Getenv("MIDTRANS_SERVER"), midtrans.Sandbox)

	req := &coreapi.ChargeReq{
		PaymentType: coreapi.CoreapiPaymentType(paymentType),
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: grossAmmount,
		},
	}

	chargeRes, _ := coreapi.ChargeTransaction(req)
	return chargeRes
}
