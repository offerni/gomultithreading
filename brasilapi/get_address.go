package brasilapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl string = "https://brasilapi.com.br/api/cep/v1"

func GetAddress(ctx context.Context, cep string) (*address, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s", baseUrl, cep), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var address address
	if err := json.Unmarshal(resBody, &address); err != nil {
		return nil, err
	}

	return &address, nil
}

type address struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}
