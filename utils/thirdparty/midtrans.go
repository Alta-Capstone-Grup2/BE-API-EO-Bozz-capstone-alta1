package thirdparty

import (
	cfg "capstone-alta1/config"
	"capstone-alta1/utils/helper"
	"errors"
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

func OrderMidtransCore(orderId string, grossAmmount int64, paymentType midtrans.Bank, orderTime string) *coreapi.ChargeResponse {
	Init()
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER")
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT")
	midtrans.Environment = midtrans.Sandbox
	c := coreapi.Client{}
	c.New(os.Getenv("MIDTRANS_SERVER"), midtrans.Sandbox)

	req := &coreapi.ChargeReq{
		PaymentType:  coreapi.PaymentTypeBankTransfer,
		BankTransfer: &coreapi.BankTransferDetails{Bank: paymentType},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderId,
			GrossAmt: grossAmmount,
		},
		CustomExpiry: &coreapi.CustomExpiry{
			OrderTime:      orderTime,
			ExpiryDuration: cfg.PAYMENT_EXPIRED_DURATION,
			Unit:           cfg.PAYMENT_EXPIRED_UNIT,
		},
	}

	chargeRes, _ := coreapi.ChargeTransaction(req)
	return chargeRes
}

func GetVABank(input string) (midtrans.Bank, error) {
	helper.LogDebug("get va bank ", input)
	if input == string(cfg.VABNI) {
		return midtrans.BankBni, nil
	}
	if input == string(cfg.VABca) {
		return midtrans.BankBca, nil
	}
	if input == string(cfg.VABri) {
		return midtrans.BankBri, nil
	}
	if input == string(cfg.VAPermata) {
		return midtrans.BankPermata, nil
	}
	helper.LogDebug("Thirdpary - Midtrans - GetVABank | Error get VA Bank. Input  = ", input)
	return "", errors.New("Failed to get VA Bank. Please check input again or choose other payment method.")
}

func GetVABankTitle(input string) string {
	helper.LogDebug("get va bank ", input)
	if input == string(cfg.VABNI) {
		return "BNI Virtual Account"
	}
	if input == string(cfg.VABca) {
		return "BCA Virtual Account"
	}
	if input == string(cfg.VABri) {
		return "BRI Virtual Account"
	}
	if input == string(cfg.VAPermata) {
		return "Permata Virtual Account"
	}
	helper.LogDebug("Thirdpary - Midtrans - GetVABankTitle | Error get VA Bank. Input  = ", input)
	return ""
}
