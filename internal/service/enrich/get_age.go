package enrich

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type agifyResponse struct {
	Age int `json:"age"`
}

func GetAge(ctx context.Context, name string) (int, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var ar agifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&ar); err != nil {
		return 0, err
	}

	return ar.Age, nil
}
