package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateTestAccount(address string) error {
	resp, err := http.Get("https://horizon-testnet.stellar.org/friendbot?addr=" + address)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if _, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := CreateTestAccount("GB5GY3VURL6ITZ6SD53YUCFCIKD65FV4WD44ACN24DLDPVDWBZ6MNS2Q"); err != nil {
		panic(err)
	}
}
