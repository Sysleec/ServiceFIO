package enrich

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type nationalizeResponse struct {
	Country []struct {
		CountryID string `json:"country_id"`
	} `json:"country"`
}

func GetNationality(ctx context.Context, name string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var nr nationalizeResponse
	if err := json.NewDecoder(resp.Body).Decode(&nr); err != nil {
		return "", err
	}

	if len(nr.Country) == 0 {
		return "", fmt.Errorf("no countries found for name %s", name)
	}

	return nr.Country[0].CountryID, nil
}
