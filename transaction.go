package main

import (
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
)

func BuildPaymentTransaction(sourceSeed, destinationAddr, amount, memoText string) (string, error) {
	tx, err := build.Transaction(
		build.TestNetwork,
		build.SourceAccount{sourceSeed},
		build.AutoSequence{horizon.DefaultTestNetClient},
		build.MemoText{memoText},
		build.Payment(
			build.Destination{destinationAddr},
			build.NativeAmount{amount},
		),
	)

	if err != nil {
		return "", err
	}

	// Sign the transaction before submitting to the network
	signedTx, err := tx.Sign(sourceSeed)
	if err != nil {
		return "", err
	}

	return signedTx.Base64()
}

func Transfer(sourceSeed, destinationAddr, amount, memoText string) error {
	// Verify destination account exists
	if _, err := horizon.DefaultTestNetClient.LoadAccount(destinationAddr); err != nil {
		return err
	}

	// Build & sign payment transaction
	tx, err := BuildPaymentTransaction(sourceSeed, destinationAddr, amount, memoText)
	if err != nil {
		return err
	}

	// Submit the transaction to the test network
	_, err = horizon.DefaultTestNetClient.SubmitTransaction(tx)
	return err
}

func main() {
	err := Transfer("SCI2RW2WK2GBVU2HNB4ABFH52EHSMAR2PV2SLE4VFS6OFN4RZYELKOR4", "GDAVIJXB6QXBKU66CBKOLHEMOUJOFQZKB7M3Y4IDXFTY2F5K6B25LDT4", "500", "hey there")
	if err != nil {
		panic(err)
	}
}
