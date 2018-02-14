package main

import (
	"context"
	"fmt"
	"github.com/stellar/go/clients/horizon"
)

func main() {
	const address = "GDAVIJXB6QXBKU66CBKOLHEMOUJOFQZKB7M3Y4IDXFTY2F5K6B25LDT4"

	ctx := context.Background()

	cursor := horizon.Cursor("now")

	fmt.Println("Waiting for a payment...")

	err := horizon.DefaultTestNetClient.StreamPayments(ctx, address, &cursor, func(payment horizon.Payment) {
		fmt.Println("Payment type", payment.Type)
		fmt.Println("Payment Paging Token", payment.PagingToken)
		fmt.Println("Payment From", payment.From)
		fmt.Println("Payment To", payment.To)
		fmt.Println("Payment Asset Type", payment.AssetType)
		fmt.Println("Payment Asset Code", payment.AssetCode)
		fmt.Println("Payment Asset Issuer", payment.AssetIssuer)
		fmt.Println("Payment Amount", payment.Amount)
		fmt.Println("Payment Memo Type", payment.Memo.Type)
		fmt.Println("Payment Memo", payment.Memo.Value)
	})

	if err != nil {
		panic(err)
	}

}
