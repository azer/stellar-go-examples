package main

import (
	"errors"
	"fmt"
	"github.com/stellar/go/clients/horizon"
)

func GetBalance(balanceType, address string) (string, error) {
	account, err := horizon.DefaultTestNetClient.LoadAccount(address)
	if err != nil {
		return "", err
	}

	for _, balance := range account.Balances {
		if balance.Type == balanceType {
			return balance.Balance, nil
		}
	}

	return "", errors.New("Invalid balance type")
}

func main() {
	balance, err := GetBalance("native", "GB5GY3VURL6ITZ6SD53YUCFCIKD65FV4WD44ACN24DLDPVDWBZ6MNS2Q")
	if err != nil {
		panic(err)
	}

	fmt.Println(balance)
}
