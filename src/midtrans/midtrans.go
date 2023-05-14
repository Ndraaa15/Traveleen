package midtrans

import (
	"gin/src/entity"
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/example"
	"github.com/midtrans/midtrans-go/snap"
)

var s snap.Client

func InitializeSnapClient() {
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
}

func CreateTransaction(user entity.User, cart entity.Cart) (*snap.Response, error) {
	s.Options.SetPaymentAppendNotification("https://traveleen.vercel.app/history")

	s.Options.SetPaymentOverrideNotification("https://traveleen.vercel.app/history")

	resp, err := s.CreateTransaction(GenerateSnapReq(user, cart))
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func GenerateSnapReq(user entity.User, cart entity.Cart) *snap.Request {

	custAddress := &midtrans.CustomerAddress{
		FName: user.Username,
		Phone: user.Contact,
		City:  user.Region,
	}

	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "MID-GO-ID-" + example.Random(),
			GrossAmt: int64(cart.TotalPrice),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName:    user.Username,
			Email:    user.Email,
			Phone:    user.Contact,
			BillAddr: custAddress,
			ShipAddr: custAddress,
		},
		EnabledPayments: snap.AllSnapPaymentType,
	}
	return snapReq
}
