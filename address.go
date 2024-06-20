package gomultithreading

type AddressResponse struct {
	Cep          string `json:"cep"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Service      string `json:"service"`
	State        string `json:"state"`
	Street       string `json:"street"`
}
