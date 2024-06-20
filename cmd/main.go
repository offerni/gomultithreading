package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/offerni/gomultithreading"
	"github.com/offerni/gomultithreading/brasilapi"
	"github.com/offerni/gomultithreading/viacep"
)

const timeoutDeadline time.Duration = time.Second

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDeadline)
	defer cancel()

	addressChannel := make(chan gomultithreading.AddressResponse)

	// hack just to get the first argument
	cep := os.Args[1]

	go buildAddressFromBrasilApi(ctx, addressChannel, cep)

	go buildAddressFromViaCep(ctx, addressChannel, cep)

	select {
	case address := <-addressChannel:
		adJson, err := json.Marshal(address)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(adJson))
	case <-ctx.Done():
		log.Fatalf("%s Timeout reached, Context canceled", timeoutDeadline)
	}
}

func buildAddressFromBrasilApi(ctx context.Context, ch chan<- gomultithreading.AddressResponse, cep string) {
	// time.Sleep(time.Second) // Uncomment this if you want to test forcing returning results from the other concurrent API call
	brasilApiAddress, err := brasilapi.GetAddress(ctx, cep)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got it from %s:\n", brasilapi.ServiceName)
	ch <- gomultithreading.AddressResponse{
		Cep:          brasilApiAddress.Cep,
		City:         brasilApiAddress.City,
		Neighborhood: brasilApiAddress.Neighborhood,
		Service:      brasilapi.ServiceName,
		State:        brasilApiAddress.State,
		Street:       brasilApiAddress.Street,
	}
}

func buildAddressFromViaCep(ctx context.Context, ch chan<- gomultithreading.AddressResponse, cep string) {
	// time.Sleep(time.Second) // Uncomment this if you want to test forcing returning results from the other concurrent API call
	viacepAddress, err := viacep.GetAddress(ctx, cep)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got it from %s:\n", viacep.ServiceName)
	ch <- gomultithreading.AddressResponse{
		Cep:          viacepAddress.Cep,
		City:         viacepAddress.Localidade,
		Neighborhood: viacepAddress.Bairro,
		Service:      viacep.ServiceName,
		State:        viacepAddress.Uf,
		Street:       viacepAddress.Logradouro,
	}
}
