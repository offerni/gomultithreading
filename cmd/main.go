package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/offerni/gomultithreading"
	"github.com/offerni/gomultithreading/viacep"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// address, err := brasilapi.GetAddress(ctx, "17055250")
	// if err != nil {
	// 	panic(err)
	// }

	// addressResponse := gomultithreading.AddressResponse{
	// 	Cep:          address.Cep,
	// 	City:         address.City,
	// 	Neighborhood: address.Neighborhood,
	// 	Service:      "brasilapi",
	// 	State:        address.State,
	// 	Street:       address.Street,
	// }

	address, err := viacep.GetAddress(ctx, "17055250")
	if err != nil {
		panic(err)
	}

	addressResponse := gomultithreading.AddressResponse{
		Cep:          address.Cep,
		City:         address.Localidade,
		Neighborhood: address.Bairro,
		Service:      "viacep",
		State:        address.Uf,
		Street:       address.Logradouro,
	}

	adJson, err := json.Marshal(addressResponse)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(adJson))
}
