package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/offerni/gomultithreading"
	"github.com/offerni/gomultithreading/brasilapi"
	"github.com/offerni/gomultithreading/viacep"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	addressChannel := make(chan gomultithreading.AddressResponse)

	// hack just to get the first argument
	cep := os.Args[1]

	go getAddressFromBrasilApi(ctx, addressChannel, cep)

	go getAddressFromViaCep(ctx, addressChannel, cep)

	adJson, err := json.Marshal(<-addressChannel)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(adJson))
}

func getAddressFromBrasilApi(ctx context.Context, ch chan<- gomultithreading.AddressResponse, cep string) {
	brasilApiAddress, err := brasilapi.GetAddress(ctx, cep)
	if err != nil {
		panic(err)
	}
	ch <- gomultithreading.AddressResponse{
		Cep:          brasilApiAddress.Cep,
		City:         brasilApiAddress.City,
		Neighborhood: brasilApiAddress.Neighborhood,
		Service:      "brasilapi",
		State:        brasilApiAddress.State,
		Street:       brasilApiAddress.Street,
	}
}

func getAddressFromViaCep(ctx context.Context, ch chan<- gomultithreading.AddressResponse, cep string) {
	viacepAddress, err := viacep.GetAddress(ctx, cep)
	if err != nil {
		panic(err)
	}
	ch <- gomultithreading.AddressResponse{
		Cep:          viacepAddress.Cep,
		City:         viacepAddress.Localidade,
		Neighborhood: viacepAddress.Bairro,
		Service:      "viacep",
		State:        viacepAddress.Uf,
		Street:       viacepAddress.Logradouro,
	}
}
