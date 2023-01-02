package thirdparty

import (
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

func OrderMidtrans(orderId string, price int64) *snap.Response {
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
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT")
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New(os.Getenv("MIDTRANS_SERVER"), midtrans.Sandbox)

	res, _ := c.CheckTransaction(orderId)
	return res
}

func OrderMidtransCore(orderId string, grossAmmount int64, paymentType midtrans.Bank) *coreapi.ChargeResponse {
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
