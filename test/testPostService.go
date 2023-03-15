package main

import (
	"fmt"
	ps "github.com/uan190176/postservice"
)

func main() {
	PostToken := "2jfrbrey4fXCX6T8LLJeXL0roRd0Bceq"
	PostAuth := "dmlzdWFsdGVjaEBtYWlsLnJ1OnZpc3VhbCEzMjEh"
	URL := "https://otpravka-api.pochta.ru/1.0/clean/address"

	credentials := ps.PostCredentials{
		URL:   URL,
		Token: PostToken,
		Auth:  PostAuth,
	}

	addr := ps.Address{
		Id:              "1",
		OriginalAddress: "Миасское ул Комсомола д 1",
	}

	var addresses = []ps.Address{addr}

	p := ps.NewPostService(credentials)

	normAddresses, err := p.GetNormalizedAddresses(addresses)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(normAddresses)
}
